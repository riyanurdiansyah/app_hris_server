package repository

import (
	"app-hris-server/data/entity"

	"gorm.io/gorm"
)

type AttendanceRepositoryImpl struct{}

func NewAttendanceRepository() AttendanceRepository {
	return &AttendanceRepositoryImpl{}
}

// Attendance implements AttendanceRepository
func (*AttendanceRepositoryImpl) Attendance(db *gorm.DB, ent *entity.Attendance) *entity.Attendance {
	result := db.Table("attendance").Select("*").Create(&ent)
	if result.Error != nil {
		ent.ID = 0
		return ent
	}

	return ent
}

// UpdateAttendance implements AttendanceRepository
func (*AttendanceRepositoryImpl) UpdateAttendance(db *gorm.DB, ent *entity.Attendance) *entity.Attendance {
	result := db.Table("attendance").Where("uuid_user = ?", ent.UserId).Where("date = ?", ent.Date).Updates(&ent)
	if result.Error != nil {
		///handle panic
		panic(result.Error)
	}

	return ent
}

// CheckAttendance implements AttendanceRepository
func (*AttendanceRepositoryImpl) CheckAttendance(db *gorm.DB, userid string, date string) *entity.Attendance {
	var attendance = entity.Attendance{}
	result := db.Table("attendance").Select("*").Where("uuid_user = ?", userid).Where("date = ?", date).Scan(&attendance)
	if result.Error != nil {
		attendance.ID = 0
		return &attendance
	}
	return &attendance
}
