package repository

import (
	"app-ecommerce-server/data/entity"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	InsertCategory(db *gorm.DB, category *entity.Category) *entity.Category
	FindAllCategory(db *gorm.DB) []*entity.Category
	FindByIdCategory(db *gorm.DB, categoryId int) *entity.Category
	DeleteCategory(db *gorm.DB, categoryId int) int
	UpdateCategory(db *gorm.DB, category *entity.Category) *entity.Category
}
