package dto

import (
	"app-travel-server/data/entity"
)

func ToAuthResponseDTO(user *entity.User) *UserResponseDTO {
	return &UserResponseDTO{
		Id:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Password:    user.Password,
		SignupWith:  user.SignupWith,
		Role:        user.Role,
		Created:     user.CreatedAt,
		Updated:     user.UpdatedAt,
	}
}

func ToListAuthResponseDTO(user []*entity.User) []*UserResponseDTO {
	var listTemp = []*UserResponseDTO{}
	for _, data := range user {
		listTemp = append(listTemp, &UserResponseDTO{
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
