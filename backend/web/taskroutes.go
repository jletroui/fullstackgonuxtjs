package web

import (
	"backend/logic"

	"github.com/gin-gonic/gin"
)

func InstallTaskRoutes(svr *gin.RouterGroup, taskRepo logic.TaskRepository) {
	svr.GET("/tasks/count", func(ctx *gin.Context) {
		count, err := taskRepo.Count()
		if err != nil {
			ctx.AbortWithError(500, err)
		}

		ctx.JSON(200, gin.H{
			"count": count,
		})
	})
}
