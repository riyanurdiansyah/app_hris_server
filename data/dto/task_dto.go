package dto

import (
	"app-hris-server/data/entity"
)

type TaskResponseDTO struct {
	// ID       int               `json:"id"`
	// Uuid     string            `json:"uuid"`
	Title    string            `json:"title"`
	Status   int               `json:"status"`
	Progress float64           `json:"progress"`
	TaskBy   TaskByResponseDTO `json:"task_by"`
	Error    bool              `json:"-"`
	Message  string            `json:"-"`
}

type TaskByResponseDTO struct {
	NamaDepan    string `json:"nama_depan"`
	NamaBelakang string `json:"nama_belakang"`
	Image        string `json:"image"`
	NoHp         string `json:"no_hp"`
}

type TaskCreateDTO struct {
	IdUser   string  `validate:"required" json:"user_id"`
	Title    string  `validate:"required" json:"title"`
	TaskBy   int     `validate:"required" json:"task_by"`
	Status   int     `validate:"required" json:"status"`
	Progress float64 `validate:"required" json:"progress"`
}

type TaskUpdateDTO struct {
	ID       string  `validate:"required" json:"id"`
	IdUser   string  `validate:"required" json:"user_id"`
	Title    string  `validate:"required" json:"title"`
	TaskBy   int     `validate:"required" json:"task_by"`
	Status   int     `validate:"required" json:"status"`
	Progress float64 `validate:"required" json:"progress"`
	Error    bool    `json:"-"`
	Message  string  `json:"-"`
}

func ToTaskByResponseDTO(ent *entity.TaskBy) *TaskByResponseDTO {
	return &TaskByResponseDTO{
		NamaDepan:    ent.NamaDepan,
		NamaBelakang: ent.NamaBelakang,
		Image:        ent.Image,
		NoHp:         ent.NoHp,
	}
}

func ToListTaskResponseDTO(promo []*entity.Task) []*TaskResponseDTO {
	var listTemp = []*TaskResponseDTO{}

	for _, data := range promo {
		listTemp = append(listTemp, &TaskResponseDTO{
			// ID:       data.ID,
			Title: data.Title,
			// IdUser:   data.IdUser,
			Progress: data.Progress,
			Status:   data.Status,
			TaskBy: TaskByResponseDTO{
				NamaDepan:    data.NamaDepan,
				NamaBelakang: data.NamaBelakang,
				Image:        data.Image,
				NoHp:         data.NoHp,
			},
		})
	}
	return listTemp
}
