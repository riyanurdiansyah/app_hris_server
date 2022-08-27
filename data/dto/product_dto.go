package dto

type ProductDTO struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Condition    string `json:"condition"`
	Price        int32  `json:"price"`
	Discount     int32  `json:"discount"`
	DiscountType int8   `json:"discount_type"`
	Weight       int32  `json:"weight"`
	IdCategory   int8   `json:"category_id"`
	Description  string `json:"description"`
	Minimum      int8   `json:"minimum"`
	Status       int8   `json:"status"`
	Created      string `json:"created_at"`
	Updated      string `json:"updated_at"`
}
