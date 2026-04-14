package models

import "time"

type Transaction struct {
	Id          int       `json:"transaction_id" db:"id"`
	TrxId       string    `json:"trx_id" db:"trx_id"`
	UserId      int       `json:"user_id" db:"user_id"`
	OrderDate   time.Time `json:"order_date" db:"order_date"`
	Fullname    string    `json:"fullname" db:"fullname"`
	Email       string    `json:"email" db:"email"`
	Address     string    `json:"address" db:"address"`
	Delivery    string    `json:"delivery" db:"delivery"`
	DeliveryFee int       `json:"delivery_fee" db:"delivery_fee"`
	Tax         int       `json:"tax" db:"tax"`
	Total       int       `json:"total" db:"total"`
	OrderStatus string    `json:"status_order" db:"status_order"`
}

type CreateTransactionRequest struct {
	TrxId       string `json:"trx_id" db:"trx_id"`
	UserId      int    `json:"user_id" db:"user_id"`
	Fullname    string `json:"fullname" db:"fullname"`
	Email       string `json:"email" db:"email"`
	Address     string `json:"address" db:"address"`
	Delivery    string `json:"delivery" db:"delivery"`
	DeliveryFee int    `json:"delivery_fee" db:"delivery_fee"`
	Tax         int    `json:"tax" db:"tax"`
	Total       int    `json:"total" db:"total"`
	OrderStatus string `json:"status_order" db:"status_order"`
}

type UpdateTransactionRequest struct {
	TrxId       string `json:"trx_id" db:"trx_id"`
	UserId      int    `json:"user_id" db:"user_id"`
	Fullname    string `json:"fullname" db:"fullname"`
	Email       string `json:"email" db:"email"`
	Address     string `json:"address" db:"address"`
	Delivery    string `json:"delivery" db:"delivery"`
	DeliveryFee int    `json:"delivery_fee" db:"delivery_fee"`
	Tax         int    `json:"tax" db:"tax"`
	Total       int    `json:"total" db:"total"`
	OrderStatus string `json:"status_order" db:"status_order"`
}

type TransactionItem struct {
	ProductId   int    `json:"product_id"`
	ProductName string `json:"product_name"`
	Price       int    `json:"price"`
	Qty         int    `json:"qty"`
	Image       string `json:"image"`
}

type TransactionDetailResponse struct {
	Id          int               `json:"id"`
	TrxId       string            `json:"trx_id"`
	OrderDate   time.Time         `json:"order_date"`
	Total       int               `json:"total"`
	StatusOrder string            `json:"status_order"`
	Items       []TransactionItem `json:"items"`
}

type TransactionHistoryResponse struct {
	Id          int    `json:"id"`
	TrxId       string `json:"trx_id"`
	OrderDate   string `json:"order_date"`
	Total       int    `json:"total"`
	StatusOrder string `json:"status_order"`
}
