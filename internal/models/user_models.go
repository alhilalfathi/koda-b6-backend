package models

type Users struct {
	Id       int    `json:"id"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserRequest struct {
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type LoginUserResponse struct {
	Token string `json:"token"`
}
type UpdateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
