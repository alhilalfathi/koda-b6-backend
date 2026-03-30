package models

type Product struct {
	Id          int    `json:"product_id" db:"id"`
	ProductName string `json:"product_name" db:"product_name"`
	Desc        string `json:"product_desc" db:"product_desc"`
	Price       int    `json:"price" db:"price"`
	Stock       int    `json:"stock" db:"stock"`
	Category    string `json:"category" db:"category"`
	Images      string `json:"path" db:"path"`
}
type CreateProductRequest struct {
	ProductName string `json:"product_name" db:"product_name"`
	Desc        string `json:"product_desc" db:"product_desc"`
	Price       int    `json:"price" db:"price"`
	Stock       int    `json:"stock" db:"stock"`
}
type UpdateProductRequest struct {
	ProductName string `json:"product_name" db:"product_name"`
	Desc        string `json:"product_desc" db:"product_desc"`
	Price       int    `json:"price" db:"price"`
	Stock       int    `json:"stock" db:"stock"`
}
type RecommendedProduct struct {
	Id          int    `json:"product_id" db:"id"`
	ProductName string `json:"product_name" db:"product_name"`
	Desc        string `json:"product_desc" db:"product_desc"`
	Price       int    `json:"price" db:"price"`
	Stock       int    `json:"stock" db:"stock"`
	Images      string `json:"path" db:"path"`
	CountReview int    `db:"total_review"`
}
