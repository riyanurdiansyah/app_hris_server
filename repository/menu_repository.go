package repository

import (
	"app-hris-server/data/entity"

	"gorm.io/gorm"
)

type MenuRepository interface {
	InsertMenu(db *gorm.DB, ent *entity.Menu) *entity.Menu
	UpdateMenu(db *gorm.DB, ent *entity.Menu) *entity.Menu
	GetMenuById(db *gorm.DB, id int) *entity.Menu
	GetMenu(db *gorm.DB) []*entity.Menu
	CheckMenu(db *gorm.DB, id int) bool
}
