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

	// Routes for users
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

	// Routes for friends
	router.Route("/friend", func(r chi.Router) {
		friendHandlers := handlers.FriendHandlers{
			IFriendService: service.FriendService{
				IFriendRepo: repositories.FriendRepo{
					Db: db.Conn,
				},
				IUserRepo: repositories.UserRepo{
					Db: db.Conn,
				},
			},
			IUserService: service.UserService{
				IUserRepo: repositories.UserRepo{
					Db: db.Conn,
				},
			},
		}
		r.MethodFunc(http.MethodPost, "/", friendHandlers.CreateFriend)
		r.MethodFunc(http.MethodPost, "/friends-list", friendHandlers.GetFriendsList)
		r.MethodFunc(http.MethodPost, "/common-friends", friendHandlers.GetCommonFriends)
		r.MethodFunc(http.MethodPost, "/receive-updates", friendHandlers.ReceiveUpdate)
	})

	// Routes for subscribe
	router.Route("/subscribe", func(r chi.Router) {
		SubscribeHandlers := handlers.SubscribeHandlers{
			ISubscribeService: service.SubscribeService{
				ISubscribeRepo: repositories.SubscribeRepo{
					Db: db.Conn,
				},
			},
			IUserService: service.UserService{
				IUserRepo: repositories.UserRepo{
					Db: db.Conn,
				},
			},
		}
		r.MethodFunc(http.MethodPost, "/", SubscribeHandlers.CreateSubscribe)
	})

	// Routes for block
	router.Route("/block", func(r chi.Router) {
		blockHandlers := handlers.BlockHandlers{
			IBlockService: service.BlockService{
				IBlockRepo: repositories.BlockRepo{
					Db: db.Conn,
				},
			},
			IUserService: service.UserService{
				IUserRepo: repositories.UserRepo{
					Db: db.Conn,
				},
			},
		}
		r.MethodFunc(http.MethodPost, "/", blockHandlers.CreateBlock)
	})
}
