package servers

import (
	"database/sql"
	"server/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func SetupServer(router *gin.Engine, db *sql.DB) {
	SetupMiddlewares(router)
	SetupController(router, db)
	router.NoRoute(InvalidRoute)
}

func SetupMiddlewares(router *gin.Engine) {
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
	router.Use(middlewares...)
}
