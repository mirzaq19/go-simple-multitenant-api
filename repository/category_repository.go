package repository

import (
	"multi-tenant/app"
	"multi-tenant/model/domain"
)

type CategoryRepository interface {
	Create(db app.TenantDBInstance, category domain.Category) domain.Category
	Update(db app.TenantDBInstance, category domain.Category) domain.Category
	Delete(db app.TenantDBInstance, id uint)
	FindById(db app.TenantDBInstance, id uint) (domain.Category, error)
	FindAll(db app.TenantDBInstance) []domain.Category
}
