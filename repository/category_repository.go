package repository

import (
	"app-travel-server/data/entity"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	InsertCategory(db *gorm.DB, category *entity.Category) *entity.Category
	FindAllCategory(db *gorm.DB, page int) ([]*entity.Category, int64)
	FindByIdCategory(db *gorm.DB, categoryId int) *entity.Category
	DeleteCategory(db *gorm.DB, category *entity.Category) *entity.Category
	UpdateCategory(db *gorm.DB, category *entity.Category) *entity.Category
}
