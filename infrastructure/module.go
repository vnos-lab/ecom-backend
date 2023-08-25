package infrastructure

import (
	"erp/infrastructure/cache"
	"erp/infrastructure/db"
	"erp/infrastructure/search"

	"go.uber.org/fx"
)

var Module = fx.Options(fx.Provide(db.NewDatabase), fx.Provide(search.NewSearchClient), fx.Provide(cache.NewCacheClient))
