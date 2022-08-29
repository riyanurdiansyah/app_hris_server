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

// FindPromoById implements PromoRepository
func (repo *PromoRepositoryImpl) FindPromoById(db *gorm.DB, promoId int) *entity.Promo {
	var promo = entity.Promo{}
	result := db.Table("promos_slider").Select("*").Where("id = ?", promoId).Scan(&promo)
	if result.Error != nil {
		promo.ID = -99
		return &promo
	}
	return &promo
}

// UpdatePromo implements PromoRepository
func (repo *PromoRepositoryImpl) UpdatePromo(db *gorm.DB, promo *entity.Promo) *entity.Promo {
	result :=
		db.Table("promos_slider").Where("id = ?", promo.ID).Updates(&promo)
	if result.Error != nil {
		promo.ID = -99
		return promo
	}
	return promo
}

func (repo *PromoRepositoryImpl) DeletePromo(db *gorm.DB, promo *entity.Promo) *entity.Promo {
	result :=
		db.Table("promos_slider").Where("id = ?", promo.ID).Delete(&promo)
	if result.Error != nil {
		promo.ID = -99
		return promo
	}
	return promo
}
