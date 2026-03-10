package models

type Category struct {
	Id       int    `json:"category_id" db:"id"`
	Category string `json:"category" db:"category"`
}

type CreateCategoryRequest struct {
	Category string `json:"category" db:"category"`
}

type UpdateCategoryRequest struct {
	Category string `json:"category" db:"category"`
}
