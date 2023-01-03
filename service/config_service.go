package service

import (
	"app-hris-server/data/dto"
)

type ConfigService interface {
	InsertConfig(request *dto.ConfigCreateDTO) *dto.ConfigResponseDTO
	UpdateConfig(request *dto.ConfigUpdateDTO) *dto.ConfigResponseDTO
	GetConfig() []*dto.ConfigResponseDTO
	GetConfigByName(name string) *dto.ConfigResponseDTO
	CheckConfig(name string) bool
}
