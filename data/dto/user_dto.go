package dto

type UserDTO struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	SignupWith  int    `json:"signup_with"`
	Created     string `json:"created_at"`
	Updated     string `json:"updated_at"`
}

type UserCreateDTO struct {
	Username    string `validate:"required,min=1" json:"username"`
	Email       string `json:"email"`
	Password    string `validate:"required,min=8" json:"password"`
	PhoneNumber string `json:"phone_number"`
	SignupWith  int    `validate:"required" json:"signup_with"`
	Created     string `validate:"required" json:"created_at"`
	Updated     string `validate:"required" json:"updated_at"`
}

type UserResponseDTO struct {
	Id          int    `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	SignupWith  int    `json:"signup_with"`
	Created     string `json:"created_at"`
	Updated     string `json:"updated_at"`
	Error       bool   `json:"-"`
	Message     string `json:"-"`
}
