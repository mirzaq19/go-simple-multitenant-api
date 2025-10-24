package middleware

import (
	"errors"
	"multi-tenant/app"
	"multi-tenant/constant"
	"multi-tenant/model/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TenantMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tenantId := ctx.GetHeader(constant.TenantIdHeader)
		if tenantId == "" {
			web.ErrorResponse(ctx, http.StatusBadRequest, 400, "BAD_REQUEST", errors.New("missing X-Tenant-ID"))
			ctx.Abort()
			return
		}

		_, exists := app.TenantsDB[tenantId]
		if !exists {
			web.ErrorResponse(ctx, http.StatusBadRequest, 400, "BAD_REQUEST", errors.New("invalid tenant id"))
			ctx.Abort()
			return
		}

		ctx.Set("tenantName", tenantId)
		ctx.Next()
	}
}
