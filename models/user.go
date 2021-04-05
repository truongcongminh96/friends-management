package models

import "net/http"

type User struct {
	Email        string   `json:"email"`
	Friends      []string `json:"friends"`
	Subscription []string `json:"subscription"`
	Blocked      []string `json:"blocked"`
}

type UserList struct {
	Users []User
}

func (u UserList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
