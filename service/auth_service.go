package service

import (
	"app-ecommerce-server/data/dto"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	SignUp(ctx *gin.Context, request *dto.UserCreateDTO) *dto.UserResponseDTO
	FindUserByEmail(ctx *gin.Context, request *dto.UserLoginDTO) *dto.UserResponseDTO
	FindUserByUsername(ctx *gin.Context, request *dto.UserLoginDTO) *dto.UserResponseDTO
	CheckEmail(ctx *gin.Context, email string) bool
	CheckUsername(ctx *gin.Context, username string) bool
}
