package main

import (
	"ecom/config"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	config := config.NewConfig()

	// "postgres://postgres:postgres@localhost:5432/example?sslmode=disable"
	dns := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name, config.Database.SSLMode)
	m, err := migrate.New(
		"file://migrations",
		dns)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}
