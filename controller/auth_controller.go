package controller

import (
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	SignUp(c *gin.Context)
	SigninWithUsername(c *gin.Context)
	FindUserByEmail(c *gin.Context)
	FindUserByUsername(c *gin.Context)
	CheckEmail(email string) bool
	CheckUsername(username string) bool
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}
