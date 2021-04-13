package routes

import (
	"database/sql"
	"github.com/friends-management/database"
	"github.com/friends-management/repositories"
	"github.com/friends-management/service"
	"net/http"

	_ "github.com/friends-management/database"
	"github.com/friends-management/handlers"
	"github.com/go-chi/chi"
)

type Database struct {
	Conn *sql.DB
}

var dbInstance Database

func NewHandler(db database.Database) http.Handler {
	router := chi.NewRouter()
	dbInstance = Database(db)
	router.MethodNotAllowed(handlers.MethodNotAllowedHandler)
	router.NotFound(handlers.NotFoundHandler)
	router.Route("/api/v1/", createRoutes)
	return router
}

func createRoutes(router chi.Router) {
	db := dbInstance

	router.Route("/user", func(r chi.Router) {
		UserHandler := handlers.UserHandler{
			IUserService: service.UserService{
				IUserRepo: repositories.UserRepo{
					Db: db.Conn,
				},
			},
		}
		r.MethodFunc(http.MethodPost, "/", UserHandler.CreateUser)
	})
}
