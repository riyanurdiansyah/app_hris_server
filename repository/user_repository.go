package repository

import (
	"app-hris-server/data/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	UpdateUserInfoPersonal(db *gorm.DB, ent *entity.UserPersonalInfo) *entity.UserPersonalInfo
	AddUserInfoPersonal(db *gorm.DB, ent *entity.UserPersonalInfo) *entity.UserPersonalInfo
	CheckUser(db *gorm.DB, employeeId string) bool
}
