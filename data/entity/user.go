package entity

type User struct {
	ID          int    `gorm:"column:id;primaryKey;autoIncrement"`
	Username    string `gorm:"column:username"`
	Email       string `gorm:"column:email"`
	Password    string `gorm:"column:password"`
	PhoneNumber string `gorm:"column:phone_number"`
	SignupWith  int    `gorm:"column:signup_with"`
	Role        int    `gorm:"column:role"`
	CreatedAt   string `gorm:"column:created_at"`
	UpdatedAt   string `gorm:"column:updated_at"`
}

type UserCompany struct {
	ID        int `gorm:"column:id;primaryKey;autoIncrement"`
	IdUser    int `gorm:"column:id_user"`
	IdCompany int `gorm:"column:id_company"`
	Status    int `gorm:"column:status"`
}
