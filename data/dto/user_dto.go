package dto

type UserDTO struct {
	ID          int    `json:"id"`
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
	Username    string `validate:"required,min=1" json:"username"`
	Email       string `json:"email"`
	Password    string `validate:"required,min=8" json:"password"`
	PhoneNumber string `json:"phone_number"`
	SignupWith  int    `validate:"required" json:"register_by"`
	Role        int    `validate:"required" json:"role"`
	Created     string `json:"created_at"`
	Updated     string `json:"updated_at"`
}

type UserLoginDTO struct {
	Email    string `validate:"required,min=1" json:"email"`
	Username string `validate:"required,min=1" json:"username"`
	Password string `validate:"required,min=8" json:"password"`
}

type UserResponseDTO struct {
	Id          int    `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"-"`
	PhoneNumber string `json:"phone_number"`
	SignupWith  int    `json:"register_by"`
	Role        int    `json:"role"`
	Created     string `json:"created_at"`
	Updated     string `json:"updated_at"`
	Error       bool   `json:"-"`
	Message     string `json:"-"`
}
