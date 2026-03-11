package models

type ForgotPass struct {
	Id        int    `json:"forgot_pass_id" db:"id"`
	CreatedAt string `json:"created_at" db:"created_at"`
	Email     string `json:"email" db:"email"`
	Code      int    `json:"code" db:"code"`
}

type CreateForgotPassRequest struct {
	Email string `json:"email" db:"email"`
}

type UpdateForgotPassRequest struct {
	Email    string `json:"email" db:"email"`
	Code     int    `json:"code" db:"code"`
	Password string `json:"password"`
}
