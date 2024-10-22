package servers

import (
	"net/http"
	"server/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func SetupRoute(h *HandlerOps) *gin.Engine {
	g := gin.New()
	g.ContextWithFallback = true
	// g.Use(gin.Recovery(), middlewares.LoggerMiddleware(), middlewares.ErrorHandler)
	limiter := rate.NewLimiter(rate.Limit(50), 50)

	middlewares := []gin.HandlerFunc{
		middlewares.LoggerMiddleware(),
		middlewares.ErrorHandler,
		middlewares.RateLimiter(limiter),
		middlewares.RequestTimeout(),
		cors.New(cors.Config{
			AllowMethods:     []string{"*"},
			AllowHeaders:     []string{"*", "Authorization", "Content-Type"},
			AllowOrigins:     []string{"http://localhost:3000", "http://localhost:3001", "http://localhost:5173", "http://localhost:5174"},
			AllowCredentials: true,
		}),
		gin.Recovery(),
	}
	RegisterValidators()
	g.Use(middlewares...)

	g.NoRoute(InvalidRoute)

	SetupYourRoute(g, h)

	return g
}

func SetupYourRoute(g *gin.Engine, h *HandlerOps) {
	u := g.Group("/")
	u.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})

	// u.POST("/penerimaan-barang", controllers.CreatePenerimaanBarang)
	// u.POST("/pengeluaran-barang", controllers.CreatePengeluaranBarang)
}
