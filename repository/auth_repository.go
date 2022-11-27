package repository

import (
	"app-hris-server/data/entity"

	"gorm.io/gorm"
)

type AuthRepository interface {
	SignUp(db *gorm.DB, user *entity.User) *entity.User
	FindUserByEmail(db *gorm.DB, email string) *entity.User
	FindUserByUsername(db *gorm.DB, username string) *entity.User
	CheckEmail(db *gorm.DB, email string) bool
	CheckUsername(db *gorm.DB, username string) bool
	CheckCompany(db *gorm.DB, key string) bool
}
