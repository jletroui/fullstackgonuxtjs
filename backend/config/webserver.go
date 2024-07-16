package config

import (
	"fmt"
	"strings"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func RunServer(config *Config, installApiRoutes func(*gin.RouterGroup)) error {
	svr := gin.Default()

	apiRoutes := svr.Group("/api")
	installApiRoutes(apiRoutes)

	svr.Use(createStaticHandler())

	return svr.Run(fmt.Sprintf("0.0.0.0:%d", config.Port))
}

func createStaticHandler() gin.HandlerFunc {
	innerHandler := static.Serve("/", static.LocalFile("www/", true))

	return func(ctx *gin.Context) {
		// We only want to pay the price of looking on the file system for non-API requests
		if !strings.HasPrefix(ctx.Request.URL.Path, "/api") {
			innerHandler(ctx)
		}
	}
}
