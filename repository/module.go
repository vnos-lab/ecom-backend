package repository

import (
	"erp/repository/postgres"

	"go.uber.org/fx"
)

var Module = fx.Provide(
	postgres.NewUserRepository,
)
