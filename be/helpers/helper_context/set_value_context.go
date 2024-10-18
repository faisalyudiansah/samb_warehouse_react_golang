package helpercontext

import (
	"context"
	"database/sql"
	"server/constants"
)

func SetTx(c context.Context, tx *sql.Tx) context.Context {
	var ctx constants.Ctx = "ctx"
	return context.WithValue(c, ctx, tx)
}
