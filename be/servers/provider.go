package servers

import (
	"database/sql"
	"server/controllers"
	"server/helpers"
	"server/helpers/logger"
	"server/repositories"
	"server/services"

	"github.com/gin-gonic/gin"
)

var (
	Bcrypt   *helpers.BcryptStruct
	Jwt      *helpers.JwtProviderHS256
	GetParam *helpers.GetParam
)

var (
	TransactorRepository *repositories.TransactorRepositoryImplementation
	ItemRepository       *repositories.ItemRepositoryImplementation
)

var (
	ItemService *services.ItemServiceImplementation
)

var (
	ItemController *controllers.ItemController
)

func SetupController(router *gin.Engine, db *sql.DB) {
	logrusLogger := logger.NewLogger()
	logger.SetLogger(logrusLogger)

	Bcrypt = helpers.NewBcryptStruct()
	Jwt = helpers.NewJWTProviderHS256()
	GetParam = helpers.NewGetParams()
	TransactorRepository = repositories.NewTransactorRepositoryImplementation(db)

	SetupItemController(db)
	SetupItemRoute(router, ItemController)
}

func SetupItemController(db *sql.DB) {
	ItemRepository = repositories.NewItemRepositoryImplementation(db)
	ItemService = services.NewItemServiceImplementation(ItemRepository, TransactorRepository)
	ItemController = controllers.NewItemController(ItemService)
}
