package repository

import (
	"app-ecommerce-server/data/entity"

	"gorm.io/gorm"
)

type ProductRepository interface {
	InsertProductDraft(db *gorm.DB, product *entity.Product) *entity.Product
}
