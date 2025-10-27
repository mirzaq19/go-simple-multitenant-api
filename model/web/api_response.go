package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type APIResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Error   any    `json:"error,omitempty"`
}

func SuccesResponse(ctx *gin.Context, code int, message string, data any) {
	if len(message) == 0 {
		message = "Success"
	}

	ctx.JSON(code, APIResponse{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func ValidationErrorResponse(ctx *gin.Context, validationMessage validator.ValidationErrorsTranslations) {
	ctx.JSON(http.StatusUnprocessableEntity, APIResponse{
		Code:    http.StatusUnprocessableEntity,
		Message: "VALIDATION_ERROR",
		Error:   validationMessage,
	})
}

func ErrorResponse(ctx *gin.Context, statusCode int, errorCode int, message string, err error) {
	ctx.JSON(statusCode, APIResponse{
		Code:    errorCode,
		Message: message,
		Error:   err.Error(),
	})
}
