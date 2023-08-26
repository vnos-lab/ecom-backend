package utils

import "github.com/jackc/pgx/v4"

func ErrNoRows(err error) bool {
	return err == pgx.ErrNoRows
}
