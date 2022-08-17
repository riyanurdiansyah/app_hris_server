package repository

import (
	"app-ecommerce-server/data/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	InsertCategory(ctx *gin.Context, db *gorm.DB, category *entity.Category) *entity.Category
	FindAllCategory(ctx *gin.Context, db *gorm.DB) []*entity.Category
	FindByIdCategory(ctx *gin.Context, db *gorm.DB, categoryId int) *entity.Category
	DeleteCategory(ctx *gin.Context, db *gorm.DB, categoryId int) int
	UpdateCategory(ctx *gin.Context, db *gorm.DB, category *entity.Category) *entity.Category
}
