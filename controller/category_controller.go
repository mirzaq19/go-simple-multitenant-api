package controller

import "github.com/gin-gonic/gin"

type CategoryController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Show(ctx *gin.Context)
	Index(ctx *gin.Context)
}
