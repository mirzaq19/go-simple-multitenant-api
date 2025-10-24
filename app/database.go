package app

import (
	"gorm.io/gorm"
)

type TenantDBManager interface {
	GetConnection(tenantId string) (*gorm.DB, error)
	OpenConnection(tenant TenantDB)
}
