package dto

import "app-hris-server/data/entity"

type CompanyDTO struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	SecretKey string `json:"secret_key"`
	Created   string `json:"created_at"`
	Updated   string `json:"updated_at"`
}

type CompanyResponseDTO struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	SecretKey string `json:"secret_key"`
	Created   string `json:"created_at"`
	Updated   string `json:"updated_at"`
	Error     bool   `json:"-"`
	Message   string `json:"-"`
}

type CompanyCreateDTO struct {
	Name      string `validate:"required" json:"name"`
	SecretKey string `json:"secret_key"`
}

func ToCompanyResponseDTO(ent *entity.Company) *CompanyResponseDTO {
	return &CompanyResponseDTO{
		ID:        ent.ID,
		Name:      ent.Name,
		SecretKey: ent.SecretKey,
		Created:   ent.CreatedAt,
		Updated:   ent.UpdatedAt,
	}
}

func ToListCompanyResponseDTO(ent []*entity.Company) []*CompanyResponseDTO {
	var listTemp = []*CompanyResponseDTO{}
	for _, data := range ent {
		listTemp = append(listTemp, &CompanyResponseDTO{
			ID:        data.ID,
			Name:      data.Name,
			SecretKey: data.SecretKey,
			Created:   data.CreatedAt,
			Updated:   data.UpdatedAt,
		})
	}
	return listTemp
}
