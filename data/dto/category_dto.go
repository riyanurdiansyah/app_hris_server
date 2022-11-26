package dto

import (
	"app-hris-server/data/entity"
	"mime/multipart"
)

type CategoryDTO struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Created string `json:"created_at"`
	Updated string `json:"updated_at"`
}

type CategoryCreateDTO struct {
	Name  string                `form:"name" validate:"required"`
	Image *multipart.FileHeader `form:"image" validate:"required"`
	Path  string
}

type CategoryUpdateDTO struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CategoryResponseDTO struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Image   string `json:"image"`
	Status  int    `json:"status"`
	Created string `json:"created_at"`
	Updated string `json:"updated_at"`
	Error   bool   `json:"-"`
	Message string `json:"-"`
}

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
