package repositories

import (
	"context"
	"database/sql"
	"fmt"
	helpercontext "server/helpers/helper_context"
)

type TransactorRepositoryInterface interface {
	Atomic(c context.Context, fn func(context.Context) (any, error)) (any, error)
}

type TransactorRepositoryImplementation struct {
	db *sql.DB
}

func NewTransactorRepositoryImplementation(db *sql.DB) *TransactorRepositoryImplementation {
	return &TransactorRepositoryImplementation{
		db: db,
	}
}

func (dc *TransactorRepositoryImplementation) Atomic(c context.Context, fn func(context.Context) (any, error)) (any, error) {
	tx, err := dc.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				err = fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
			}
		}
		err = tx.Commit()
	}()

	result, err := fn(helpercontext.SetTx(c, tx))
	if err != nil {
		return nil, err
	}
	return result, nil
}
