package middleware

import (
	"fmt"
	"multi-tenant/exception"

	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Panic occured: ", err.(error).Error())
				exception.ErrorHandler(ctx, err)
				ctx.Abort()
				return
			}
		}()
		ctx.Next()
	}
}
