package dto

import (
	"app-hris-server/data/entity"
	"mime/multipart"
)

type MenuResponseDTO struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Image   string `json:"image"`
	Status  int    `json:"status"`
	Route   string `json:"route"`
	Error   bool   `json:"-"`
	Message string `json:"-"`
}

type MenuCreateDTO struct {
	Title  string                `validate:"required" form:"title"`
	Image  *multipart.FileHeader `form:"image" validate:"required"`
	Path   string                `validate:"required"`
	Status int                   `validate:"required" form:"status"`
	Route  string                `validate:"required" form:"route"`
}

type MenuUpdateDTO struct {
	Id     int                   `validate:"required" form:"id"`
	Title  string                `validate:"required" form:"title"`
	Image  *multipart.FileHeader `form:"image" validate:"required"`
	Path   string                `validate:"required"`
	Status int                   `validate:"required" form:"status"`
	Route  string                `validate:"required" form:"route"`
}

func ToMenuResponseDTO(ent *entity.Menu) *MenuResponseDTO {

	return &MenuResponseDTO{
		ID:     ent.ID,
		Title:  ent.Title,
		Image:  ent.Image,
		Status: ent.Status,
		Route:  ent.Route,
	}
}

func ToListMenuResponseDTO(promo []*entity.Menu) []*MenuResponseDTO {
	var listTemp = []*MenuResponseDTO{}
	for _, data := range promo {
		listTemp = append(listTemp, &MenuResponseDTO{
			ID:     data.ID,
			Title:  data.Title,
			Image:  data.Image,
			Status: data.Status,
			Route:  data.Route,
		})
	}
	return listTemp
}
