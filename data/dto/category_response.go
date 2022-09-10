package dto

import (
	"app-ecommerce-server/data/entity"
)

func ToCategoryResponseDTO(category *entity.Category) *CategoryResponseDTO {
	return &CategoryResponseDTO{
		Id:      category.ID,
		Name:    category.Name,
		Image:   category.Image,
		Status:  category.Status,
		Created: category.CreatedAt,
		Updated: category.UpdatedAt,
	}
}

func ToListCategoryResponseDTO(category []*entity.Category) []*CategoryResponseDTO {
	var listTemp = []*CategoryResponseDTO{}
	for _, data := range category {
		listTemp = append(listTemp, &CategoryResponseDTO{
			Id:      data.ID,
			Name:    data.Name,
			Image:   data.Image,
			Status:  data.Status,
			Created: data.CreatedAt,
			Updated: data.UpdatedAt,
		})
	}
	return listTemp
}
