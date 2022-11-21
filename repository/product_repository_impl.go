package repository

import (
	"app-travel-server/data/entity"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

// InsertProductDraft implements ProductRepository
func (*ProductRepositoryImpl) InsertProductDraft(db *gorm.DB, product *entity.Product) *entity.Product {
	result := db.Table("products_draft").Select("*").Create(&product)
	if result.Error != nil {
		product.ID = -99
		return product
	}

	return product
}
