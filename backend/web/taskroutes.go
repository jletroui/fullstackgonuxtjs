package web

import "github.com/gin-gonic/gin"

func InstallTaskRoutes(svr *gin.RouterGroup) {
	svr.GET("/tasks/count", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"count": 5,
		})
	})
}
