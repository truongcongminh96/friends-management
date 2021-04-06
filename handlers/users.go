package handlers

import (
	"encoding/json"
	"github.com/friends-management/models"
	"github.com/friends-management/service"
	"github.com/go-chi/render"
	"log"
	"net/http"
)

func getUserList(w http.ResponseWriter, r *http.Request) {
	users, err := dbInstance.GetUserList()
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, users); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}

// https://golang.org/src/net/http/example_test.go
func createUser(service service.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &models.UserRequest{}

		if err := render.Bind(r, req); err != nil {
			render.Render(w, r, ErrBadRequest)
			return
		}

		if !isEmailValid(req.Email) {
			log.Println("Email address is invalid")
			render.Render(w, r, ErrBadRequest)
			return
		}

		response, err := service.CreateUser(req)
		if err != nil {
			render.Render(w, r, ServerErrorRenderer(err))
			return
		}

		json.NewEncoder(w).Encode(response)
	}
}
