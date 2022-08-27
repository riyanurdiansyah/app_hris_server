package service

import (
	"app-ecommerce-server/data/dto"
	"app-ecommerce-server/data/entity"
	"app-ecommerce-server/helper"
	"app-ecommerce-server/repository"
	"app-ecommerce-server/validation"
	"time"

	"github.com/gin-gonic/gin"
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
func (service *PromoServiceImpl) GetAllPromo(ctx *gin.Context) []*dto.PromoResponseDTO {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	if tx.Error != nil {
		return []*dto.PromoResponseDTO{}
	} else {
		listPromo := service.PromoRepository.GetAllPromo(ctx, tx)
		return dto.ToListPromoResponseDTO(listPromo)
	}
}

// InsertPromo implements PromoService
func (service *PromoServiceImpl) InsertPromo(ctx *gin.Context, request *dto.PromoCreateDTO) *dto.PromoResponseDTO {
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

		promoResponse := service.PromoRepository.InsertPromo(ctx, tx, &promo)

		return dto.ToPromoResponseDTO(promoResponse)
	}
}
