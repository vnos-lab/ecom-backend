package infrastructure

import (
	"ecom/infrastructure/cache"
	"ecom/infrastructure/db"
	"ecom/infrastructure/search"

	"go.uber.org/fx"
)

var Module = fx.Options(fx.Provide(db.NewDatabase), fx.Provide(search.NewSearchClient), fx.Provide(cache.NewCacheClient))
