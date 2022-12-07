package entity

type Absent struct {
	ID         int     `gorm:"column:id;primaryKey;autoIncrement"`
	IdUser     string  `gorm:"column:id_user"`
	IdEmployee string  `gorm:"column:id_employee"`
	Tanggal    string  `gorm:"column:tanggal"`
	Latitude   float64 `gorm:"column:latitude"`
	Longitude  float64 `gorm:"column:longitude"`
	Catatan    string  `gorm:"column:catatan"`
	Tipe       int     `gorm:"column:tipe"`
	Photo      string  `gorm:"column:photo"`
	CreatedAt  string  `gorm:"column:created_at"`
	UpdatedAt  string  `gorm:"column:updated_at"`
}
