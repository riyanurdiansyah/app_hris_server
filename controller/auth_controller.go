package controller

import (
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	SignUp(c *gin.Context)
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}
