package models

type ProductVariant struct {
	Id        int `json:"product_category_id" db:"id"`
	ProductId int `json:"product_id" db:"product_id"`
	VariantId int `json:"category_id" db:"category_id"`
}

type CreateProductVariantRequest struct {
	ProductId int `json:"product_id" db:"product_id"`
	VariantId int `json:"variant_id" db:"variant_id"`
}
type UpdateProductVariantRequest struct {
	ProductId int `json:"product_id" db:"product_id"`
	VariantId int `json:"variant_id" db:"variant_id"`
}
