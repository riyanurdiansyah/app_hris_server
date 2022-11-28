package controller

import (
	"github.com/gin-gonic/gin"
)

type UserController interface {
	AddUserInfoPersonal(c *gin.Context)
	UpdateUserInfoPersonal(c *gin.Context)
	CheckUser(email string) bool
}
