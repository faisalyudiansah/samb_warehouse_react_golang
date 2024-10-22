package repositories

import "database/sql"

type ItemRepositoryInterface interface {
}

type ItemRepositoryImplementation struct {
	db *sql.DB
}

func NewItemRepositoryImplementation(db *sql.DB) *ItemRepositoryImplementation {
	return &ItemRepositoryImplementation{
		db: db,
	}
}
