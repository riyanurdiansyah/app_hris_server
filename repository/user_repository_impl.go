package repository

import (
	"app-hris-server/data/entity"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

// UpdatePersonalInfo implements UserRepository
func (*UserRepositoryImpl) AddUserInfoPersonal(db *gorm.DB, ent *entity.UserPersonalInfo) *entity.UserPersonalInfo {
	result := db.Table("user_info_personal").Select("*").Create(&ent)
	if result.Error != nil {
		///handle panic
		panic(result.Error)
	}
	return ent
}

// UpdatePersonalInfo implements UserRepository
func (*UserRepositoryImpl) UpdateUserInfoPersonal(db *gorm.DB, ent *entity.UserPersonalInfo) *entity.UserPersonalInfo {
	var userInfo = entity.UserPersonalInfo{}
	result :=
		db.Table("categories").Where("id = ?", ent.IdEmployee).Updates(&userInfo)
	if result.Error != nil {
		///handle panic
		panic(result.Error)
	}
	return ent
}

// CheckUser implements UserRepository
func (*UserRepositoryImpl) CheckUser(db *gorm.DB, employeeId string) bool {
	var user = entity.UserPersonalInfo{}
	db.Table("user_info_personal").Select("*").Where("id_employee = ?", employeeId).Scan(&user)
	if user.IdEmployee == "" {
		return false
	} else {
		return true
	}
}
