package repository

import (
	"app-ecommerce-server/data/entity"

	"gorm.io/gorm"
)

type PromoRepository interface {
	InsertPromo(db *gorm.DB, promo *entity.Promo) *entity.Promo
	GetAllPromo(db *gorm.DB) []*entity.Promo
}
