package entity

type SubCategoy struct {
	ID         uint64   `gorm:"column:id;primaryKey;autoIncrement"`
	Name       string   `gorm:"column:name"`
	IdCategory uint64   `gorm:"column:id_category"`
	CreatedAt  string   `gorm:"column:created_at"`
	UpdatedAt  string   `gorm:"column:updated_at"`
	Categoy    Category `gorm:"foreignKey:IdCategory"`
}
