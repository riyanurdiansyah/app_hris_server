package entity

type Menu struct {
	ID     int    `gorm:"column:id;primaryKey;autoIncrement"`
	Title  string `gorm:"column:title"`
	Image  string `gorm:"column:image"`
	Status int    `gorm:"column:status"`
	Route  string `gorm:"column:route"`
}
