package servers

import (
	"net/http"
	"server/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoute(h *HandlerOps) *gin.Engine {
	g := gin.New()
	g.ContextWithFallback = true
	g.Use(gin.Recovery(), middlewares.LoggerMiddleware(), middlewares.ErrorHandler)

	g.NoRoute(InvalidRoute)

	SetupYourRoute(g, h)

	return g
}

func SetupYourRoute(g *gin.Engine, h *HandlerOps) {
	u := g.Group("/")
	u.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})
}
