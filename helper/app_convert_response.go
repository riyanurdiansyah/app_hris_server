package helper

import (
	"belajar/dto"
	"belajar/entity"
)

func ToCategoryResponseDTO(category entity.Category) dto.CategoryResponseDTO {
	return dto.CategoryResponseDTO{
		Id:      category.ID,
		Name:    category.Name,
		Created: category.CreatedAt,
		Updated: category.UpdatedAt,
	}
}

func ToListCategoryResponseDTO(category []entity.Category) []dto.CategoryResponseDTO {
	var listTemp = []dto.CategoryResponseDTO{}
	for _, data := range category {
		listTemp = append(listTemp, dto.CategoryResponseDTO{
			Id:      data.ID,
			Name:    data.Name,
			Created: data.CreatedAt,
			Updated: data.UpdatedAt,
		})
	}
	return listTemp
}
