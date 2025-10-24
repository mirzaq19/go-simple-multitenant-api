package web

import (
	"github.com/gin-gonic/gin"
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

func ErrorResponse(ctx *gin.Context, statusCode int, errorCode int, message string, err error) {
	ctx.JSON(statusCode, APIResponse{
		Code:    errorCode,
		Message: message,
		Error:   err.Error(),
	})
}
