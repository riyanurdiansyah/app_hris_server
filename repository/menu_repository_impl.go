package repository

import (
	"app-hris-server/data/entity"

	"gorm.io/gorm"
)

type MenuRepositoryImpl struct{}

func NewMenuRepository() MenuRepository {
	return &MenuRepositoryImpl{}
}

// CheckMenu implements MenuRepository
func (*MenuRepositoryImpl) CheckMenu(db *gorm.DB, id int) bool {
	var menu = entity.Menu{}
	db.Table("menu").Select("*").Where("id = ?", id).Scan(&menu)
	if menu.Title == "" {
		return false
	} else {
		return true
	}
}

// InsertMenu implements MenuRepository
func (*MenuRepositoryImpl) InsertMenu(db *gorm.DB, ent *entity.Menu) *entity.Menu {
	result := db.Table("menu").Select("*").Create(&ent)
	if result.Error != nil {
		///handle panic
		panic(result.Error)
	}
	return ent
}

// UpdateMenu implements MenuRepository
func (*MenuRepositoryImpl) UpdateMenu(db *gorm.DB, ent *entity.Menu) *entity.Menu {

	var menu = entity.Menu{}
	result :=
		db.Table("menu").Where("id= ?", ent.ID).Updates(&menu)
	if result.Error != nil {
		///handle panic
		panic(result.Error)
	}
	return ent
}

// GetMenu implements MenuRepository
func (*MenuRepositoryImpl) GetMenu(db *gorm.DB) []*entity.Menu {
	var menu []*entity.Menu
	return menu
}
