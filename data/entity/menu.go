package entity

type Menu struct {
	ID          int    `gorm:"column:id;primaryKey;autoIncrement"`
	Category    string `gorm:"column:category"`
	Position    int    `gorm:"column:position"`
	Title       string `gorm:"column:title"`
	Image       string `gorm:"column:image"`
	Status      int    `gorm:"column:status"`
	Language    string `gorm:"column:language"`
	Route       string `gorm:"column:route"`
	Description string `gorm:"column:description"`
}
