package utils

import (
	sq "github.com/Masterminds/squirrel"
)

func Psql() sq.StatementBuilderType {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
}
