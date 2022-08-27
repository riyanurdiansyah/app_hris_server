package repository

import (
	"app-ecommerce-server/data/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PromoRepository interface {
	InsertPromo(c *gin.Context, db *gorm.DB, promo *entity.Promo) *entity.Promo
	GetAllPromo(c *gin.Context, db *gorm.DB) []*entity.Promo
}
