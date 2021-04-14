package models

import (
	"errors"
	"github.com/friends-management/helper"
)

type Friend struct {
	User1 int
	User2 int
}

type FriendRequest struct {
	Friends []string `json:"friends"`
}

type FriendsListRequest struct {
	Email string `json:"email"`
}

func (r FriendsListRequest) Validate() error {
	if r.Email == "" {
		return errors.New("\"email\" is required")
	}

	if !helper.IsEmailValid(r.Email) {
		return errors.New("\"email\"'s format is not valid. (ex: \"andy@example\")")
	}

	return nil
}

type FriendsResponse struct {
	Success bool     `json:"success"`
	Friends []string `json:"friends"`
	Count   int      `json:"count"`
}

func (r FriendRequest) Validate() error {
	emailUser1 := r.Friends[0]
	emailUser2 := r.Friends[1]

	if emailUser1 == "" {
		return errors.New("your email is required")
	}

	if emailUser2 == "" {
		return errors.New("your friend is required")
	}

	if !helper.IsEmailValid(emailUser1) || !helper.IsEmailValid(emailUser2) {
		return errors.New("\"email\"'s format is not valid. (ex: \"andy@example\")")
	}

	return nil
}

type CommonFriendsRequest struct {
	Friends []string `json:"friends"`
}

func (r CommonFriendsRequest) Validate() error {
	emailUser1 := r.Friends[0]
	emailUser2 := r.Friends[1]

	if emailUser1 == "" {
		return errors.New("your email is required")
	}

	if emailUser2 == "" {
		return errors.New("your friend is required")
	}

	if !helper.IsEmailValid(emailUser1) || !helper.IsEmailValid(emailUser2) {
		return errors.New("\"email\"'s format is not valid. (ex: \"andy@example\")")
	}

	return nil
}

type ReceiveUpdateRequest struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

func (r ReceiveUpdateRequest) Validate() interface{} {
	if r.Sender == "" {
		return errors.New("sender is required")
	}

	if !helper.IsEmailValid(r.Sender) {
		return errors.New("\"email\"'s format is not valid. (ex: \"andy@example\")")
	}
	return nil
}

type ReceiveUpdateResponse struct {
	Success    bool     `json:"success"`
	Recipients []string `json:"recipients"`
}
