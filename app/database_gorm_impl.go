package app

import (
	"fmt"
	"log"
	"multi-tenant/exception"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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

func (t *TenantDBManagerImpl) GetConnection(tenantName string) (*gorm.DB, error) {
	conn, ok := t.connections.Load(tenantName)
	if !ok {
		return nil, exception.NewInternalServerError(500, "Database connection not found:"+tenantName)
	}

	return conn.(*gorm.DB), nil
}

func (t *TenantDBManagerImpl) OpenConnection(tenant TenantDB) {
	db, err := gorm.Open(mysql.Open(tenant.DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to tenant %s with DSN %s: %v", tenant.Name, tenant.DSN, err)
		return
	}

	// Tune the connection pool
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("❌ Failed to get SQL DB for tenant %s: %v", tenant.Name, err)
		return
	}

	// Example tuning — adjust depending on your workload and DB capacity
	sqlDB.SetMaxOpenConns(20)               // Max open connections
	sqlDB.SetMaxIdleConns(10)               // Max idle connections
	sqlDB.SetConnMaxLifetime(1 * time.Hour) // Recycle connections every hour
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	fmt.Printf("✅ Connected to database for tenant %s with pool configured\n", tenant.Name)

	// Store the pool for future requests
	t.connections.Store(tenant.Name, db)
}
