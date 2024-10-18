package middlewares

import (
	"errors"
	"net/http"
	"server/apperrors"
	"server/helpers"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler(c *gin.Context) {
	c.Next()
	if len(c.Errors) == 0 {
		return
	}
	if len(c.Errors) > 0 {
		var ve validator.ValidationErrors
		if errors.As(c.Errors[0].Err, &ve) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": helpers.FormatterErrorInput(ve)})
			return
		}

		errorMappings := map[error]int{
			apperrors.ErrISE:                http.StatusInternalServerError,
			apperrors.ErrUnauthorization:    http.StatusUnauthorized,
			apperrors.ErrUrlNotFound:        http.StatusNotFound,
			apperrors.ErrRequestBodyInvalid: http.StatusBadRequest,
		}
		if statusCode, exists := errorMappings[c.Errors[0].Err]; exists {
			helpers.PrintError(c, statusCode, c.Errors[0].Err.Error())
			return
		}
		helpers.PrintError(c, http.StatusInternalServerError, apperrors.ErrISE.Error())
	}
}
