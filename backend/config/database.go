package config

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/url"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type migrateLogger struct {
}

func NewDatabase(cfg *Config) (*sql.DB, error) {
	return sql.Open(
		"postgres",
		fmt.Sprintf(
			"postgres://%s:%s@%s:5432/%s?sslmode=disable",
			cfg.PostgresAdminUser,
			url.QueryEscape(cfg.PostgresAdminPassword),
			cfg.PostgresHost,
			cfg.PostgresDatabase,
		),
	)
}

func MigrateDatabase(cfg *Config) error {
	migrator, err := migrate.New(
		fmt.Sprintf(
			"file://%sops/db/migrations",
			cfg.BasePath,
		),
		fmt.Sprintf(
			"postgres://%s:%s@%s:5432/%s?sslmode=disable",
			cfg.PostgresAdminUser,
			url.QueryEscape(cfg.PostgresAdminPassword),
			cfg.PostgresHost,
			cfg.PostgresDatabase,
		),
	)
	if err != nil {
		return err
	}
	migrator.Log = &migrateLogger{}
	err = migrator.Up()
	if errors.Is(err, migrate.ErrNoChange) {
		log.Printf("No database migrations to apply.")
	} else if err != nil {
		return err
	}
	return nil
}

func (*migrateLogger) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func (*migrateLogger) Verbose() bool {
	return true
}
