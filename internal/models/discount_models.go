package models

type Discount struct {
	Id          int     `json:"discount_id" db:"id"`
	Rate        float32 `json:"discount_rate" db:"discount_rate"`
	Desc        string  `json:"description" db:"description"`
	IsFlashSale bool    `json:"is_flashsale" db:"is_flashsale"`
}

type CreateDiscountRequest struct {
	Rate        float32 `json:"discount_rate" db:"discount_rate"`
	Desc        string  `json:"description" db:"description"`
	IsFlashSale *bool   `json:"is_flashsale" db:"is_flashsale"`
}

type UpdateDiscountRequest struct {
	Rate        float32 `json:"discount_rate" db:"discount_rate"`
	Desc        string  `json:"description" db:"description"`
	IsFlashSale *bool   `json:"is_flashsale" db:"is_flashsale"`
}
