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

type CartResponse struct {
	Id          int    `json:"cart_id" db:"id"`
	ProductId   int    `json:"product_id" db:"product_id"`
	ProductName string `json:"product_name" db:"product_name"`
	Price       int    `json:"price" db:"price"`
	Images      string `json:"path" db:"path"`
	Quantity    int    `json:"quantity" db:"quantity"`
	Size        string `json:"size" db:"size"`
	Variant     string `json:"variant" db:"variant"`
	UserId      int    `json:"user_id" db:"user_id"`
}
