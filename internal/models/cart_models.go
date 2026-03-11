package models

type Cart struct {
	Id        int `json:"cart_id" db:"id"`
	Quantity  int `json:"quantity" db:"quantity"`
	SizeId    int `json:"size_id" db:"size_id"`
	VariantId int `json:"variant_id" db:"variant_id"`
	UserId    int `json:"user_id" db:"user_id"`
	ProductId int `json:"product_id" db:"product_id"`
}

type CreateCartRequest struct {
	Quantity  int `json:"quantity" db:"quantity"`
	SizeId    int `json:"size_id" db:"size_id"`
	VariantId int `json:"variant_id" db:"variant_id"`
	UserId    int `json:"user_id" db:"user_id"`
	ProductId int `json:"product_id" db:"product_id"`
}

type UpdateCartRequest struct {
	Quantity  int `json:"quantity" db:"quantity"`
	SizeId    int `json:"size_id" db:"size_id"`
	VariantId int `json:"variant_id" db:"variant_id"`
	UserId    int `json:"user_id" db:"user_id"`
	ProductId int `json:"product_id" db:"product_id"`
}
