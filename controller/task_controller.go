package controller

import (
	"github.com/gin-gonic/gin"
)

type TaskController interface {
	InsertTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	GetTaskByUserId(c *gin.Context)
	CheckTask(idTask int) bool
}
