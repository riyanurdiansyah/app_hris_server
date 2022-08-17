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
