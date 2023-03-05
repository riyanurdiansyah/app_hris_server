package service

import (
	"app-hris-server/data/dto"
)

type AttendanceService interface {
	Attendance(request *dto.AttendanceCreateDTO) *dto.AttendanceResponseDTO
	CheckAttendance(request *dto.AttendanceCreateDTO) bool
}
