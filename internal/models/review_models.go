package models

type Review struct {
	Id        int    `json:"review_id" db:"id"`
	UserId    int    `json:"user_id" db:"user_id"`
	ProductId int    `json:"product_id" db:"product_id"`
	Messages  string `json:"messages" db:"messages"`
	Rating    int    `json:"rating" db:"rating"`
	Path      string `json:"path" db:"path"`
}

type CreateReviewRequest struct {
	UserId    int    `json:"user_id" db:"user_id"`
	ProductId int    `json:"product_id" db:"product_id"`
	Messages  string `json:"messages" db:"messages"`
	Rating    int    `json:"rating" db:"rating"`
}

type UpdateReviewRequest struct {
	UserId    int    `json:"user_id" db:"user_id"`
	ProductId int    `json:"product_id" db:"product_id"`
	Messages  string `json:"messages" db:"messages"`
	Rating    int    `json:"rating" db:"rating"`
}
