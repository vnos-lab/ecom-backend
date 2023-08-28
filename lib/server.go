package lib

import (
	"context"
	"ecom/api/middlewares"
	config "ecom/config"
	constants "ecom/constants"
	"ecom/infrastructure/db"
	"ecom/utils"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"go.uber.org/fx"
)

func NewServerGroup(instance *gin.Engine) *gin.RouterGroup {
	return instance.Group("/api")
}

func NewServer(lifecycle fx.Lifecycle, zap *zap.Logger, config *config.Config, db *db.Database, middlewares *middlewares.GinMiddleware) *gin.Engine {
	switch config.Server.Env {
	case constants.Dev, constants.Local:
		gin.SetMode(gin.DebugMode)
	case constants.Prod:
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	//gin.LoggerWithConfig(gin.LoggerConfig{
	//	Formatter: nil,
	//	Output:    nil,
	//	SkipPaths: nil,
	//})
	instance := gin.New()

	//instance.Use(gozap.RecoveryWithZap(zap, true))

	instance.Use(middlewares.CORS)
	instance.Use(middlewares.JWT(config, db))
	instance.Use(middlewares.Logger(zap))
	instance.Use(middlewares.ErrorHandler(zap))
	instance.Use(middlewares.JSONMiddleware)

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			zap.Info("Starting HTTP server")

			SeedRoutes(instance, db)
			go func() {
				addr := fmt.Sprint(config.Server.Host, ":", config.Server.Port)
				if err := instance.Run(addr); err != nil {
					zap.Fatal(fmt.Sprint("HTTP server failed to start %w", err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			zap.Info("Stopping HTTP server")
			return nil
		},
	})

	return instance
}

func SeedRoutes(engine *gin.Engine, db *db.Database) error {
	// Delete all routes
	db.DB.MustExec("DELETE FROM routes")
	qb := utils.Psql().Insert("routes").Columns("method", "path")
	args := []interface{}{}
	for _, r := range engine.Routes() {
		qb = qb.Values(r.Method, r.Path)
		args = append(args, r.Method, r.Path)
	}
	query, _, _ := qb.ToSql()

	db.DB.MustExecContext(context.Background(), query, args...)
	return nil
}
