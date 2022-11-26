package repository

import (
	"app-hris-server/data/entity"

	"gorm.io/gorm"
)

type PromoRepository interface {
	InsertPromo(db *gorm.DB, promo *entity.Promo) *entity.Promo
	GetAllPromo(db *gorm.DB) []*entity.Promo
	FindPromoById(db *gorm.DB, promoId int) *entity.Promo
	UpdatePromo(db *gorm.DB, promo *entity.Promo) *entity.Promo
	DeletePromo(db *gorm.DB, promo *entity.Promo) *entity.Promo
}
