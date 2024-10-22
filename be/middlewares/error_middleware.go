package middlewares

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"server/apperrors"
	"server/constants"
	"server/dtos"
	"server/helpers"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler(c *gin.Context) {
	c.Next()
	if len(c.Errors) == 0 {
		return
	}
	if len(c.Errors) > 0 {
		err := c.Errors.Last()
		switch e := err.Err.(type) {
		case validator.ValidationErrors:
			var ve validator.ValidationErrors
			if errors.As(c.Errors[0].Err, &ve) {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": helpers.FormatterErrorInput(ve)})
				return
			}
		case *json.SyntaxError:
			handleJsonSyntaxError(c)
		case *json.UnmarshalTypeError:
			handleJsonUnmarshalTypeError(c, e)
		case *time.ParseError:
			handleParseTimeError(c, e)
		default:
			if errors.Is(e, io.EOF) {
				c.AbortWithStatusJSON(http.StatusBadRequest, dtos.WebResponse[any]{
					Message: constants.RequestBodyInvalid,
				})
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
}

func handleJsonSyntaxError(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.WebResponse[any]{
		Message: constants.JsonSyntaxErrorMessage,
	})
}

func handleJsonUnmarshalTypeError(ctx *gin.Context, err *json.UnmarshalTypeError) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.WebResponse[any]{
		Message: fmt.Sprintf(constants.InvalidJsonValueTypeErrorMessage, err.Field),
	})
}

func handleParseTimeError(ctx *gin.Context, err *time.ParseError) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, dtos.WebResponse[any]{
		Message: fmt.Sprintf("please send time in format of %s, got: %s", constants.ConvertGoTimeLayoutToReadable(err.Layout), err.Value),
	})
}
