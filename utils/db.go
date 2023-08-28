package utils

import (
	"ecom/infrastructure/db"

	"github.com/jackc/pgx/v4"
)

func ErrNoRows(err error) bool {
	return err == pgx.ErrNoRows
}

func MustHaveDb(db *db.Database) {
	if db == nil {
		panic("Database engine is null")
	}
}
