package service

import (
	"app-hris-server/data/dto"
)

type CompanyService interface {
	InsertCompany(request *dto.CompanyCreateDTO) *dto.CompanyResponseDTO
}
