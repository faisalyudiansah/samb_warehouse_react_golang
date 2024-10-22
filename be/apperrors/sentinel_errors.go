package apperrors

import (
	"errors"
	"server/constants"
)

var (
	ErrISE                = errors.New(constants.ISE)
	ErrUnauthorization    = errors.New(constants.Unauthorization)
	ErrUrlNotFound        = errors.New(constants.UrlNotFound)
	ErrRequestBodyInvalid = errors.New(constants.RequestBodyInvalid)
	ErrLimitError         = errors.New(constants.TooManyRequestsErrorMessage)
	ErrTimeoutError       = errors.New(constants.RequestTimeoutErrorMessage)
)
