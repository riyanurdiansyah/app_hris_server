package controller

import (
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	SignUp(c *gin.Context)
	SigninWithUsername(c *gin.Context)
	SigninWithEmail(c *gin.Context)
	CheckEmail(email string) bool
	CheckUsername(username string) bool
	CheckCompany(key string) bool
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}
