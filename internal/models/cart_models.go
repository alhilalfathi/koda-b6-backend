package models

type Cart struct {
	Id        int    `json:"cart_id" db:"id"`
	Quantity  int    `json:"quantity" db:"quantity"`
	Size      string `json:"size" db:"size"`
	Variant   string `json:"variant" db:"variant"`
	UserId    int    `json:"user_id" db:"user_id"`
	ProductId int    `json:"product_id" db:"product_id"`
}

type CreateCartRequest struct {
	Quantity  int    `json:"quantity" db:"quantity"`
	Size      string `json:"size" db:"size"`
	Variant   string `json:"variant" db:"variant"`
	UserId    int    `json:"user_id" db:"user_id"`
	ProductId int    `json:"product_id" db:"product_id"`
}

type UpdateCartRequest struct {
	Quantity  int    `json:"quantity" db:"quantity"`
	Size      string `json:"size" db:"size"`
	Variant   string `json:"variant" db:"variant"`
	UserId    int    `json:"user_id" db:"user_id"`
	ProductId int    `json:"product_id" db:"product_id"`
}
