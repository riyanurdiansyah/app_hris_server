package repository

import (
	"app-hris-server/data/dto"
	"app-hris-server/data/entity"

	"gorm.io/gorm"
)

type AttendanceRepositoryImpl struct{}

func NewAttendanceRepository() AttendanceRepository {
	return &AttendanceRepositoryImpl{}
}

// CheckIfDoneClockin implements AttendanceRepository
func (*AttendanceRepositoryImpl) CheckIfDoneClockin(db *gorm.DB, absent *dto.ClockinCreateDTO) bool {
	panic("unimplemented")
}

// CheckIfDoneClockout implements AttendanceRepository
func (*AttendanceRepositoryImpl) CheckIfDoneClockout(db *gorm.DB, absent *dto.ClockoutCreateDTO) bool {
	panic("unimplemented")
}

// Clockin implements AttendanceRepository
func (*AttendanceRepositoryImpl) Clockin(db *gorm.DB, absent *entity.ClockIn) *entity.ClockIn {
	result := db.Table("attendance").Select("*").Create(&absent)
	if result.Error != nil {
		///handle panic
		panic(result.Error)
	}
	return absent
}

// Clockout implements AttendanceRepository
func (*AttendanceRepositoryImpl) Clockout(db *gorm.DB, absent *entity.ClockOut) *entity.ClockOut {
	panic("unimplemented")
}
