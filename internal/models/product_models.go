package models

type Product struct {
	Id          int    `json:"product-id"`
	ProductName string `json:"product-name"`
	Desc        string `json:"product-desc"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
}
type CreateProductRequest struct {
	ProductName string `json:"product-name"`
	Desc        string `json:"product-desc"`
	Price       int    `json:"price"`
	Category    string `json:"category"`
	Stock       int    `json:"stock"`
}
type UpdateProductRequest struct {
	ProductName string `json:"product-name"`
	Desc        string `json:"product-desc"`
	Price       int    `json:"price"`
	Category    string `json:"category"`
	Stock       int    `json:"stock"`
}
