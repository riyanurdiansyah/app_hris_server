package service

import (
	"app-hris-server/data/dto"
)

type AbsentService interface {
	InsertAbsent(request *dto.AbsentCreateDTO) *dto.AbsentResponseDTO
	CheckIfDoneAbsent(request *dto.AbsentCreateDTO) bool
}
