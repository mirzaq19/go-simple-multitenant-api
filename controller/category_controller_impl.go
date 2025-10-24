package controller

import (
	"multi-tenant/exception"
	"multi-tenant/helper"
	"multi-tenant/model/web"
	"multi-tenant/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) *CategoryControllerImpl {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (c *CategoryControllerImpl) GetTenant(ctx *gin.Context) string {
	tenantName, exists := ctx.Get("tenantName")
	if !exists {
		helper.PanicIfError(exception.NewInvariantError(400, "Failed to get tenant name"))
	}

	return tenantName.(string)
}

func (c *CategoryControllerImpl) Create(ctx *gin.Context) {
	tenantName := c.GetTenant(ctx)

	var categoryCreateRequest web.CategoryCreateRequest
	if err := ctx.ShouldBindJSON(&categoryCreateRequest); err != nil {
		panic(err)
	}

	category := c.CategoryService.Create(tenantName, categoryCreateRequest)
	web.SuccesResponse(ctx, http.StatusCreated, "Category created succesfully", category)
}

func (c *CategoryControllerImpl) Update(ctx *gin.Context) {
	tenantName := c.GetTenant(ctx)

	var categoryUpdateRequest web.CategoryUpdateRequest
	if err := ctx.ShouldBindJSON(&categoryUpdateRequest); err != nil {
		panic(err)
	}

	categoryId := ctx.Param("id")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)
	categoryUpdateRequest.Id = id

	category := c.CategoryService.Update(tenantName, categoryUpdateRequest)
	web.SuccesResponse(ctx, http.StatusOK, "Category updated succesfully", category)
}

func (c *CategoryControllerImpl) Delete(ctx *gin.Context) {
	tenantName := c.GetTenant(ctx)

	categoryId := ctx.Param("id")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	c.CategoryService.Delete(tenantName, uint(id))
	web.SuccesResponse(ctx, http.StatusOK, "Category deleted succesfully", nil)
}

func (c *CategoryControllerImpl) Show(ctx *gin.Context) {
	tenantName := c.GetTenant(ctx)

	categoryId := ctx.Param("id")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	category := c.CategoryService.FindById(tenantName, uint(id))
	web.SuccesResponse(ctx, http.StatusOK, "Success fetch category", category)
}

func (c *CategoryControllerImpl) Index(ctx *gin.Context) {
	tenantName := c.GetTenant(ctx)

	category := c.CategoryService.FindAll(tenantName)
	web.SuccesResponse(ctx, http.StatusOK, "Success fetch categories", category)
}
