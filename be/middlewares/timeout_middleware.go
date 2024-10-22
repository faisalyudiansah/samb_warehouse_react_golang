package middlewares

import (
	"context"
	"server/apperrors"
	"time"

	"github.com/gin-gonic/gin"
)

func RequestTimeout() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		timeoutCtx, cancel := context.WithTimeout(
			ctx.Request.Context(),
			time.Duration(5)*time.Second,
		)
		defer cancel()

		done := make(chan struct{})
		ctx.Request = ctx.Request.WithContext(timeoutCtx)

		go next(ctx, done)

		select {
		case <-timeoutCtx.Done():
			ctx.Error(apperrors.ErrTimeoutError)
			ctx.Abort()
		case <-done:
		}
	}
}

func next(ctx *gin.Context, done chan struct{}) {
	defer func() {
		close(done)
		if err, ok := recover().(error); ok && err != nil {
			ctx.Error(err)
			ctx.Abort()
		}
	}()

	ctx.Next()
}
