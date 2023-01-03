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

type ConfigServiceImpl struct {
	ConfigRepository repository.ConfigRepository

	DB       *gorm.DB
	Validate *validator.Validate
}

func NewConfigService(repo repository.ConfigRepository, DB *gorm.DB, validate *validator.Validate) ConfigService {
	return &ConfigServiceImpl{
		ConfigRepository: repo,
		DB:               DB,
		Validate:         validate,
	}
}

// CheckConfig implements ConfigService
func (service *ConfigServiceImpl) CheckConfig(name string) bool {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	if tx.Error != nil {
		return true
	} else {
		result := service.ConfigRepository.CheckNameConfig(tx, name)
		return result
	}
}

// GetConfig implements ConfigService
func (service *ConfigServiceImpl) GetConfig() []*dto.ConfigResponseDTO {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	if tx.Error != nil {
		return []*dto.ConfigResponseDTO{}
	} else {
		configs := service.ConfigRepository.GetConfig(tx)
		return dto.ToListConfigResponseDTO(configs)
	}
}

// GetConfigByName implements ConfigService
func (service *ConfigServiceImpl) GetConfigByName(name string) *dto.ConfigResponseDTO {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	if tx.Error != nil {
		return &dto.ConfigResponseDTO{}
	} else {
		menu := service.ConfigRepository.GetConfigByName(tx, name)
		return dto.ToConfigResponseDTO(menu)
	}
}

// InsertConfig implements ConfigService
func (service *ConfigServiceImpl) InsertConfig(request *dto.ConfigCreateDTO) *dto.ConfigResponseDTO {
	errorValidation := service.Validate.Struct(request)
	if errorValidation != nil {
		msgError := validation.TextValidation(errorValidation.Error())
		return &dto.ConfigResponseDTO{
			Error:   true,
			Message: msgError,
		}
	}
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	if tx.Error != nil {
		msgError := validation.TextValidation(tx.Error.Error())
		return &dto.ConfigResponseDTO{
			Error:   true,
			Message: msgError,
		}
	} else {
		config := entity.Config{
			Name:      request.Name,
			Value:     request.Value,
			CreatedAt: time.Now().Format(time.RFC3339),
			UpdatedAt: time.Now().Format(time.RFC3339),
		}

		response := service.ConfigRepository.InsertConfig(tx, &config)

		return dto.ToConfigResponseDTO(response)
	}
}

// UpdateConfig implements ConfigService
func (service *ConfigServiceImpl) UpdateConfig(request *dto.ConfigUpdateDTO) *dto.ConfigResponseDTO {
	errorValidation := service.Validate.Struct(request)
	if errorValidation != nil {
		msgError := validation.TextValidation(errorValidation.Error())
		return &dto.ConfigResponseDTO{
			Error:   true,
			Message: msgError,
		}
	}
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	if tx.Error != nil {
		msgError := validation.TextValidation(tx.Error.Error())
		return &dto.ConfigResponseDTO{
			Error:   true,
			Message: msgError,
		}
	} else {
		config := entity.Config{
			ID:        request.ID,
			Name:      request.Name,
			Value:     request.Value,
			CreatedAt: time.Now().Format(time.RFC3339),
			UpdatedAt: time.Now().Format(time.RFC3339),
		}

		response := service.ConfigRepository.UpdateConfig(tx, &config)

		return dto.ToConfigResponseDTO(response)
	}
}
