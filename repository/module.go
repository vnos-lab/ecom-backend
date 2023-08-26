package repository

import (
	"ecom/repository/postgres"

	"go.uber.org/fx"
)

var Module = fx.Provide(
	postgres.NewUserRepository,
)
