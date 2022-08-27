package dto

type CategoryDTO struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Created string `json:"created_at"`
	Updated string `json:"updated_at"`
}

type CategoryCreateDTO struct {
	Name    string `validate:"required,min=1" json:"name"`
	Created string `json:"created_at"`
	Updated string `json:"updated_at"`
}

type CategoryUpdateDTO struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Created string `json:"created_at"`
	Updated string `json:"updated_at"`
}

type CategoryResponseDTO struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Created string `json:"created_at"`
	Updated string `json:"updated_at"`
	Error   bool   `json:"-"`
	Message string `json:"-"`
}
