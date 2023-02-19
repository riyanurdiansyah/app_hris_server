package dto

type UserDTO struct {
	// ID          int    `json:"id"`
	Uuid        string `json:"uuid"`
	EmployeeId  string `json:"employee_id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	SignupWith  int    `json:"register_by"`
	Role        int    `json:"role"`
	Created     string `json:"created_at"`
	Updated     string `json:"updated_at"`
}

type UserCreateDTO struct {
	Username         string `validate:"required" json:"username"`
	EmployeeId       string `validate:"required" json:"employee_id"`
	Email            string `validate:"required" json:"email"`
	Password         string `validate:"required,min=8" json:"password"`
	PhoneNumber      string `json:"phone_number"`
	Role             int    `validate:"required" json:"role"`
	CompanySecretKey string `validate:"required" json:"company_secret_key"`
	Created          string `json:"created_at"`
	Updated          string `json:"updated_at"`
}

type UserLoginUsernameDTO struct {
	Username string `validate:"required,min=1" json:"username"`
	Password string `validate:"required,min=8" json:"password"`
}

type UserLoginEmailDTO struct {
	Email    string `validate:"required,min=1" json:"email"`
	Password string `validate:"required" json:"password"`
}
type UserResponseDTO struct {
	// Id               int    `json:"id"`
	Uuid             string `json:"uuid"`
	EmployeeId       string `json:"employee_id"`
	Username         string `json:"username"`
	Email            string `json:"email"`
	Password         string `json:"-"`
	PhoneNumber      string `json:"phone_number"`
	Role             int    `json:"role"`
	CompanySecretKey string `json:"company_secret_key"`
	Created          string `json:"created_at"`
	Updated          string `json:"updated_at"`
	Error            bool   `json:"-"`
	Message          string `json:"-"`
}
