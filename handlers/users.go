package handlers

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

func users(router chi.Router) {
	router.Get("/", getHomeRoot)
	router.Get("/user_list", getUserList)

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
