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

func (service *CategoryServiceImpl) InsertCategory(ctx *gin.Context, request *dto.CategoryCreateDTO) *dto.CategoryResponseDTO {
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
			CreatedAt: time.Now().Local().String(),
			UpdatedAt: time.Now().Local().String(),
		}

		categoryResponse := service.CategoryRepository.InsertCategory(ctx, tx, &category)

		return dto.ToCategoryResponseDTO(categoryResponse)
	}
}

func (service *CategoryServiceImpl) FindAllCategory(ctx *gin.Context) []*dto.CategoryResponseDTO {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	if tx.Error != nil {
		return []*dto.CategoryResponseDTO{}
	} else {
		listCategory := service.CategoryRepository.FindAllCategory(ctx, tx)
		return dto.ToListCategoryResponseDTO(listCategory)
	}
}

func (service *CategoryServiceImpl) FindByIdCategory(ctx *gin.Context, categoryId int) *dto.CategoryResponseDTO {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	if tx.Error != nil {
		return &dto.CategoryResponseDTO{
			Error:   true,
			Message: "terjadi kesalahan.. silahkan coba lagi",
		}
	} else {
		category := service.CategoryRepository.FindByIdCategory(ctx, tx, categoryId)
		return dto.ToCategoryResponseDTO(category)
	}
}

func (service *CategoryServiceImpl) DeleteCategory(ctx *gin.Context, categoryId int) int {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	if tx.Error != nil {
		return -1
	} else {
		count := service.CategoryRepository.DeleteCategory(ctx, tx, categoryId)
		return count
	}
}

func (service *CategoryServiceImpl) UpdateCategory(ctx *gin.Context, request *dto.CategoryUpdateDTO) *dto.CategoryResponseDTO {
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

		category := service.CategoryRepository.FindByIdCategory(ctx, tx, request.Id)
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

			categoryResponse := service.CategoryRepository.UpdateCategory(ctx, tx, &categorys)

			return dto.ToCategoryResponseDTO(categoryResponse)
		}
	}
}
