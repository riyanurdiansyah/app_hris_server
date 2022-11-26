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

type CompanyServiceImpl struct {
	CompanyRepository repository.CompanyRepository

	DB       *gorm.DB
	Validate *validator.Validate
}

func NewCompanyService(repo repository.CompanyRepository, DB *gorm.DB, validate *validator.Validate) CompanyService {
	return &CompanyServiceImpl{
		CompanyRepository: repo,
		DB:                DB,
		Validate:          validate,
	}
}

// InsertCompany implements CompanyService
func (service *CompanyServiceImpl) InsertCompany(request *dto.CompanyCreateDTO) *dto.CompanyResponseDTO {
	errorValidation := service.Validate.Struct(request)
	if errorValidation != nil {
		msgError := validation.TextValidation(errorValidation.Error())
		return &dto.CompanyResponseDTO{
			Error:   true,
			Message: msgError,
		}
	}
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	if tx.Error != nil {
		msgError := validation.TextValidation(tx.Error.Error())
		return &dto.CompanyResponseDTO{
			Error:   true,
			Message: msgError,
		}
	} else {
		user := entity.Company{
			Name:      request.Name,
			SecretKey: request.SecretKey,
			CreatedAt: time.Now().Format(time.RFC3339),
			UpdatedAt: time.Now().Format(time.RFC3339),
		}

		response := service.CompanyRepository.InsertCompany(tx, &user)

		return dto.ToCompanyResponseDTO(response)
	}
}
