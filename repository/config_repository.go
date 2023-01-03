package repository

import (
	"app-hris-server/data/entity"

	"gorm.io/gorm"
)

type ConfigRepository interface {
	InsertConfig(db *gorm.DB, ent *entity.Config) *entity.Config
	UpdateConfig(db *gorm.DB, ent *entity.Config) *entity.Config
	GetConfig(db *gorm.DB) []*entity.Config
	GetConfigByName(db *gorm.DB, name string) *entity.Config
	CheckNameConfig(db *gorm.DB, name string) bool
}
