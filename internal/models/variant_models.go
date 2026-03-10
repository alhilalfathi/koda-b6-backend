package models

type Variant struct {
	Id         int    `json:"variant_id" db:"id"`
	Variant    string `json:"variant" db:"variant"`
	AddedPrice int    `json:"add_price" db:"add_price"`
}

type CreateVariantRequest struct {
	Variant    string `json:"variant" db:"variant"`
	AddedPrice int    `json:"add_price" db:"add_price"`
}

type UpdateVariantRequest struct {
	Variant    string `json:"variant" db:"variant"`
	AddedPrice int    `json:"add_price" db:"add_price"`
}
