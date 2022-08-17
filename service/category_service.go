package service

import (
	"belajar/dto"

	"github.com/gin-gonic/gin"
)

type CategoryService interface {
	InsertCategory(ctx *gin.Context, request dto.CategoryCreateDTO) dto.CategoryResponseDTO
	FindAllCategory(ctx *gin.Context) []dto.CategoryResponseDTO
	FindByIdCategory(ctx *gin.Context, categoryId int) dto.CategoryResponseDTO
	DeleteCategory(ctx *gin.Context, categoryId int) int
	UpdateCategory(ctx *gin.Context, request dto.CategoryUpdateDTO) dto.CategoryResponseDTO
}
