package ginutils

import (
	"net/http"
	"reflect"
	"server/constants"
	"server/dtos"

	"github.com/gin-gonic/gin"
)

func ResponseOK[T any](ctx *gin.Context, data T) {
	ResponseJSON(ctx, http.StatusOK, constants.ResponseSuccessMessage, data, nil)
}

func ResponseOKPlain(ctx *gin.Context) {
	ResponseOK[any](ctx, nil)
}

func ResponseOKPagination[T any](ctx *gin.Context, data T, paging *dtos.PageMetaData) {
	ResponseJSON(ctx, http.StatusOK, constants.ResponseSuccessMessage, data, paging)
}

func ResponseOKSeekPagination[T any](ctx *gin.Context, data T, paging *dtos.SeekPageMetaData) {
	ResponseSeekJSON(ctx, http.StatusOK, constants.ResponseSuccessMessage, data, paging)
}

func ResponseCreated[T any](ctx *gin.Context, data T) {
	ResponseJSON(ctx, http.StatusCreated, constants.ResponseSuccessMessage, data, nil)
}

func ResponseCreatedPlain(ctx *gin.Context) {
	ResponseCreated[any](ctx, nil)
}

func ResponseNoContent(ctx *gin.Context) {
	ResponseJSON[any](ctx, http.StatusNoContent, constants.ResponseSuccessMessage, nil, nil)
}

func ResponseSeekJSON[T any](ctx *gin.Context, statusCode int, message string, data T, paging *dtos.SeekPageMetaData) {
	ctx.JSON(statusCode, dtos.WebResponseSeek[T]{
		Message: message,
		Data:    data,
		Paging:  paging,
	})
}

func ResponseJSON[T any](ctx *gin.Context, statusCode int, message string, data T, paging *dtos.PageMetaData) {
	v := reflect.ValueOf(data)

	if v.Kind() == reflect.Slice && v.IsNil() {
		ctx.JSON(statusCode, dtos.WebResponseDataEmptyArray[[]T]{
			Message: message,
			Data:    []T{},
			Paging:  paging,
		})
		return
	}

	if v.Kind() == reflect.Slice && v.Len() == 0 {
		ctx.JSON(statusCode, dtos.WebResponseDataEmptyArray[T]{
			Message: message,
			Data:    data,
			Paging:  paging,
		})
		return
	}

	ctx.JSON(statusCode, dtos.WebResponse[T]{
		Message: message,
		Data:    data,
		Paging:  paging,
	})
}
