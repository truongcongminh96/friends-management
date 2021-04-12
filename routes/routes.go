package routes

import (
	"github.com/friends-management/database"
	"github.com/friends-management/service"
	"net/http"

	_ "github.com/friends-management/database"
	"github.com/friends-management/handlers"
	"github.com/go-chi/chi"
)

var dbInstance database.Database

func NewHandler(db database.Database) http.Handler {
	router := chi.NewRouter()
	dbInstance = db
	router.MethodNotAllowed(handlers.MethodNotAllowedHandler)
	router.NotFound(handlers.NotFoundHandler)
	router.Route("/api/v1/", createRoutes)
	return router
}

func createRoutes(router chi.Router) {
	db := service.DbInstance{Db: dbInstance}

	router.Route("/user", func(r chi.Router) {
		r.MethodFunc(http.MethodGet, "/list", handlers.GetUserList(db))
		r.MethodFunc(http.MethodPost, "/", handlers.CreateUser(db))
	})

	router.Route("/friend", func(r chi.Router) {
		r.MethodFunc(http.MethodPost, "/", handlers.CreateFriendConnection(db))
		r.MethodFunc(http.MethodPost, "/retrieveFriendList", handlers.RetrieveFriendList(db))
		r.MethodFunc(http.MethodPost, "/commonFriends", handlers.GetCommonFriendsList(db))
		r.MethodFunc(http.MethodPost, "/blockFriend", handlers.CreateBlockFriend(db))
		r.MethodFunc(http.MethodPost, "/receiveUpdates", handlers.ReceiveUpdate(db))
	})

	router.Route("/subscribe", func(r chi.Router) {
		r.MethodFunc(http.MethodPost, "/", handlers.CreateSubscribe(db))
	})
}
