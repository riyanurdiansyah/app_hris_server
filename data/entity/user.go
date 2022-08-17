package entity

type User struct {
	ID          int    `gorm:"column:id;primaryKey;autoIncrement"`
	Username    string `gorm:"column:username"`
	Email       string `gorm:"column:email"`
	Password    string `gorm:"column:password"`
	PhoneNumber string `gorm:"column:phone_number"`
	SignupWith  int    `gorm:"column:signup_with"`
	CreatedAt   string `gorm:"column:created_at"`
	UpdatedAt   string `gorm:"column:updated_at"`
}
