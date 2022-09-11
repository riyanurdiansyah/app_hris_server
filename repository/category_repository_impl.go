package repository

import (
	"app-ecommerce-server/data/entity"
	"app-ecommerce-server/helper"

	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) InsertCategory(db *gorm.DB, category *entity.Category) *entity.Category {
	result := db.Table("categories").Select("*").Create(&category)
	if result.Error != nil {
		///handle panic
		panic(result.Error)
	}
	return category
}

func (repository *CategoryRepositoryImpl) FindAllCategory(db *gorm.DB, page int) ([]*entity.Category, int64) {
	var listCategory = []*entity.Category{}
	var total int64
	if page == 0 {
		result :=
			db.Table("categories").Select("*").Scan(&listCategory).Count(&total)
		helper.PanicIfError(result.Error)
	} else {
		perPage := 4
		result :=
			db.Table("categories").Select("*").Limit(4).Offset(((page - 1) * perPage)).Scan(&listCategory).Count(&total)
		helper.PanicIfError(result.Error)
	}
	db.Table("categories").Select("*").Count(&total)
	return listCategory, total
}

func (repository *CategoryRepositoryImpl) FindByIdCategory(db *gorm.DB, categoryId int) *entity.Category {
	var category = entity.Category{}
	result :=
		db.Table("categories").Select("*").Where("id = ?", categoryId).Scan(&category)
	helper.PanicIfError(result.Error)
	return &category
}

func (repository *CategoryRepositoryImpl) DeleteCategory(db *gorm.DB, category *entity.Category) *entity.Category {
	result :=
		db.Table("categories").Where("id = ?", category.ID).Delete(&category)
	if result.Error != nil {
		category.ID = -99
		return category
	}
	return category
}

func (repository *CategoryRepositoryImpl) UpdateCategory(db *gorm.DB, category *entity.Category) *entity.Category {
	var count int64
	var tempCategory = entity.Category{}
	checkid :=
		db.Table("categories").Select("*").Where("id = ?", category.ID).Count(&count).Scan(&tempCategory)
	if checkid.Error != nil {
		///handle error
		/// ID Set ke -2 untuk penanda error saat check id
		category.ID = -2
		return category
	}
	if count > 0 {
		result :=
			db.Table("categories").Where("id = ?", category.ID).Updates(&category)
		if result.Error != nil {
			///handle error
			/// ID Set ke -2 untuk penanda error result
			category.ID = -2
			return category
		}
		return category
	} else {
		///handle error
		/// ID Set ke -2 untuk penanda id not found
		category.ID = -1
		return category
	}

}
