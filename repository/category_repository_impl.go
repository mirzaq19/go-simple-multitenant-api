package repository

import (
	"errors"
	"multi-tenant/helper"
	"multi-tenant/model/domain"

	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (r *CategoryRepositoryImpl) Create(tx *gorm.DB, category domain.Category) domain.Category {
	err := tx.Create(&category).Error
	helper.PanicIfError(err)

	return category
}

func (r *CategoryRepositoryImpl) Update(tx *gorm.DB, category domain.Category) domain.Category {
	err := tx.Save(&category).Error
	helper.PanicIfError(err)

	return category
}

func (r *CategoryRepositoryImpl) Delete(tx *gorm.DB, id uint) {
	err := tx.Delete(&domain.Category{}, id).Error
	helper.PanicIfError(err)
}

func (r *CategoryRepositoryImpl) FindById(tx *gorm.DB, id uint) (domain.Category, error) {
	category := domain.Category{}
	err := tx.First(&category, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return category, errors.New("category not found")
	}

	helper.PanicIfError(err)

	return category, nil
}

func (r *CategoryRepositoryImpl) FindAll(tx *gorm.DB) []domain.Category {
	categories := []domain.Category{}

	err := tx.Find(&categories).Error
	helper.PanicIfError(err)
	return categories
}
