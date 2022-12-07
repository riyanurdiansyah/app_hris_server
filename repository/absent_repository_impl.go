package repository

import (
	"app-hris-server/data/dto"
	"app-hris-server/data/entity"

	"gorm.io/gorm"
)

type AbsentRepositoryImpl struct{}

func NewAbsentRepository() AbsentRepository {
	return &AbsentRepositoryImpl{}
}

// InsertAbsent implements AbsentRepository
func (repo *AbsentRepositoryImpl) InsertAbsent(db *gorm.DB, absent *entity.Absent) *entity.Absent {
	result := db.Table("user_absent").Select("*").Create(&absent)
	if result.Error != nil {
		///handle panic
		panic(result.Error)
	}
	return absent
}

// CheckIfDoneAbsent implements AbsentRepository
func (*AbsentRepositoryImpl) CheckIfDoneAbsent(db *gorm.DB, absent *dto.AbsentCreateDTO) bool {
	var user = entity.User{}
	db.Table("user_absent").Select("*").Where("id_user = ?", absent.IdUser).Scan(&user)
	if user.Email == "" {
		return false
	} else {
		return true
	}
}
