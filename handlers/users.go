package handlers

import (
	"github.com/friends-management/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

func users(router chi.Router) {
	router.Get("/", getHomeRoot)
	router.Get("/user_list", getUserList)
	router.Post("/registration", createUser)
}

func getHomeRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi"))
}

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

func createUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	if err := render.Bind(r, user); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	if err := dbInstance.CreateUser(user.Email); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}

	if err := render.Render(w, r, user); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}
