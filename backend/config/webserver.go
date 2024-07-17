package config

import (
	"strings"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func CreateRouter(cfg *Config, installApiRoutes func(*gin.RouterGroup)) *gin.Engine {
	router := gin.Default()

	apiRoutes := router.Group("/api")
	installApiRoutes(apiRoutes)

	router.Use(createStaticHandler())

	return router
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
