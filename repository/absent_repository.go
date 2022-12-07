package repository

import (
	"app-hris-server/data/dto"
	"app-hris-server/data/entity"

	"gorm.io/gorm"
)

type AbsentRepository interface {
	InsertAbsent(db *gorm.DB, absent *entity.Absent) *entity.Absent
	CheckIfDoneAbsent(db *gorm.DB, absent *dto.AbsentCreateDTO) bool
}
