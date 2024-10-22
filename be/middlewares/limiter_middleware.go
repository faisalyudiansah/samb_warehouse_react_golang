package middlewares

import (
	"server/apperrors"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func RateLimiter(limiter *rate.Limiter) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !limiter.Allow() {
			ctx.Error(apperrors.ErrLimitError)
			return
		}

		ctx.Next()
	}
}
