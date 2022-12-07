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

type AbsentServiceImpl struct {
	AbsentRepository repository.AbsentRepository

	DB       *gorm.DB
	Validate *validator.Validate
}

func NewAbsentRepository(repository repository.AbsentRepository, db *gorm.DB, validate *validator.Validate) AbsentService {
	return &AbsentServiceImpl{
		AbsentRepository: repository,
		DB:               db,
		Validate:         validate,
	}
}

// CheckIfDoneAbsent implements AbsentService
func (service *AbsentServiceImpl) CheckIfDoneAbsent(request *dto.AbsentCreateDTO) bool {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	if tx.Error != nil {
		return true
	} else {
		result := service.AbsentRepository.CheckIfDoneAbsent(tx, request)
		return result
	}
}

// InsertAbsent implements AbsentService
func (service *AbsentServiceImpl) InsertAbsent(request *dto.AbsentCreateDTO) *dto.AbsentResponseDTO {
	errorValidation := service.Validate.Struct(request)
	if errorValidation != nil {
		msgError := validation.TextValidation(errorValidation.Error())
		return &dto.AbsentResponseDTO{
			Error:   true,
			Message: msgError,
		}
	}
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	if tx.Error != nil {
		msgError := validation.TextValidation(tx.Error.Error())
		return &dto.AbsentResponseDTO{
			Error:   true,
			Message: msgError,
		}
	} else {
		absent := entity.Absent{
			IdUser:     request.IdUser,
			IdEmployee: request.IdEmployee,
			Tanggal:    request.Tanggal,
			Latitude:   request.Latitude,
			Longitude:  request.Longitude,
			Catatan:    request.Catatan,
			Tipe:       request.Tipe,
			Photo:      request.Photo,
			CreatedAt:  time.Now().Format(time.RFC3339),
			UpdatedAt:  time.Now().Format(time.RFC3339),
		}

		absentResponse := service.AbsentRepository.InsertAbsent(tx, &absent)

		return dto.ToAbsentResponseDTO(absentResponse)
	}
}
