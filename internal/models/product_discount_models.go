package models

type ProductDiscount struct {
	Id         int `json:"product_discount_id" db:"id"`
	ProductId  int `json:"product_id" db:"product_id"`
	DiscountId int `json:"discount_id" db:"discount_id"`
}

type CreateProductDiscountRequest struct {
	ProductId  int `json:"product_id" db:"product_id"`
	DiscountId int `json:"discount_id" db:"discount_id"`
}
type UpdateProductDiscountRequest struct {
	ProductId  int `json:"product_id" db:"product_id"`
	DiscountId int `json:"discount_id" db:"discount_id"`
}
