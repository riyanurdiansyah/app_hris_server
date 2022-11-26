package service

import (
	"app-hris-server/data/dto"
)

type PromoService interface {
	InsertPromo(request *dto.PromoCreateDTO) *dto.PromoResponseDTO
	GetAllPromo() []*dto.PromoResponseDTO
	FindPromoById(promoId int) *dto.PromoResponseDTO
	UpdatePromo(request *dto.PromoResponseDTO) *dto.PromoResponseDTO
	DeletePromo(request *dto.PromoResponseDTO) *dto.PromoResponseDTO
}
