package service

import (
	"app-ecommerce-server/data/dto"

	"github.com/gin-gonic/gin"
)

type PromoService interface {
	InsertPromo(ctx *gin.Context, request *dto.PromoCreateDTO) *dto.PromoResponseDTO
	GetAllPromo(ctx *gin.Context) []*dto.PromoResponseDTO
}
