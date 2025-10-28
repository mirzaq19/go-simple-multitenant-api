package repository

import (
	"errors"
	"multi-tenant/app"
	"multi-tenant/helper"
	"multi-tenant/model/domain"

	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (r *CategoryRepositoryImpl) Create(db app.TenantDBInstance, category domain.Category) domain.Category {
	tx := db.GetInstance().(*gorm.DB)
	err := tx.Create(&category).Error
	helper.PanicIfError(err)

	return category
}

func (r *CategoryRepositoryImpl) Update(db app.TenantDBInstance, category domain.Category) domain.Category {
	tx := db.GetInstance().(*gorm.DB)
	err := tx.Save(&category).Error
	helper.PanicIfError(err)

	return category
}

func (r *CategoryRepositoryImpl) Delete(db app.TenantDBInstance, id uint) {
	tx := db.GetInstance().(*gorm.DB)
	err := tx.Delete(&domain.Category{}, id).Error
	helper.PanicIfError(err)
}

func (r *CategoryRepositoryImpl) FindById(db app.TenantDBInstance, id uint) (domain.Category, error) {
	category := domain.Category{}
	tx := db.GetInstance().(*gorm.DB)
	err := tx.First(&category, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return category, errors.New("category not found")
	}

	helper.PanicIfError(err)

	return category, nil
}

func (r *CategoryRepositoryImpl) FindAll(db app.TenantDBInstance) []domain.Category {
	categories := []domain.Category{}

	tx := db.GetInstance().(*gorm.DB)
	err := tx.Find(&categories).Error
	helper.PanicIfError(err)
	return categories
}
