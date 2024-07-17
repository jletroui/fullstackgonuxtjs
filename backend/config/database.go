package config

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
)

func Migrate(cfg *Config) error {
	migrator, err := migrate.New(
		"file://ops/db/migrations",
		fmt.Sprintf(
			"postgres://%s:%s@%s:5432/%s?sslmode=disable",
			cfg.PostgresAdminUser,
			cfg.PostgresAdminPassword,
			cfg.PostgresHost,
			cfg.PostgresDatabase,
		),
	)
	if err != nil {
		return err
	}
	return migrator.Up()
}
