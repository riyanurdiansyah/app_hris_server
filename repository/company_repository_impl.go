package repository

import (
	"app-hris-server/data/entity"

	"gorm.io/gorm"
)

type CompanyRepositoryImpl struct {
}

func NewCompanyRepository() CompanyRepository {
	return &CompanyRepositoryImpl{}
}

// InsertCategory implements CompanyRepository
func (*CompanyRepositoryImpl) InsertCompany(db *gorm.DB, ent *entity.Company) *entity.Company {
	result := db.Table("companies").Select("*").Create(&ent)
	if result.Error != nil {
		///handle panic
		panic(result.Error)
	}
	return ent
}
