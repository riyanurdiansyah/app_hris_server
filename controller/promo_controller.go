package controller

import "github.com/gin-gonic/gin"

type PromoController interface {
	InsertPromo(c *gin.Context)
	GetAllPromo(c *gin.Context)
	UpdatePromo(c *gin.Context)
	DeletePromo(c *gin.Context)
}
