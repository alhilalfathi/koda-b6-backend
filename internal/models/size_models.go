package models

type Size struct {
	Id         int    `json:"size_id" db:"id"`
	Size       string `json:"size" db:"size"`
	AddedPrice int    `json:"add_price" db:"add_price"`
}

type CreateSizeRequest struct {
	Size       string `json:"size" db:"size"`
	AddedPrice int    `json:"add_price" db:"add_price"`
}

type UpdateSizeRequest struct {
	Size       string `json:"size" db:"size"`
	AddedPrice int    `json:"add_price" db:"add_price"`
}
