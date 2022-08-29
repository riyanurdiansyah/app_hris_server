package dto

import "mime/multipart"

type PromoDTO struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	KodePromo   string `json:"kode_promo"`
	Expired     int    `json:"expired"`
	Status      int    `json:"status"`
	Created     string `json:"created_at"`
	Updated     string `json:"updated_at"`
}

type PromoCreateDTO struct {
	Name        string                `form:"name" validate:"required"`
	Image       *multipart.FileHeader `form:"image" validate:"required"`
	Path        string                `validate:"required"`
	Description string                `form:"description" validate:"required"`
	KodePromo   string                `form:"kode_promo" validate:"required"`
	Expired     int                   `form:"expired" validate:"required"`
}

type PromoUpdateDTO struct {
	ID     int `json:"id" validate:"required"`
	Status int `json:"status" validate:"required"`
}

type PromoResponseDTO struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	KodePromo   string `json:"kode_promo"`
	Expired     int    `json:"expired"`
	Status      int    `json:"status"`
	Created     string `json:"created_at"`
	Updated     string `json:"updated_at"`
	Error       bool   `json:"-"`
	Message     string `json:"-"`
}
