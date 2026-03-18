package models

type Users struct {
	Id       int    `json:"id" db:"id"`
	FullName string `json:"fullname" db:"fullname"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type CreateUserRequest struct {
	FullName string `json:"fullname" db:"fullname"`
	Email    string `json:"email" db:"email" binding:"required,email"`
	Password string `json:"password" db:"password" binding:"required,min=4"`
}
type LoginUserRequest struct {
	Email    string `json:"email" db:"email" binding:"required,email"`
	Password string `json:"password" db:"password" binding:"required,min=4"`
}
type LoginUserResponse struct {
	Token    string `json:"token"`
	Fullname string `json:"fullname"`
}
type UpdateUserRequest struct {
	FullName string `json:"fullname" db:"fullname"`
	Email    string `json:"email" db:"email" binding:"required,email"`
	Password string `json:"password" db:"password" binding:"required,min=4"`
}
