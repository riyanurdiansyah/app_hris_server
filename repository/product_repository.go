package repository

import (
	"app-ecommerce-server/data/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductRepository interface {
	InsertProductDraft(c *gin.Context, db *gorm.DB, product *entity.Product) *entity.Product
}
