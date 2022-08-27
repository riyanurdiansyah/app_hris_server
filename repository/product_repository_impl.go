package repository

import (
	"app-ecommerce-server/data/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

// InsertProductDraft implements ProductRepository
func (*ProductRepositoryImpl) InsertProductDraft(c *gin.Context, db *gorm.DB, product *entity.Product) *entity.Product {
	result := db.Table("products_draft").Select("*").Create(&product)
	if result.Error != nil {
		product.ID = -99
		return product
	}

	return product
}
