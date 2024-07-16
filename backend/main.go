package main

import (
	"backend/config"
	"backend/web"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadFromEnv("dev")
	if err != nil {
		log.Fatalf("Cannot load config: %s", err)
	}

	err = config.RunServer(cfg, func(svr *gin.RouterGroup) {
		web.InstallTaskRoutes(svr)
	})

	if err != nil {
		log.Fatalf("Cannot start server at port %d: %s", cfg.Port, err)
	}
}
