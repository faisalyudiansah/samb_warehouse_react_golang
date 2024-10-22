package helpers

import (
	"server/dtos"
	"server/helpers/validationutils"

	"github.com/go-playground/validator/v10"
)

func FormatterErrorInput(ve validator.ValidationErrors) []dtos.ResponseApiError {
	result := make([]dtos.ResponseApiError, len(ve))
	for i, fe := range ve {
		result[i] = dtos.ResponseApiError{
			Field: jsonFieldName(fe.Field()),
			Msg:   validationutils.TagToMsg(fe),
		}
	}
	return result
}
