package service

import (
	"app-ecommerce-server/data/dto"
	"app-ecommerce-server/data/entity"
	"app-ecommerce-server/helper"
	"app-ecommerce-server/repository"
	"app-ecommerce-server/validation"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository

	DB       *gorm.DB
	Validate *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *gorm.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) InsertCategory(request *dto.CategoryCreateDTO) *dto.CategoryResponseDTO {
	errorValidation := service.Validate.Struct(request)
	if errorValidation != nil {
		msgError := validation.TextValidation(errorValidation.Error())
		return &dto.CategoryResponseDTO{
			Error:   true,
			Message: msgError,
		}
	}
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	if tx.Error != nil {
		msgError := validation.TextValidation(tx.Error.Error())
		return &dto.CategoryResponseDTO{
			Error:   true,
			Message: msgError,
		}
	} else {

		category := entity.Category{
			Name:      request.Name,
			Image:     request.Path,
			Status:    1,
			CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
			UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
		}

		categoryResponse := service.CategoryRepository.InsertCategory(tx, &category)

		return dto.ToCategoryResponseDTO(categoryResponse)
	}
}

func (service *CategoryServiceImpl) FindAllCategory(ctx *gin.Context) ([]*dto.CategoryResponseDTO, int64) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	if tx.Error != nil {
		return []*dto.CategoryResponseDTO{}, 0
	} else {
		listCategory, total := service.CategoryRepository.FindAllCategory(tx, page)
		return dto.ToListCategoryResponseDTO(listCategory), total
	}
}

func (service *CategoryServiceImpl) FindByIdCategory(categoryId int) *dto.CategoryResponseDTO {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	if tx.Error != nil {
		msgError := validation.TextValidation(tx.Error.Error())
		return &dto.CategoryResponseDTO{
			Error:   true,
			Message: msgError,
		}
	} else {
		category := service.CategoryRepository.FindByIdCategory(tx, categoryId)
		return dto.ToCategoryResponseDTO(category)
	}
}

func (service *CategoryServiceImpl) DeleteCategory(request *dto.CategoryResponseDTO) *dto.CategoryResponseDTO {
	errorValidation := service.Validate.Struct(request)
	if errorValidation != nil {
		msgError := validation.TextValidation(errorValidation.Error())
		return &dto.CategoryResponseDTO{
			Error:   true,
			Message: msgError,
		}
	}
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	if tx.Error != nil {
		return &dto.CategoryResponseDTO{
			Error:   true,
			Message: "terjadi kesalahan... silahkan coba beberapa saat lagi",
		}
	} else {
		promo := entity.Category{
			ID:        request.Id,
			Name:      request.Name,
			Image:     request.Image,
			CreatedAt: request.Created,
			UpdatedAt: time.Now().Local().String(),
		}
		promoResponse := service.CategoryRepository.DeleteCategory(tx, &promo)

		return dto.ToCategoryResponseDTO(promoResponse)
	}
}

func (service *CategoryServiceImpl) UpdateCategory(request *dto.CategoryUpdateDTO) *dto.CategoryResponseDTO {
	errorValidation := service.Validate.Struct(request)
	if errorValidation != nil {
		msgError := validation.TextValidation(errorValidation.Error())
		return &dto.CategoryResponseDTO{
			Error:   true,
			Message: msgError,
		}
	}
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	if tx.Error != nil {
		return &dto.CategoryResponseDTO{
			Error:   true,
			Message: "terjadi kesalahan... silahkan coba beberapa saat lagi",
		}
	} else {

		category := service.CategoryRepository.FindByIdCategory(tx, request.Id)
		if category.Name == "" {
			return &dto.CategoryResponseDTO{
				Error:   true,
				Message: "id is not found",
			}
		} else {
			categorys := entity.Category{
				ID:        request.Id,
				Name:      request.Name,
				CreatedAt: category.CreatedAt,
				UpdatedAt: time.Now().Local().String(),
			}

			categoryResponse := service.CategoryRepository.UpdateCategory(tx, &categorys)

			return dto.ToCategoryResponseDTO(categoryResponse)
		}
	}
}
