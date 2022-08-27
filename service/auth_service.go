package service

import (
	"app-ecommerce-server/data/dto"
)

type AuthService interface {
	SignUp(request *dto.UserCreateDTO) *dto.UserResponseDTO
	FindUserByEmail(request *dto.UserLoginEmailDTO) *dto.UserResponseDTO
	FindUserByUsername(request *dto.UserLoginUsernameDTO) *dto.UserResponseDTO
	CheckEmail(email string) bool
	CheckUsername(username string) bool
}
