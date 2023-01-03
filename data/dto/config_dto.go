package dto

import (
	"app-hris-server/data/entity"
)

type ConfigResponseDTO struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Value     string `json:"value"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Error     bool   `json:"-"`
	Message   string `json:"-"`
}

type ConfigCreateDTO struct {
	Name    string `validate:"required" json:"name"`
	Value   string `validate:"required" json:"value"`
	Created string `json:"created_at"`
	Updated string `json:"updated_at"`
}

type ConfigUpdateDTO struct {
	ID    int    `validate:"required" json:"id"`
	Name  string `validate:"required" json:"name"`
	Value string `validate:"required" json:"value"`
}

func ToConfigResponseDTO(ent *entity.Config) *ConfigResponseDTO {

	return &ConfigResponseDTO{
		ID:        ent.ID,
		Name:      ent.Name,
		Value:     ent.Value,
		CreatedAt: ent.CreatedAt,
		UpdatedAt: ent.UpdatedAt,
	}
}

func ToListConfigResponseDTO(promo []*entity.Config) []*ConfigResponseDTO {
	var listTemp = []*ConfigResponseDTO{}
	for _, data := range promo {
		listTemp = append(listTemp, &ConfigResponseDTO{
			ID:        data.ID,
			Name:      data.Name,
			Value:     data.Value,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,
		})
	}
	return listTemp
}
