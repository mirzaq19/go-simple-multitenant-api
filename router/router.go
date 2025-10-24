package router

import (
	"multi-tenant/app"
	"multi-tenant/controller"
	"multi-tenant/middleware"
	"multi-tenant/repository"
	"multi-tenant/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.TenantMiddleware(), middleware.RecoveryMiddleware())

	dbManager := app.NewTenantDBManager(app.TenantsDB)
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(dbManager, categoryRepository, validate)
	categoryController := controller.NewCategoryController(categoryService)

	categoryRouter := router.Group("/categories")

	{
		categoryRouter.POST("/", categoryController.Create)
		categoryRouter.PUT("/:id", categoryController.Update)
		categoryRouter.DELETE("/:id", categoryController.Delete)
		categoryRouter.GET("/:id", categoryController.Show)
		categoryRouter.GET("/", categoryController.Index)
	}

	return router
}
