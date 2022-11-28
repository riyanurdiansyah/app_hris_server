package service

import (
	"app-hris-server/data/dto"
	"app-hris-server/data/entity"
	"app-hris-server/helper"
	"app-hris-server/repository"
	"app-hris-server/validation"
	"time"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository

	DB       *gorm.DB
	Validate *validator.Validate
}

func NewUserService(repo repository.UserRepository, DB *gorm.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: repo,
		DB:             DB,
		Validate:       validate,
	}
}

// AddUserInfoPersonal implements UserService
func (service *UserServiceImpl) AddUserInfoPersonal(request *dto.UserInfoCreateDTO) *dto.UserPersonalInfoResponseDTO {
	errorValidation := service.Validate.Struct(request)
	if errorValidation != nil {
		msgError := validation.TextValidation(errorValidation.Error())
		return &dto.UserPersonalInfoResponseDTO{
			Error:   true,
			Message: msgError,
		}
	}
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	if tx.Error != nil {
		msgError := validation.TextValidation(tx.Error.Error())
		return &dto.UserPersonalInfoResponseDTO{
			Error:   true,
			Message: msgError,
		}
	} else {
		user := entity.UserPersonalInfo{
			IdEmployee:        request.IdEmployee,
			NamaDepan:         request.NamaDepan,
			NamaBelakang:      request.NamaBelakang,
			JenisKelamin:      request.JenisKelamin,
			TempatLahir:       request.TempatLahir,
			TanggalLahir:      request.TanggalLahir,
			NoHP:              request.NoHP,
			IdUser:            request.IdUser,
			Telepon:           request.Telepon,
			StatusPernikahan:  request.StatusPernikahan,
			Agama:             request.Agama,
			NomorId:           request.NomorId,
			TipeId:            request.TipeId,
			TanggalKadaluarsa: request.TanggalKadaluarsa,
			AlamatKTP:         request.AlamatKTP,
			AlamatDomisili:    request.AlamatDomisili,
			GolonganDarah:     request.GolonganDarah,
			CreatedAt:         time.Now().Format(time.RFC3339),
			UpdatedAt:         time.Now().Format(time.RFC3339),
		}

		response := service.UserRepository.AddUserInfoPersonal(tx, &user)

		return dto.ToUserPersonalInfoResponseDTO(response)
	}
}

// CheckUser implements UserService
func (service *UserServiceImpl) CheckUser(employeeId string) bool {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	if tx.Error != nil {
		return true
	} else {
		result := service.UserRepository.CheckUser(tx, employeeId)
		return result
	}
}

// UpdateUserInfoPersonal implements UserService
func (service *UserServiceImpl) UpdateUserInfoPersonal(request *dto.UserInfoCreateDTO) *dto.UserPersonalInfoResponseDTO {
	errorValidation := service.Validate.Struct(request)
	if errorValidation != nil {
		msgError := validation.TextValidation(errorValidation.Error())
		return &dto.UserPersonalInfoResponseDTO{
			Error:   true,
			Message: msgError,
		}
	}
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	if tx.Error != nil {
		msgError := validation.TextValidation(tx.Error.Error())
		return &dto.UserPersonalInfoResponseDTO{
			Error:   true,
			Message: msgError,
		}
	} else {
		user := entity.UserPersonalInfo{
			IdEmployee:        request.IdEmployee,
			NamaDepan:         request.NamaDepan,
			NamaBelakang:      request.NamaBelakang,
			JenisKelamin:      request.JenisKelamin,
			TempatLahir:       request.TempatLahir,
			TanggalLahir:      request.TanggalLahir,
			NoHP:              request.NoHP,
			IdUser:            request.IdUser,
			Telepon:           request.Telepon,
			StatusPernikahan:  request.StatusPernikahan,
			Agama:             request.Agama,
			NomorId:           request.NomorId,
			TipeId:            request.TipeId,
			TanggalKadaluarsa: request.TanggalKadaluarsa,
			AlamatKTP:         request.AlamatKTP,
			AlamatDomisili:    request.AlamatDomisili,
			GolonganDarah:     request.GolonganDarah,
			UpdatedAt:         time.Now().Format(time.RFC3339),
		}

		response := service.UserRepository.UpdateUserInfoPersonal(tx, &user)

		return dto.ToUserPersonalInfoResponseDTO(response)
	}
}
