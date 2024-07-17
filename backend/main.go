package main

import (
	"backend/config"
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

	err = config.RunServer(cfg, func(apiRoutes *gin.RouterGroup) {
		web.InstallTaskRoutes(apiRoutes)
	})

	if err != nil {
		log.Fatalf("Cannot start server at port %d: %s", cfg.Port, err)
	}
}
