package entity

type Promo struct {
	ID          int    `gorm:"column:id;primaryKey;autoIncrement"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
	Image       string `gorm:"column:image"`
	KodePromo   string `gorm:"column:kode_promo"`
	Expired     int    `gorm:"column:expired"`
	Status      int    `gorm:"column:status"`
	CreatedAt   string `gorm:"column:created_at"`
	UpdatedAt   string `gorm:"column:updated_at"`
}
