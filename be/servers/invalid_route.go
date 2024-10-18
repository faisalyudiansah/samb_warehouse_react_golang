package servers

import (
	"server/apperrors"

	"github.com/gin-gonic/gin"
)

func InvalidRoute(c *gin.Context) {
	c.Error(apperrors.ErrUrlNotFound)
	c.Abort()
}
