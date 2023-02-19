package repository

import (
	"app-hris-server/data/dto"
	"app-hris-server/data/entity"

	"gorm.io/gorm"
)

type AttendanceRepository interface {
	Clockin(db *gorm.DB, absent *entity.ClockIn) *entity.ClockIn
	Clockout(db *gorm.DB, absent *entity.ClockOut) *entity.ClockOut
	CheckIfDoneClockin(db *gorm.DB, absent *dto.ClockinCreateDTO) bool
	CheckIfDoneClockout(db *gorm.DB, absent *dto.ClockoutCreateDTO) bool
}
