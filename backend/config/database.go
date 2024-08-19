package config

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net"
	"net/url"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

const (
	DEFAULT_PORT = 5432
	WAIT_MS      = 100
)

type migrateLogger struct {
}

func waitDbReady(cfg *Config) error {
	maxTries := cfg.PostgresWaitTimeoutMs / WAIT_MS
	isReady := isDbReady(cfg)

	if !isReady {
		slog.Info(fmt.Sprintf("Postgres seems not ready. Will retry to connect %d times.", maxTries))
	}

	for maxTries > 0 && !isReady {
		time.Sleep(WAIT_MS * time.Millisecond)
		isReady = isDbReady(cfg)
		maxTries--
	}

	if isReady {
		slog.Info("Postgres is ready. Proceeding.")
		return nil
	}
	return fmt.Errorf("can't connect to %s:%d after %d milliseconds", cfg.PostgresHost, DEFAULT_PORT, cfg.PostgresWaitTimeoutMs)
}

func isDbReady(cfg *Config) bool {
	_, err := net.Dial("tcp", fmt.Sprintf("%s:%d", cfg.PostgresHost, DEFAULT_PORT))
	return err == nil
}

func NewDatabase(cfg *Config) (*sql.DB, error) {
	return sql.Open(
		"postgres",
		fmt.Sprintf(
			"postgres://%s:%s@%s:%d/%s?sslmode=disable",
			cfg.PostgresAdminUser,
			url.QueryEscape(cfg.PostgresAdminPassword),
			cfg.PostgresHost,
			DEFAULT_PORT,
			cfg.PostgresDatabase,
		),
	)
}

func MigrateDatabase(cfg *Config) error {
	err := waitDbReady(cfg)
	if err != nil {
		return err
	}

	migrator, err := migrate.New(
		fmt.Sprintf(
			"file://%sops/db/migrations",
			cfg.BasePath,
		),
		fmt.Sprintf(
			"postgres://%s:%s@%s:%d/%s?sslmode=disable",
			cfg.PostgresAdminUser,
			url.QueryEscape(cfg.PostgresAdminPassword),
			cfg.PostgresHost,
			DEFAULT_PORT,
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
