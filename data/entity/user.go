package entity

type User struct {
	ID               int    `gorm:"column:id;primaryKey;autoIncrement"`
	EmployeeId       string `gorm:"column:employee_id"`
	Username         string `gorm:"column:username"`
	Email            string `gorm:"column:email"`
	Password         string `gorm:"column:password"`
	PhoneNumber      string `gorm:"column:phone_number"`
	Role             int    `gorm:"column:role"`
	CompanySecretKey string `gorm:"column:company_secret_key"`
	CreatedAt        string `gorm:"column:created_at"`
	UpdatedAt        string `gorm:"column:updated_at"`
}
