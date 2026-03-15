package models

type TransactionProduct struct {
	Id        int    `json:"id" db:"id"`
	TrId      int    `json:"transaction_id" db:"transaction_id"`
	ProductId int    `json:"product_id" db:"Product_id"`
	Quantity  int    `json:"quantity" db:"quantity"`
	Size      string `json:"size" db:"size"`
	Variant   string `json:"variant" db:"variant"`
}

type CreateTransactionProductRequest struct {
	TrId      int    `json:"transaction_id" db:"transaction_id"`
	ProductId int    `json:"product_id" db:"Product_id"`
	Quantity  int    `json:"quantity" db:"quantity"`
	Size      string `json:"size" db:"size"`
	Variant   string `json:"variant" db:"variant"`
}

type UpdateTransactionProductRequest struct {
	TrId      int    `json:"transaction_id" db:"transaction_id"`
	ProductId int    `json:"product_id" db:"Product_id"`
	Quantity  int    `json:"quantity" db:"quantity"`
	Size      string `json:"size" db:"size"`
	Variant   string `json:"variant" db:"variant"`
}
