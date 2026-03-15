package models

type ProductSize struct {
	Id        int `json:"product_category_id" db:"id"`
	ProductId int `json:"product_id" db:"product_id"`
	SizeId    int `json:"category_id" db:"category_id"`
}

type CreateProductSizeRequest struct {
	ProductId int `json:"product_id" db:"product_id"`
	SizeId    int `json:"size_id" db:"size_id"`
}
type UpdateProductSizeRequest struct {
	ProductId int `json:"product_id" db:"product_id"`
	SizeId    int `json:"size_id" db:"size_id"`
}
