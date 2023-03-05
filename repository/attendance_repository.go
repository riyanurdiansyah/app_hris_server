package repository

import (
	"app-hris-server/data/entity"

	"gorm.io/gorm"
)

type AttendanceRepository interface {
	Attendance(db *gorm.DB, ent *entity.Attendance) *entity.Attendance
	UpdateAttendance(db *gorm.DB, ent *entity.Attendance) *entity.Attendance
	CheckAttendance(db *gorm.DB, userid string, date string) *entity.Attendance
}
