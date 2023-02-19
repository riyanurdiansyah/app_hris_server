package entity

type Task struct {
	ID           int     `gorm:"column:id;primaryKey;autoIncrement"`
	Uuid         string  `gorm:"column:uuid_user"`
	IdUser       string  `gorm:"column:id_user"`
	Title        string  `gorm:"column:title"`
	NamaDepan    string  `gorm:"column:nama_depan"`
	NamaBelakang string  `gorm:"column:nama_belakang"`
	Image        string  `gorm:"column:image"`
	NoHp         string  `gorm:"column:no_hp"`
	Status       int     `gorm:"column:status"`
	Progress     float64 `gorm:"column:progress"`
}

type TaskBy struct {
	NamaDepan    string
	NamaBelakang string
	Image        string
	NoHp         string
}
