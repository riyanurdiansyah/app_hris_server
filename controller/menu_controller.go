package controller

import (
	"github.com/gin-gonic/gin"
)

type MenuController interface {
	InsertMenu(c *gin.Context)
	UpdateMenu(c *gin.Context)
	CheckMenu(id int) bool
	GetMenu(c *gin.Context)
	GetMenuById(c *gin.Context)
}
