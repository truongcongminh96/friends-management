package models

import (
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Email        string   `json:"email"`
	Friends      []string `json:"friends"`
	Subscription []string `json:"subscription"`
	Blocked      []string `json:"blocked"`
}

type UserList struct {
	Users []User
}

type UserListResponse struct {
	Users []User `json:"users"`
}

type UserRequest struct {
	Email string `json:"email"`
}

func (email *UserRequest) Bind(r *http.Request) error {
	if email.Email == "" {
		log.Print("email is a required field")
		return fmt.Errorf("email is a required field")
	}
	return nil
}

type ResultResponse struct {
	Success bool `json:"success"`
}

func (u UserList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type FriendConnectionRequest struct {
	Friends []string `json:"friends"`
}

func (f FriendConnectionRequest) Bind(r *http.Request) error {
	userEmailOne := f.Friends[0]
	userEmailTwo := f.Friends[1]

	if userEmailOne == "" || userEmailTwo == "" {
		return fmt.Errorf("email is a required field")
	}
	if userEmailOne == userEmailTwo {
		return fmt.Errorf("can not connect your self")
	}
	return nil
}

type FriendListRequest struct {
	Email string `json:"email"`
}

func (email FriendListRequest) Bind(r *http.Request) error {
	if email.Email == "" {
		log.Print("email is a required field")
		return fmt.Errorf("email is a required field")
	}
	return nil
}

type FriendListResponse struct {
	Success bool     `json:"success"`
	Friends []string `json:"friends"`
	Count   int      `json:"count"`
}

type CommonFriendsListRequest struct {
	Friends []string `json:"friends"`
}

func (c CommonFriendsListRequest) Bind(r *http.Request) error {
	emailOne := c.Friends[0]
	emailTwo := c.Friends[1]
	if emailOne == emailTwo {
		return fmt.Errorf("email request are duplicate")
	}
	return nil
}
