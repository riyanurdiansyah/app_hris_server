package repository

import (
	"app-hris-server/data/entity"
	"fmt"

	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
}

func NewAuthRepository() AuthRepository {
	return &AuthRepositoryImpl{}
}

// SignUp implements AuthRepository
func (repo *AuthRepositoryImpl) SignUp(db *gorm.DB, user *entity.User) *entity.User {
	result := db.Table("users").Select("*").Create(&user)
	if result.Error != nil {
		user.ID = -99
		return user
	}

	return user
}

// FindUserByEmail implements AuthRepository
func (*AuthRepositoryImpl) FindUserByEmail(db *gorm.DB, email string) *entity.User {
	fmt.Println("KOCAK 2", email)
	var user = entity.User{}
	result := db.Table("users").Select("*").Where("email = ?", email).Scan(&user)
	if result.Error != nil {
		user.ID = -99
		return &user
	}
	return &user
}

// FindUserByUsername implements AuthRepository
func (*AuthRepositoryImpl) FindUserByUsername(db *gorm.DB, username string) *entity.User {
	var user = entity.User{}
	result := db.Table("users").Select("*").Where("username = ?", username).Scan(&user)
	if result.Error != nil {
		user.ID = -99
		return &user
	}
	return &user
}

// CheckEmail implements AuthRepository
func (*AuthRepositoryImpl) CheckEmail(db *gorm.DB, email string) bool {
	var user = entity.User{}
	db.Table("users").Select("*").Where("email = ?", email).Scan(&user)
	if user.Email == "" {
		return false
	} else {
		return true
	}
}

// CheckUsername implements AuthRepository
func (*AuthRepositoryImpl) CheckUsername(db *gorm.DB, username string) bool {
	var user = entity.User{}
	db.Table("users").Select("*").Where("username = ?", username).Scan(&user)
	if user.Username == "" {
		return false
	} else {
		return true
	}
}

// CheckCompany implements AuthRepository
func (*AuthRepositoryImpl) CheckCompany(db *gorm.DB, key string) bool {
	var company = entity.Company{}
	db.Table("companies").Select("*").Where("secret_key = ?", key).Scan(&company)
	if company.Name == "" {
		return false
	} else {
		return true
	}
}
