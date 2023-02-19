package dto

import (
	"app-hris-server/data/entity"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}

func ToAuthResponseDTO(user *entity.User) *UserResponseDTO {

	return &UserResponseDTO{
		// Id:               user.ID,
		Uuid:             user.Uuid,
		EmployeeId:       user.EmployeeId,
		Username:         user.Username,
		Email:            user.Email,
		PhoneNumber:      user.PhoneNumber,
		Password:         user.Password,
		Role:             user.Role,
		CompanySecretKey: user.CompanySecretKey,
		Created:          user.CreatedAt,
		Updated:          user.UpdatedAt,
	}
}

func ToListAuthResponseDTO(user []*entity.User) []*UserResponseDTO {
	var listTemp = []*UserResponseDTO{}
	for _, data := range user {
		listTemp = append(listTemp, &UserResponseDTO{
			// Id:          data.ID,
			Uuid:        data.Uuid,
			Username:    data.Username,
			Email:       data.Email,
			PhoneNumber: data.PhoneNumber,
			Created:     data.CreatedAt,
			Updated:     data.UpdatedAt,
		})
	}
	return listTemp
}
