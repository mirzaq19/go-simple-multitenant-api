package exception

import (
	"multi-tenant/model/web"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler(ctx *gin.Context, err any) {
	switch value := err.(type) {
	case NotFoundError, InvariantError:
		appError := value.(ApplicationError)
		web.ErrorResponse(ctx, appError.GetStatusCode(), appError.GetErrorCode(), appError.GetErrorName(), appError)
	case validator.ValidationErrors:
		validationError := value
		web.ErrorResponse(ctx, 422, 422, "VALIDATION_ERROR", validationError)
	default:
		internalError := value.(error)
		web.ErrorResponse(ctx, 500, 500, "INTERNAL_SERVER_ERROR", internalError)
	}
}
