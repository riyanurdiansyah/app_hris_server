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

type AuthServiceImpl struct {
	AuthRepository repository.AuthRepository

	DB       *gorm.DB
	Validate *validator.Validate
}

func NewAuthService(authRepository repository.AuthRepository, DB *gorm.DB, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		AuthRepository: authRepository,
		DB:             DB,
		Validate:       validate,
	}

}

// SignUp implements AuthService
func (service *AuthServiceImpl) SignUp(request *dto.UserCreateDTO) *dto.UserResponseDTO {
	errorValidation := service.Validate.Struct(request)
	if errorValidation != nil {
		msgError := validation.TextValidation(errorValidation.Error())
		return &dto.UserResponseDTO{
			Error:   true,
			Message: msgError,
		}
	}
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	if tx.Error != nil {
		msgError := validation.TextValidation(tx.Error.Error())
		return &dto.UserResponseDTO{
			Error:   true,
			Message: msgError,
		}
	} else {
		user := entity.User{
			EmployeeId:       request.EmployeeId,
			Username:         request.Username,
			Email:            request.Email,
			Password:         request.Password,
			PhoneNumber:      request.PhoneNumber,
			Role:             request.Role,
			CompanySecretKey: request.CompanySecretKey,
			CreatedAt:        time.Now().Format(time.RFC3339),
			UpdatedAt:        time.Now().Format(time.RFC3339),
		}

		userResponse := service.AuthRepository.SignUp(tx, &user)

		return dto.ToAuthResponseDTO(userResponse)
	}
}

// FindUserByEmail implements AuthService
func (service *AuthServiceImpl) FindUserByEmail(request *dto.UserLoginEmailDTO) *dto.UserResponseDTO {
	errorValidation := service.Validate.Struct(request)
	if errorValidation != nil {
		msgError := validation.TextValidation(errorValidation.Error())
		return &dto.UserResponseDTO{
			Error:   true,
			Message: msgError,
		}
	}
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	if tx.Error != nil {
		msgError := validation.TextValidation(tx.Error.Error())
		return &dto.UserResponseDTO{
			Error:   true,
			Message: msgError,
		}
	} else {
		userResponse := service.AuthRepository.FindUserByEmail(tx, request.Email)
		return dto.ToAuthResponseDTO(userResponse)
	}
}

// FindUserByUsername implements AuthService
func (service *AuthServiceImpl) FindUserByUsername(request *dto.UserLoginUsernameDTO) *dto.UserResponseDTO {
	errorValidation := service.Validate.Struct(request)
	if errorValidation != nil {
		msgError := validation.TextValidation(errorValidation.Error())
		return &dto.UserResponseDTO{
			Error:   true,
			Message: msgError,
		}
	}
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	if tx.Error != nil {
		msgError := validation.TextValidation(tx.Error.Error())
		return &dto.UserResponseDTO{
			Error:   true,
			Message: msgError,
		}
	} else {
		userResponse := service.AuthRepository.FindUserByUsername(tx, request.Username)
		return dto.ToAuthResponseDTO(userResponse)
	}
}

// CheckEmail implements AuthService
func (service *AuthServiceImpl) CheckEmail(email string) bool {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	if tx.Error != nil {
		return true
	} else {
		result := service.AuthRepository.CheckEmail(tx, email)
		return result
	}
}

// CheckUsername implements AuthService
func (service *AuthServiceImpl) CheckUsername(username string) bool {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	if tx.Error != nil {
		return true
	} else {
		result := service.AuthRepository.CheckUsername(tx, username)
		return result
	}
}

// CheckCompany implements AuthService
func (service *AuthServiceImpl) CheckCompany(key string) bool {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	if tx.Error != nil {
		return true
	} else {
		result := service.AuthRepository.CheckCompany(tx, key)
		return result
	}
}
