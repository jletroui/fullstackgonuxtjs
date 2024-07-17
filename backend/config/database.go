package config

import (
	"errors"
	"fmt"
	"log"
	"net/url"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type migrateLogger struct {
}

func MigrateDatabase(cfg *Config) error {
	migrator, err := migrate.New(
		"file://ops/db/migrations",
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
