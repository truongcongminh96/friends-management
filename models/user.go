package models

import (
	"errors"
	"github.com/friends-management/helper"
)

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

func (u UserRequest) Validate() error {
	if u.Email == "" {
		return errors.New("\"email\" is required")
	}

	if !helper.IsEmailValid(u.Email) {
		return errors.New("\"email\"'s format is not valid. (ex: \"andy@example\")")
	}

	return nil
}
