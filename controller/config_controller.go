package controller

import (
	"github.com/gin-gonic/gin"
)

type ConfigController interface {
	InsertConfig(c *gin.Context)
	UpdateConfig(c *gin.Context)
	GetConfig(c *gin.Context)
	GetConfigByName(c *gin.Context)
	CheckConfig(name string) bool
}
