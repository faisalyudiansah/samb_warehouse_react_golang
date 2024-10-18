package servers

import (
	"database/sql"
	"fmt"
	"server/helpers"
	"server/helpers/logger"
)

type HandlerOps struct {
}

func SetupController(db *sql.DB) *HandlerOps {
	logrusLogger := logger.NewLogger()
	logger.SetLogger(logrusLogger)

	bcrypt := helpers.NewBcryptStruct()
	jwt := helpers.NewJWTProviderHS256()
	validationReqBody := helpers.NewValidationReqBody()
	getParam := helpers.NewGetParams()

	fmt.Println(bcrypt, jwt, validationReqBody, getParam)

	return &HandlerOps{}
}
