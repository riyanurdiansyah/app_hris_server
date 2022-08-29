package service

import (
	"app-ecommerce-server/data/dto"
)

type PromoService interface {
	InsertPromo(request *dto.PromoCreateDTO) *dto.PromoResponseDTO
	GetAllPromo() []*dto.PromoResponseDTO
	FindPromoById(request *dto.PromoUpdateDTO) *dto.PromoResponseDTO
	UpdatePromo(request *dto.PromoResponseDTO) *dto.PromoResponseDTO
}
