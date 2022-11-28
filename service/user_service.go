package service

import (
	"app-hris-server/data/dto"
)

type UserService interface {
	AddUserInfoPersonal(request *dto.UserInfoCreateDTO) *dto.UserPersonalInfoResponseDTO
	UpdateUserInfoPersonal(request *dto.UserInfoCreateDTO) *dto.UserPersonalInfoResponseDTO
	CheckUser(userId int) bool
}
