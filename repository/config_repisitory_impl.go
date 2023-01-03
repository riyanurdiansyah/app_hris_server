package repository

import (
	"app-hris-server/data/entity"
	"app-hris-server/helper"

	"gorm.io/gorm"
)

type ConfigRepositoryImpl struct {
}

func NewConfigRepository() ConfigRepository {
	return &ConfigRepositoryImpl{}
}

// GetConfig implements ConfigeRepository
func (*ConfigRepositoryImpl) GetConfig(db *gorm.DB) []*entity.Config {
	var configs = []*entity.Config{}
	result :=
		db.Table("config").Select("*").Scan(&configs)
	helper.PanicIfError(result.Error)
	return configs
}

// GetConfigByName implements ConfigeRepository
func (repo *ConfigRepositoryImpl) GetConfigByName(db *gorm.DB, name string) *entity.Config {
	config := entity.Config{}
	result :=
		db.Table("config").Where("name= ?", name).Scan(&config)
	helper.PanicIfError(result.Error)
	return &config
}

// InsertConfig implements ConfigeRepository
func (repo *ConfigRepositoryImpl) InsertConfig(db *gorm.DB, ent *entity.Config) *entity.Config {
	result := db.Table("config").Create(&ent)
	if result.Error != nil {
		///handle panic
		panic(result.Error)
	}
	return ent
}

// UpdateConfig implements ConfigeRepository
func (repo *ConfigRepositoryImpl) UpdateConfig(db *gorm.DB, ent *entity.Config) *entity.Config {
	var userInfo = entity.UserPersonalInfo{}
	result :=
		db.Table("config").Where("id= ?", ent.ID).Updates(&userInfo)
	if result.Error != nil {
		///handle panic
		panic(result.Error)
	}
	return ent
}

// CheckNameConfig implements ConfigeRepository
func (*ConfigRepositoryImpl) CheckNameConfig(db *gorm.DB, name string) bool {
	var user = entity.Config{}
	db.Table("config").Select("*").Where("name = ?", name).Scan(&user)
	if user.Name == "" {
		return false
	} else {
		return true
	}
}
