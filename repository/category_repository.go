package repository

import (
	"multi-tenant/model/domain"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(tx *gorm.DB, category domain.Category) domain.Category
	Update(tx *gorm.DB, category domain.Category) domain.Category
	Delete(tx *gorm.DB, id uint)
	FindById(tx *gorm.DB, id uint) (domain.Category, error)
	FindAll(tx *gorm.DB) []domain.Category
}
