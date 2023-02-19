package service

import (
	"app-hris-server/data/dto"
)

type AttendanceService interface {
	Clockin(request *dto.ClockinCreateDTO) *dto.AttendanceResponseDTO
	Clockout(request *dto.ClockoutCreateDTO) *dto.AttendanceResponseDTO
	CheckIfDoneClockin(request *dto.ClockinCreateDTO) bool
	CheckIfDoneClockout(request *dto.ClockoutCreateDTO) bool
}
