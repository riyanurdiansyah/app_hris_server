package repository

import (
	"app-ecommerce-server/data/entity"
	"app-ecommerce-server/helper"

	"gorm.io/gorm"
)

type PromoRepositoryImpl struct {
}

func NewPromoRepository() PromoRepository {
	return &PromoRepositoryImpl{}
}

// InsertPromo implements PromoRepository
func (repo *PromoRepositoryImpl) InsertPromo(db *gorm.DB, promo *entity.Promo) *entity.Promo {
	result := db.Table("promos_slider").Select("*").Create(&promo)
	if result.Error != nil {
		promo.ID = -99
		return promo
	}

	return promo
}

// GetAllPromo implements PromoRepository
func (*PromoRepositoryImpl) GetAllPromo(db *gorm.DB) []*entity.Promo {
	var listPromo = []*entity.Promo{}
	result :=
		db.Table("promos_slider").Select("*").Scan(&listPromo)
	helper.PanicIfError(result.Error)
	return listPromo
}
