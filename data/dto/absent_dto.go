package dto

import "app-hris-server/data/entity"

type AbsentResponseDTO struct {
	IdEmployee string  `json:"id_employee"`
	Tanggal    string  `json:"tanggal"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Catatan    string  `json:"catatan"`
	Tipe       int     `json:"tipe"`
	Photo      string  `json:"photo"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
	Error      bool    `json:"-"`
	Message    string  `json:"-"`
}
type AbsentCreateDTO struct {
	IdUser     string  `validate:"required" json:"id_user"`
	IdEmployee string  `validate:"required" json:"id_employee"`
	Tanggal    string  `validate:"required" json:"tanggal"`
	Latitude   float64 `validate:"required" json:"latitude"`
	Longitude  float64 `validate:"required" json:"longitude"`
	Catatan    string  `json:"catatan"`
	Tipe       int     `validate:"required" json:"tipe"`
	Photo      string  `validate:"required" json:"photo"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}

func ToAbsentResponseDTO(ent *entity.Absent) *AbsentResponseDTO {

	return &AbsentResponseDTO{
		IdEmployee: ent.IdEmployee,
		Tanggal:    ent.Tanggal,
		Latitude:   ent.Latitude,
		Longitude:  ent.Longitude,
		Tipe:       ent.Tipe,
		Catatan:    ent.Catatan,
		Photo:      ent.Photo,
		CreatedAt:  ent.Photo,
		UpdatedAt:  ent.UpdatedAt,
	}
}
