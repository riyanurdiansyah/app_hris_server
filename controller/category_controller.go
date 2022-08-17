package controller

import "github.com/gin-gonic/gin"

type CategoryController interface {
	InsertCategory(c *gin.Context)
	FindAllCategory(c *gin.Context)
	FindByIdCategory(c *gin.Context)
	DeleteCategory(c *gin.Context)
	UpdateCategory(c *gin.Context)
}
