package repository

import (
	"app-ecommerce-server/data/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
}

func NewAuthRepository() AuthRepository {
	return &AuthRepositoryImpl{}
}

// SignUp implements AuthRepository
func (repo *AuthRepositoryImpl) SignUp(c *gin.Context, db *gorm.DB, user *entity.User) *entity.User {
	result := db.Table("users").Select("*").Create(&user)
	if result.Error != nil {
		user.ID = -99
		return user
	}

	return user
}

// CheckEmail implements AuthRepository
func (*AuthRepositoryImpl) CheckEmail(c *gin.Context, db *gorm.DB, email string) bool {
	var user = entity.User{}
	db.Table("users").Select("*").Where("email = ?", email).Scan(&user)
	if user.Email == "" {
		return false
	} else {
		return true
	}
}

// CheckUsername implements AuthRepository
func (*AuthRepositoryImpl) CheckUsername(c *gin.Context, db *gorm.DB, username string) bool {
	var user = entity.User{}
	db.Table("users").Select("*").Where("username = ?", username).Scan(&user)
	if user.Username == "" {
		return false
	} else {
		return true
	}
}
