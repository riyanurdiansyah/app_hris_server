package service

import (
	"app-hris-server/data/dto"
)

type AuthService interface {
	SignUp(request *dto.UserCreateDTO) *dto.UserResponseDTO
	FindUserByEmail(request *dto.UserLoginEmailDTO) *dto.UserResponseDTO
	FindUserByUsername(request *dto.UserLoginUsernameDTO) *dto.UserResponseDTO
	CheckEmail(email string) bool
	CheckUsername(username string) bool
	CheckCompany(key string) bool
}
