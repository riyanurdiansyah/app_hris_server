package dto

import (
	"app-travel-server/data/entity"
)

func ToPromoResponseDTO(promo *entity.Promo) *PromoResponseDTO {
	return &PromoResponseDTO{
		ID:          promo.ID,
		Name:        promo.Name,
		Image:       promo.Image,
		Description: promo.Description,
		KodePromo:   promo.KodePromo,
		Expired:     promo.Expired,
		Status:      promo.Status,
		Created:     promo.CreatedAt,
		Updated:     promo.UpdatedAt,
	}
}

func ToListPromoResponseDTO(promo []*entity.Promo) []*PromoResponseDTO {
	var listTemp = []*PromoResponseDTO{}
	for _, data := range promo {
		listTemp = append(listTemp, &PromoResponseDTO{
			ID:          data.ID,
			Name:        data.Name,
			Image:       data.Image,
			Description: data.Description,
			KodePromo:   data.KodePromo,
			Status:      data.Status,
			Expired:     data.Expired,
			Created:     data.CreatedAt,
			Updated:     data.UpdatedAt,
		})
	}
	return listTemp
}
