package service

import (
	"app-ecommerce-server/data/dto"

	"github.com/gin-gonic/gin"
)

type CategoryService interface {
	InsertCategory(request *dto.CategoryCreateDTO) *dto.CategoryResponseDTO
	FindAllCategory(ctx *gin.Context) []*dto.CategoryResponseDTO
	FindByIdCategory(categoryId int) *dto.CategoryResponseDTO
	DeleteCategory(request *dto.CategoryResponseDTO) *dto.CategoryResponseDTO
	UpdateCategory(request *dto.CategoryUpdateDTO) *dto.CategoryResponseDTO
}
