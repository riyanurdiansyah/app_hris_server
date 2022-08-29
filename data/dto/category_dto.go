package dto

import "mime/multipart"

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
	Created string `json:"created_at"`
	Updated string `json:"updated_at"`
	Error   bool   `json:"-"`
	Message string `json:"-"`
}
