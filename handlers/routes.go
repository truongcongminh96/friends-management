package handlers

import (
	"github.com/friends-management/database"
	"github.com/friends-management/service"
	"net/http"

	_ "github.com/friends-management/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

var dbInstance database.Database

func NewHandler(db database.Database) http.Handler {
	router := chi.NewRouter()
	dbInstance = db
	router.MethodNotAllowed(methodNotAllowedHandler)
	router.NotFound(notFoundHandler)
	router.Route("/api/v1/", users)
	return router
}

func users(router chi.Router) {
	db := service.DbInstance{Db: dbInstance}
	router.Get("/user_list", getUserList(db))
	router.Post("/registration", createUser(db))
	router.Post("/friendConnection", createFriendConnection(db))
	router.Post("/retrieveFriendList", retrieveFriendList(db))
	router.Post("/commonFriends", getCommonFriendsList(db))
}

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(405)
	render.Render(w, r, ErrMethodNotAllowed)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(400)
	render.Render(w, r, ErrNotFound)
}
