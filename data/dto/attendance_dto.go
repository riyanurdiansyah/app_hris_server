package dto

import (
	"app-hris-server/data/entity"
	"mime/multipart"
)

type AttendanceResponseDTO struct {
	// ID                int     `json:"id"`
	UserId            string  `json:"user_id"`
	TimeClockin       string  `json:"time_clockin"`
	TimeClockout      string  `json:"time_clockout"`
	ImageClockin      string  `json:"image_clockin"`
	ImageClockout     string  `json:"image_clockout"`
	LatitudeClockin   float64 `json:"latitude_clockin"`
	LatitudeClockout  float64 `json:"latitude_clockout"`
	LongitudeClockin  float64 `json:"longitude_clockin"`
	LongitudeClockout float64 `json:"longitude_clockout"`
	NoteClockin       string  `json:"note_clockin"`
	NoteClockout      string  `json:"note_clockout"`
	Error             bool    `json:"-"`
	Message           string  `json:"-"`
}

type AttendanceCreateDTO struct {
	UserId    string                `validate:"required" form:"uuid_user"`
	Time      string                `validate:"required" form:"time"`
	Image     *multipart.FileHeader `form:"image" validate:"required"`
	Path      string
	Latitude  float64 `validate:"required" form:"latitude"`
	Longitude float64 `validate:"required" form:"longitude"`
	Note      string  `form:"note"`
	Kode      string  `validate:"required" form:"kode"`
	Date      string  `validate:"required" form:"date"`
}

type ClockinCreateDTO struct {
	UserId           string                `validate:"required" form:"user_id"`
	TimeClockin      string                `validate:"required" form:"time_clockin"`
	Image            *multipart.FileHeader `form:"image" validate:"required"`
	Path             string
	LatitudeClockin  float64 `validate:"required" form:"latitude_clockin"`
	LongitudeClockin float64 `validate:"required" form:"longitude_clockin"`
	NoteClockin      string  `form:"note_clockin"`
}

type ClockoutCreateDTO struct {
	UserId            string                `validate:"required" form:"user_id"`
	TimeClockout      string                `validate:"required" form:"time_clockout"`
	Image             *multipart.FileHeader `form:"image" validate:"required"`
	Path              string                `validate:"required"`
	LatitudeClockout  float64               `validate:"required" form:"latitude_clockout"`
	LongitudeClockout float64               `validate:"required" form:"longitude_clockout"`
	NoteClockout      string                `form:"note_clockout"`
}

func ToAttendanceResponseDTO(ent *entity.Attendance) *AttendanceResponseDTO {

	return &AttendanceResponseDTO{
		// ID:                ent.ID,
		UserId:            ent.UserId,
		TimeClockin:       ent.TimeClockin,
		TimeClockout:      ent.TimeClockout,
		ImageClockin:      ent.ImageClockin,
		ImageClockout:     ent.ImageClockout,
		LatitudeClockin:   ent.LatitudeClockin,
		LongitudeClockin:  ent.LongitudeClockin,
		LatitudeClockout:  ent.LatitudeClockout,
		LongitudeClockout: ent.LongitudeClockout,
		NoteClockin:       ent.NoteClockin,
		NoteClockout:      ent.NoteClockout,
	}
}
