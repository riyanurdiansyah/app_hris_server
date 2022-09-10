package entity

type Category struct {
	ID        int    `gorm:"column:id;primaryKey;autoIncrement"`
	Name      string `gorm:"column:name"`
	Image     string `gorm:"column:image"`
	Status    int    `gorm:"column:status"`
	CreatedAt string `gorm:"column:created_at"`
	UpdatedAt string `gorm:"column:updated_at"`
}
