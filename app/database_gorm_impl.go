package app

import (
	"fmt"
	"multi-tenant/exception"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TenantDBInstanceImpl struct {
	db *gorm.DB
}

func NewTenantDBInstance(db *gorm.DB) TenantDBInstance {
	return &TenantDBInstanceImpl{db}
}

func (tdbi *TenantDBInstanceImpl) GetInstance() any {
	return tdbi.db
}

func (tdbi *TenantDBInstanceImpl) GetTransactionInstance() any {
	return tdbi.db.Begin()
}

type TenantDBManagerImpl struct {
	connections *sync.Map
}

func NewTenantDBManager(tenants map[string]TenantDB) TenantDBManager {
	tenantDBManager := &TenantDBManagerImpl{
		connections: &sync.Map{},
	}

	for _, tenant := range tenants {
		tenantDBManager.OpenConnection(tenant)
	}

	return tenantDBManager
}

func (t *TenantDBManagerImpl) GetConnection(tenantName string) (TenantDBInstance, error) {
	if conn, ok := t.connections.Load(tenantName); ok {
		return conn.(TenantDBInstance), nil
	}

	tenantDB := TenantsDB[tenantName]
	if newConn := t.OpenConnection(tenantDB); newConn != nil {
		return newConn, nil
	}

	return nil, exception.NewInternalServerError(500, "Failed connect to database:"+tenantName)
}

func (t *TenantDBManagerImpl) OpenConnection(tenant TenantDB) TenantDBInstance {
	db, err := gorm.Open(mysql.Open(tenant.DSN), &gorm.Config{})
	if err != nil {
		fmt.Printf("❌ Failed to connect to tenant %s with DSN %s: %v", tenant.Name, tenant.DSN, err)
		return nil
	}

	// Tune the connection pool
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Printf("❌ Failed to get SQL DB for tenant %s: %v", tenant.Name, err)
		return nil
	}

	// Example tuning — adjust depending on your workload and DB capacity
	sqlDB.SetMaxOpenConns(20)               // Max open connections
	sqlDB.SetMaxIdleConns(10)               // Max idle connections
	sqlDB.SetConnMaxLifetime(1 * time.Hour) // Recycle connections every hour
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	fmt.Printf("✅ Connected to database for tenant %s with pool configured\n", tenant.Name)

	newDBInstance := NewTenantDBInstance(db)

	// Store the pool for future requests
	t.connections.Store(tenant.Name, newDBInstance)

	return newDBInstance
}
