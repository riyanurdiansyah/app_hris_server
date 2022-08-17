package service

import (
	"app-ecommerce-server/data/dto"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	SignUp(ctx *gin.Context, request *dto.UserCreateDTO) *dto.UserResponseDTO
}
