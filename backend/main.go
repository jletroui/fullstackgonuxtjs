package main

import (
	"backend/config"
	"backend/logic"
	"backend/web"
	"fmt"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConfigureLogging()

	cfg, err := config.LoadConfigFromEnv("dev")
	if err != nil {
		fatal(err, "Cannot load config")
	}

	err = config.MigrateDatabase(cfg)
	if err != nil {
		fatal(err, "Cannot migrate database %s @%s", cfg.PostgresDatabase, cfg.PostgresHost)
	}

	db, err := config.NewDatabase(cfg)
	if err != nil {
		fatal(err, "Cannot connect to database %s @%s", cfg.PostgresDatabase, cfg.PostgresHost)
	}

	err = config.ConfigureSuperTokens(cfg)
	if err != nil {
		fatal(err, "Cannot initialize SuperTokens client")
	}

	taskRepo := logic.NewPostgresTaskRepository(db)
	sessionVerifier := web.NewSuperTokenSessionVerifier()

	svr, err := config.CreateWebServer(cfg, func(apiRoutes *gin.RouterGroup) {
		web.InstallTaskRoutes(apiRoutes, sessionVerifier, taskRepo)
	})
	if err != nil {
		fatal(err, "Cannot initialize web server")
	}

	slog.Info("Setup done. Starting server.")

	err = svr.Run(fmt.Sprintf("0.0.0.0:%d", cfg.Port))
	if err != nil {
		fatal(err, "Cannot start server at port %d", cfg.Port)
	}
}

func fatal(err error, msgTemplate string, args ...any) {
	slog.Error(fmt.Sprintf(msgTemplate, args...), "error", err)
	os.Exit(1)
}
