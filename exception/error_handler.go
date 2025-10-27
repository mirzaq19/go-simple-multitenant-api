package exception

import (
	"multi-tenant/model/web"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	UniTranslator        *ut.UniversalTranslator
	ValidationTranslator ut.Translator
)

func NewValidatior() *validator.Validate {
	english := en.New()
	UniTranslator = ut.New(english, english)
	ValidationTranslator, _ = UniTranslator.GetTranslator("en")

	validate := validator.New()
	en_translations.RegisterDefaultTranslations(validate, ValidationTranslator)

	return validate
}

func ErrorHandler(ctx *gin.Context, err any) {
	switch value := err.(type) {
	case NotFoundError, InvariantError:
		appError := value.(ApplicationError)
		web.ErrorResponse(ctx, appError.GetStatusCode(), appError.GetErrorCode(), appError.GetErrorName(), appError)
	case validator.ValidationErrors:
		validationMessage := value.Translate(ValidationTranslator)
		web.ValidationErrorResponse(ctx, validationMessage)
	default:
		internalError := value.(error)
		web.ErrorResponse(ctx, 500, 500, "INTERNAL_SERVER_ERROR", internalError)
	}
}
