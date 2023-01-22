package dto

import (
	"app-hris-server/data/entity"
	"mime/multipart"
)

type MenuResponseDTO struct {
	ID          int    `json:"id"`
	Category    string `json:"category"`
	Position    int    `json:"position"`
	Title       string `json:"title"`
	Image       string `json:"image"`
	Status      int    `json:"status"`
	Language    string `json:"language"`
	Route       string `json:"route"`
	Description string `json:"description"`
	Error       bool   `json:"-"`
	Message     string `json:"-"`
}

type MenuCreateDTO struct {
	Category    string                `validate:"required" form:"category"`
	Position    int                   `validate:"required" form:"position"`
	Title       string                `validate:"required" form:"title"`
	Image       *multipart.FileHeader `form:"image" validate:"required"`
	Path        string                `validate:"required"`
	Language    string                `validate:"required" form:"language"`
	Status      int                   `validate:"required" form:"status"`
	Route       string                `validate:"required" form:"route"`
	Description string                `validate:"required" form:"description"`
}

type MenuUpdateDTO struct {
	Id          int                   `validate:"required" form:"id"`
	Category    string                `validate:"required" form:"category"`
	Position    int                   `validate:"required" form:"position"`
	Title       string                `validate:"required" form:"title"`
	Image       *multipart.FileHeader `form:"image" validate:"required"`
	Path        string                `validate:"required"`
	Language    string                `validate:"required" form:"language"`
	Status      int                   `validate:"required" form:"status"`
	Route       string                `validate:"required" form:"route"`
	Description string                `validate:"required" form:"description"`
}

func ToMenuResponseDTO(ent *entity.Menu) *MenuResponseDTO {

	return &MenuResponseDTO{
		ID:          ent.ID,
		Category:    ent.Category,
		Position:    ent.Position,
		Title:       ent.Title,
		Image:       ent.Image,
		Status:      ent.Status,
		Language:    ent.Language,
		Route:       ent.Route,
		Description: ent.Description,
	}
}

func ToListMenuResponseDTO(promo []*entity.Menu) []*MenuResponseDTO {
	var listTemp = []*MenuResponseDTO{}
	for _, data := range promo {
		listTemp = append(listTemp, &MenuResponseDTO{
			ID:          data.ID,
			Category:    data.Category,
			Position:    data.Position,
			Title:       data.Title,
			Image:       data.Image,
			Status:      data.Status,
			Language:    data.Language,
			Route:       data.Route,
			Description: data.Description,
		})
	}
	return listTemp
}
