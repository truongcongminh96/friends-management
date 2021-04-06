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
		_ = render.Render(w, r, ServerErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, users); err != nil {
		_ = render.Render(w, r, ErrorRenderer(err))
	}
}

// https://golang.org/src/net/http/example_test.go
// https://github.com/golang/go/issues/20803
func createUser(service service.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &models.UserRequest{}

		if err := render.Bind(r, req); err != nil {
			_ = render.Render(w, r, ErrBadRequest)
			return
		}

		if !isEmailValid(req.Email) {
			log.Println("Email address is invalid")
			_ = render.Render(w, r, ErrBadRequest)
			return
		}

		response, err := service.CreateUser(req.Email)
		if err != nil {
			_ = render.Render(w, r, ServerErrorRenderer(err))
			return
		}

		_ = json.NewEncoder(w).Encode(response)
	}
}
