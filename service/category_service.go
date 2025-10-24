package service

import (
	"multi-tenant/model/web"
)

type CategoryService interface {
	Create(tenantName string, request web.CategoryCreateRequest) web.CategoryResponse
	Update(tenantName string, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(tenantName string, id uint)
	FindById(tenantName string, id uint) web.CategoryResponse
	FindAll(tenantName string) []web.CategoryResponse
}
