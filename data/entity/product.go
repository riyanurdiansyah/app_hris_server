package entity

type Product struct {
	ID           int    `gorm:"column:id;primaryKey;autoIncrement"`
	Name         string `gorm:"column:name"`
	Condition    string `gorm:"column:condition"`
	Price        int32  `gorm:"column:price"`
	Discount     int32  `gorm:"column:discount"`
	DiscountType int8   `gorm:"column:discount_type"`
	Weight       int32  `gorm:"column:weight"`
	IdCategory   int8   `gorm:"column:id_category"`
	Description  string `gorm:"column:description"`
	Minimum      int8   `gorm:"column:minimum"`
	Status       int8   `gorm:"column:status"`
	CreatedAt    string `gorm:"column:created_at"`
	UpdatedAt    string `gorm:"column:updated_at"`
}
