package helper

import (
	"app-ecommerce-server/data/dto"
	"app-ecommerce-server/data/entity"
)

func ToAuthResponseDTO(user *entity.User) *dto.UserResponseDTO {
	return &dto.UserResponseDTO{
		Id:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		SignupWith:  user.SignupWith,
		Role:        user.Role,
		Created:     user.CreatedAt,
		Updated:     user.UpdatedAt,
	}
}

func ToListAuthResponseDTO(user []*entity.User) []*dto.UserResponseDTO {
	var listTemp = []*dto.UserResponseDTO{}
	for _, data := range user {
		listTemp = append(listTemp, &dto.UserResponseDTO{
			Id:          data.ID,
			Username:    data.Username,
			Email:       data.Email,
			PhoneNumber: data.PhoneNumber,
			SignupWith:  data.SignupWith,
			Created:     data.CreatedAt,
			Updated:     data.UpdatedAt,
		})
	}
	return listTemp
}
