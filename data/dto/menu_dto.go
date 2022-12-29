package dto

import "app-hris-server/data/entity"

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
	Title  string `validate:"required" json:"title"`
	Image  string `validate:"required" json:"image"`
	Status int    `validate:"required" json:"status"`
	Route  string `validate:"required" json:"route"`
}

type MenuUpdateDTO struct {
	Id     int    `validate:"required" json:"id"`
	Title  string `validate:"required" json:"title"`
	Image  string `validate:"required" json:"image"`
	Status int    `validate:"required" json:"status"`
	Route  string `validate:"required" json:"route"`
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
