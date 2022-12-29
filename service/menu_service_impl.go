package service

import (
	"app-hris-server/data/dto"
	"app-hris-server/data/entity"
	"app-hris-server/helper"
	"app-hris-server/repository"
	"app-hris-server/validation"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type MenuServiceImpl struct {
	MenuRepository repository.MenuRepository

	DB       *gorm.DB
	Validate *validator.Validate
}

func NewMenuService(repo repository.MenuRepository, DB *gorm.DB, validate *validator.Validate) MenuService {
	return &MenuServiceImpl{
		MenuRepository: repo,
		DB:             DB,
		Validate:       validate,
	}
}

// CheckMenu implements MenuService
func (service *MenuServiceImpl) CheckMenu(id int) bool {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	if tx.Error != nil {
		return true
	} else {
		result := service.MenuRepository.CheckMenu(tx, id)
		return result
	}
}

// InsertMenu implements MenuService
func (service *MenuServiceImpl) InsertMenu(request *dto.MenuCreateDTO) *dto.MenuResponseDTO {
	errorValidation := service.Validate.Struct(request)
	if errorValidation != nil {
		msgError := validation.TextValidation(errorValidation.Error())
		return &dto.MenuResponseDTO{
			Error:   true,
			Message: msgError,
		}
	}
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	if tx.Error != nil {
		msgError := validation.TextValidation(tx.Error.Error())
		return &dto.MenuResponseDTO{
			Error:   true,
			Message: msgError,
		}
	} else {
		menu := entity.Menu{
			Title:  request.Title,
			Image:  request.Image,
			Status: request.Status,
			Route:  request.Route,
		}

		response := service.MenuRepository.InsertMenu(tx, &menu)

		return dto.ToMenuResponseDTO(response)
	}
}

// UpdateMenu implements MenuService
func (service *MenuServiceImpl) UpdateMenu(request *dto.MenuUpdateDTO) *dto.MenuResponseDTO {
	errorValidation := service.Validate.Struct(request)
	if errorValidation != nil {
		msgError := validation.TextValidation(errorValidation.Error())
		return &dto.MenuResponseDTO{
			Error:   true,
			Message: msgError,
		}
	}
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	if tx.Error != nil {
		msgError := validation.TextValidation(tx.Error.Error())
		return &dto.MenuResponseDTO{
			Error:   true,
			Message: msgError,
		}
	} else {
		menu := entity.Menu{
			Title:  request.Title,
			Image:  request.Image,
			Status: request.Status,
			Route:  request.Route,
		}

		response := service.MenuRepository.UpdateMenu(tx, &menu)

		return dto.ToMenuResponseDTO(response)
	}
}

// GetMenu implements MenuService
func (*MenuServiceImpl) GetMenu() []*dto.MenuResponseDTO {
	panic("unimplemented")
}
