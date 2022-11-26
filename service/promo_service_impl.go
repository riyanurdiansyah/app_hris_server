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

type PromoServiceImpl struct {
	PromoRepository repository.PromoRepository

	DB       *gorm.DB
	Validate *validator.Validate
}

func NewPromoService(promoRepository repository.PromoRepository, DB *gorm.DB, validate *validator.Validate) PromoService {
	return &PromoServiceImpl{
		PromoRepository: promoRepository,
		DB:              DB,
		Validate:        validate,
	}
}

// GetAllPromo implements PromoService
func (service *PromoServiceImpl) GetAllPromo() []*dto.PromoResponseDTO {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	if tx.Error != nil {
		return []*dto.PromoResponseDTO{}
	} else {
		listPromo := service.PromoRepository.GetAllPromo(tx)
		return dto.ToListPromoResponseDTO(listPromo)
	}
}

// InsertPromo implements PromoService
func (service *PromoServiceImpl) InsertPromo(request *dto.PromoCreateDTO) *dto.PromoResponseDTO {
	errorValidation := service.Validate.Struct(request)
	if errorValidation != nil {
		msgError := validation.TextValidation(errorValidation.Error())
		return &dto.PromoResponseDTO{
			Error:   true,
			Message: msgError,
		}
	}
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	if tx.Error != nil {
		msgError := validation.TextValidation(tx.Error.Error())
		return &dto.PromoResponseDTO{
			Error:   true,
			Message: msgError,
		}
	} else {
		promo := entity.Promo{
			Name:        request.Name,
			Image:       request.Path,
			Description: request.Description,
			KodePromo:   request.KodePromo,
			Expired:     request.Expired,
			Status:      1,
			CreatedAt:   time.Now().Local().String(),
			UpdatedAt:   time.Now().Local().String(),
		}

		promoResponse := service.PromoRepository.InsertPromo(tx, &promo)

		return dto.ToPromoResponseDTO(promoResponse)
	}
}

// FindPromoById implements PromoService
func (service *PromoServiceImpl) FindPromoById(promoId int) *dto.PromoResponseDTO {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	if tx.Error != nil {
		msgError := validation.TextValidation(tx.Error.Error())
		return &dto.PromoResponseDTO{
			Error:   true,
			Message: msgError,
		}
	} else {
		promoResponse := service.PromoRepository.FindPromoById(tx, promoId)
		return dto.ToPromoResponseDTO(promoResponse)
	}
}

// UpdatePromo implements PromoService
func (service *PromoServiceImpl) UpdatePromo(request *dto.PromoResponseDTO) *dto.PromoResponseDTO {
	errorValidation := service.Validate.Struct(request)
	if errorValidation != nil {
		msgError := validation.TextValidation(errorValidation.Error())
		return &dto.PromoResponseDTO{
			Error:   true,
			Message: msgError,
		}
	}
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	if tx.Error != nil {
		return &dto.PromoResponseDTO{
			Error:   true,
			Message: "terjadi kesalahan... silahkan coba beberapa saat lagi",
		}
	} else {
		promo := entity.Promo{
			ID:          request.ID,
			Name:        request.Name,
			Image:       request.Image,
			Description: request.Description,
			KodePromo:   request.KodePromo,
			Expired:     request.Expired,
			Status:      request.Status,
			CreatedAt:   request.Created,
			UpdatedAt:   time.Now().Local().String(),
		}
		promoResponse := service.PromoRepository.UpdatePromo(tx, &promo)

		return dto.ToPromoResponseDTO(promoResponse)
	}
}

// DeletePromo implements PromoService
func (service *PromoServiceImpl) DeletePromo(request *dto.PromoResponseDTO) *dto.PromoResponseDTO {
	errorValidation := service.Validate.Struct(request)
	if errorValidation != nil {
		msgError := validation.TextValidation(errorValidation.Error())
		return &dto.PromoResponseDTO{
			Error:   true,
			Message: msgError,
		}
	}
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	if tx.Error != nil {
		return &dto.PromoResponseDTO{
			Error:   true,
			Message: "terjadi kesalahan... silahkan coba beberapa saat lagi",
		}
	} else {
		promo := entity.Promo{
			ID:          request.ID,
			Name:        request.Name,
			Image:       request.Image,
			Description: request.Description,
			KodePromo:   request.KodePromo,
			Expired:     request.Expired,
			Status:      request.Status,
			CreatedAt:   request.Created,
			UpdatedAt:   time.Now().Local().String(),
		}
		promoResponse := service.PromoRepository.DeletePromo(tx, &promo)

		return dto.ToPromoResponseDTO(promoResponse)
	}
}
