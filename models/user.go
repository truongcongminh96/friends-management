package models

type User struct {
	Email string `json:"email"`
}

type UserListResponse struct {
	Users []User `json:"users"`
}

type UserRequest struct {
	Email string `json:"email"`
}

type ResultResponse struct {
	Success bool `json:"success"`
}
