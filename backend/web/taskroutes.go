package web

import (
	"backend/logic"

	"github.com/gin-gonic/gin"
)

type NewTask struct {
	Description string `json:"description" binding:"required"`
}

type tasksController struct {
	taskRepo logic.TaskRepository
}

func InstallTaskRoutes(svr *gin.RouterGroup, sess SessionVerifier, taskRepo logic.TaskRepository) {
	controller := &tasksController{taskRepo}
	svr.GET("/tasks/count", controller.getTasksCount)
	svr.POST("/tasks", sess.VerifySession, controller.postNewTask)
}

func (c *tasksController) getTasksCount(ctx *gin.Context) {
	count, err := c.taskRepo.Count()
	if err != nil {
		_ = ctx.AbortWithError(500, err)
	}

	ctx.JSON(200, gin.H{
		"count": count,
	})
}

func (c *tasksController) postNewTask(ctx *gin.Context) {
	var body NewTask
	err := ctx.BindJSON(&body)
	if err != nil {
		return
	}
	err = c.taskRepo.CreateTask(body.Description)
	if err != nil {
		_ = ctx.AbortWithError(500, err)
	}

	ctx.String(200, "")
}
