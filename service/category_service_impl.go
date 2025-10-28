package service

import (
	"multi-tenant/app"
	"multi-tenant/exception"
	"multi-tenant/helper"
	"multi-tenant/model/domain"
	"multi-tenant/model/web"
	"multi-tenant/repository"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CategoryServiceImpl struct {
	DBManager          app.TenantDBManager
	CategoryRepository repository.CategoryRepository
	Validate           *validator.Validate
}

func NewCategoryService(
	DBManager app.TenantDBManager,
	categoryRepository repository.CategoryRepository,
	validate *validator.Validate,
) CategoryService {
	return &CategoryServiceImpl{
		DBManager:          DBManager,
		CategoryRepository: categoryRepository,
		Validate:           validate,
	}
}

func (s *CategoryServiceImpl) GetDB(tenantName string) app.TenantDBInstance {
	db, err := s.DBManager.GetConnection(tenantName)
	helper.PanicIfError(err)

	return db
}

func (s *CategoryServiceImpl) Create(tenantName string, request web.CategoryCreateRequest) web.CategoryResponse {
	err := s.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := s.GetDB(tenantName).GetTransactionInstance().(*gorm.DB)
	db := app.NewTenantDBInstance(tx)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = s.CategoryRepository.Create(db, category)
	return helper.ToCategoryResponse(category)
}

func (s *CategoryServiceImpl) Update(tenantName string, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := s.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := s.GetDB(tenantName).GetTransactionInstance().(*gorm.DB)
	db := app.NewTenantDBInstance(tx)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		ID:   uint(request.Id),
		Name: request.Name,
	}

	category = s.CategoryRepository.Update(db, category)
	return helper.ToCategoryResponse(category)
}

func (s *CategoryServiceImpl) Delete(tenantName string, id uint) {
	tx := s.GetDB(tenantName).GetTransactionInstance().(*gorm.DB)
	db := app.NewTenantDBInstance(tx)
	defer helper.CommitOrRollback(tx)

	category, err := s.CategoryRepository.FindById(db, id)
	if err != nil {
		panic(exception.NewNotFoundError(404001, "Category is not found"))
	}

	s.CategoryRepository.Delete(db, category.ID)
}

func (s *CategoryServiceImpl) FindById(tenantName string, id uint) web.CategoryResponse {
	db := s.GetDB(tenantName)

	category, err := s.CategoryRepository.FindById(db, id)
	if err != nil {
		panic(exception.NewNotFoundError(404002, "Category is not found"))
	}

	return helper.ToCategoryResponse(category)
}

func (s *CategoryServiceImpl) FindAll(tenantName string) []web.CategoryResponse {
	db := s.GetDB(tenantName)

	categories := s.CategoryRepository.FindAll(db)
	return helper.ToCategoryResponses(categories)
}
