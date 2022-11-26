package repository

import (
	"app-hris-server/data/entity"

	"gorm.io/gorm"
)

type CompanyRepository interface {
	InsertCompany(db *gorm.DB, ent *entity.Company) *entity.Company
}
