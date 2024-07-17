package main

import (
	"backend/config"
	"backend/logic"
	"backend/web"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfigFromEnv("dev")
	if err != nil {
		log.Fatalf("Cannot load config: %s", err)
	}

	err = config.MigrateDatabase(cfg)
	if err != nil {
		log.Fatalf("Cannot migrate database %s @%s: %s", cfg.PostgresDatabase, cfg.PostgresHost, err)
	}
	db, err := config.NewDatabase(cfg)
	if err != nil {
		log.Fatalf("Cannot connect to database %s @%s: %s", cfg.PostgresDatabase, cfg.PostgresHost, err)
	}

	taskRepo := logic.NewPostgresTaskRepository(db)

	err = config.RunServer(cfg, func(apiRoutes *gin.RouterGroup) {
		web.InstallTaskRoutes(apiRoutes, taskRepo)
	})

	if err != nil {
		log.Fatalf("Cannot start server at port %d: %s", cfg.Port, err)
	}
}
