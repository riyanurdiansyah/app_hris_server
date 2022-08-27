package service

import (
	"app-ecommerce-server/data/dto"
)

type PromoService interface {
	InsertPromo(request *dto.PromoCreateDTO) *dto.PromoResponseDTO
	GetAllPromo() []*dto.PromoResponseDTO
}
