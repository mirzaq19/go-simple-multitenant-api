package app

type TenantDBInstance interface {
	GetInstance() any
	GetTransactionInstance() any
}

type TenantDBManager interface {
	GetConnection(tenantId string) (TenantDBInstance, error)
	OpenConnection(tenant TenantDB)
}
