package repositories

import (
	"context"
	"database/sql"
	helpercontext "server/helpers/helper_context"
)

type TransactorRepositoryInterface interface {
	Atomic(ctx context.Context, fn func(context.Context) error) error
}

type TransactorRepositoryImplementation struct {
	db *sql.DB
}

func NewTransactorRepositoryImplementation(db *sql.DB) *TransactorRepositoryImplementation {
	return &TransactorRepositoryImplementation{
		db: db,
	}
}

func (t *TransactorRepositoryImplementation) Atomic(ctx context.Context, fn func(context.Context) error) error {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = fn(helpercontext.SetTx(ctx, tx))
	if err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			return err
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
