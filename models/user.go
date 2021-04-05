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

func (u User) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (u User) Bind(r *http.Request) error {
	if u.Email == "" {
		log.Print("email is a required field")
		return fmt.Errorf("email is a required field")
	}
	return nil
}

type UserList struct {
	Users []User
}

func (u UserList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
