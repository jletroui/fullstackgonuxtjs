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

	taskRepo := logic.NewPostgresTaskRepository(db)

	svr := config.CreateWebServer(cfg, func(apiRoutes *gin.RouterGroup) {
		web.InstallTaskRoutes(apiRoutes, taskRepo)
	})

	err = svr.Run(fmt.Sprintf("0.0.0.0:%d", cfg.Port))
	if err != nil {
		fatal(err, "Cannot start server at port %d", cfg.Port)
	}
}

func fatal(err error, msgTemplate string, args ...any) {
	slog.Error(fmt.Sprintf(msgTemplate, args...), "error", err)
	os.Exit(1)
}
