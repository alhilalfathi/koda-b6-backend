package models

type ProductCategory struct {
	Id         int `json:"product_category_id" db:"id"`
	ProductId  int `json:"product_id" db:"product_id"`
	CategoryId int `json:"category_id" db:"category_id"`
}

type CreateProductCategoryRequest struct {
	ProductId  int `json:"product_id" db:"product_id"`
	CategoryId int `json:"category_id" db:"category_id"`
}
type UpdateProductCategoryRequest struct {
	ProductId  int `json:"product_id" db:"product_id"`
	CategoryId int `json:"category_id" db:"category_id"`
}
