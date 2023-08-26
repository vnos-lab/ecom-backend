package db

import (
	config "ecom/config"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Database struct {
	*sqlx.DB
	logger *zap.Logger
}

func NewDatabase(config *config.Config, logger *zap.Logger) *Database {
	var err error

	logger.Info("Connecting to database...")
	sqlxDb, err := getSqlxDb(config)
	if err != nil {
		logger.Error("[Postgres] - Init connection exception: %s", zap.Error(err))
		// retry 5 times
		for i := 0; i < 5; i++ {
			time.Sleep(5 * time.Second)
			sqlxDb, err = getSqlxDb(config)
			if err == nil {
				break
			}
		}
	}

	err = sqlxDb.Ping()
	if err != nil {
		logger.Error("[Postgres] - Ping connection exception: %s", zap.Error(err))
		// retry 5 times
		for i := 0; i < 5; i++ {
			time.Sleep(5 * time.Second)
			err = sqlxDb.Ping()
			if err == nil {
				break
			}
		}
	}
	logger.Info("Database connected")

	db := &Database{
		DB:     sqlxDb,
		logger: logger,
	}

	return db
}

func getSqlxDb(config *config.Config) (*sqlx.DB, error) {
	var err error
	var sqlxDB *sqlx.DB
	switch config.Database.Driver {
	case "mysql":
		break
	case "postgres":
		dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
			config.Database.Host, config.Database.Username, config.Database.Password, config.Database.Name, config.Database.Port, config.Database.SSLMode, config.Database.TimeZone)

		sqlxDB, err = sqlx.Open("pgx", dns)
		if err != nil {
			return nil, err
		}

		err = sqlxDB.Ping()
		if err != nil {
			return nil, err
		}

		sqlxDB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	}

	sqlxDB.SetMaxOpenConns(config.Database.MaxOpenConnections)
	sqlxDB.SetMaxIdleConns(config.Database.MaxIdleConnections)
	sqlxDB.SetConnMaxLifetime(time.Duration(config.Database.ConnMaxLifetime))
	return sqlxDB, nil
}
