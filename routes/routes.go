package routes

import (
	"encoding/json"
	"github.com/friends-management/database"
	"github.com/friends-management/helper"
	"github.com/friends-management/models"
	"github.com/friends-management/service"
	"log"
	"net/http"

	_ "github.com/friends-management/database"
	"github.com/friends-management/handlers"
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
	router.Post("/subscribe", createSubscribe(db))
	router.Post("/blockFriend", createBlockFriend(db))
	router.Post("/receiveUpdates", receiveUpdate(db))
}

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(405)
	render.Render(w, r, handlers.ErrMethodNotAllowed)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(400)
	render.Render(w, r, handlers.ErrNotFound)
}

func getUserList(service service.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response, err := service.GetUserList()

		if err != nil {
			_ = render.Render(w, r, handlers.ServerErrorRenderer(err))
			return
		}
		json.NewEncoder(w).Encode(response)
	}
}

// https://golang.org/src/net/http/example_test.go
// https://github.com/golang/go/issues/20803
func createUser(service service.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &models.UserRequest{}

		if err := render.Bind(r, req); err != nil {
			_ = render.Render(w, r, handlers.ErrBadRequest)
			return
		}

		if !helper.IsEmailValid(req.Email) {
			log.Println("Email address is invalid")
			_ = render.Render(w, r, handlers.ErrBadRequest)
			return
		}

		response, err := service.CreateUser(req.Email)
		if err != nil {
			_ = render.Render(w, r, handlers.ServerErrorRenderer(err))
			return
		}

		_ = json.NewEncoder(w).Encode(response)
	}
}

func createFriendConnection(service service.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &models.FriendConnectionRequest{}

		if err := render.Bind(r, req); err != nil {
			render.Render(w, r, handlers.ErrBadRequest)
			return
		}

		response, err := service.CreateFriendConnection(req.Friends)

		if err != nil {
			render.Render(w, r, handlers.ServerErrorRenderer(err))
			return
		}
		json.NewEncoder(w).Encode(response)
	}
}

func retrieveFriendList(service service.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &models.FriendListRequest{}

		if err := render.Bind(r, req); err != nil {
			render.Render(w, r, handlers.ErrBadRequest)
			return
		}

		if !helper.IsEmailValid(req.Email) {
			log.Println("Email address is invalid")
			_ = render.Render(w, r, handlers.ErrBadRequest)
			return
		}

		response, err := service.RetrieveFriendList(req.Email)
		if err != nil {
			render.Render(w, r, handlers.ServerErrorRenderer(err))
			return
		}
		json.NewEncoder(w).Encode(response)
	}
}

func getCommonFriendsList(service service.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &models.CommonFriendsListRequest{}
		if err := render.Bind(r, req); err != nil {
			render.Render(w, r, handlers.ErrBadRequest)
			return
		}

		if !helper.IsEmailValid(req.Friends[0]) || !helper.IsEmailValid(req.Friends[1]) {
			log.Println("One Email request is invalid")
			_ = render.Render(w, r, handlers.ErrBadRequest)
			return
		}

		response, err := service.GetCommonFriendsList(req.Friends)
		if err != nil {
			render.Render(w, r, handlers.ServerErrorRenderer(err))
			return
		}

		json.NewEncoder(w).Encode(response)
	}
}

func createSubscribe(service service.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &models.SubscriptionRequest{}
		if err := render.Bind(r, req); err != nil {
			render.Render(w, r, handlers.ErrBadRequest)
			return
		}

		if !helper.IsEmailValid(req.Requestor) || !helper.IsEmailValid(req.Target) {
			log.Println("One Email request is invalid")
			_ = render.Render(w, r, handlers.ErrBadRequest)
			return
		}

		response, err := service.CreateSubscribe(req)
		if err != nil {
			render.Render(w, r, handlers.ServerErrorRenderer(err))
			return
		}

		json.NewEncoder(w).Encode(response)
	}
}

func createBlockFriend(service service.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &models.BlockRequest{}
		if err := render.Bind(r, req); err != nil {
			render.Render(w, r, handlers.ErrBadRequest)
			return
		}

		if !helper.IsEmailValid(req.Requestor) || !helper.IsEmailValid(req.Target) {
			log.Println("One Email request is invalid")
			_ = render.Render(w, r, handlers.ErrBadRequest)
			return
		}

		response, err := service.CreateBlockFriend(req)
		if err != nil {
			render.Render(w, r, handlers.ServerErrorRenderer(err))
			return
		}

		json.NewEncoder(w).Encode(response)
	}
}

func receiveUpdate(service service.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &models.SendUpdateEmailRequest{}
		if err := render.Bind(r, req); err != nil {
			render.Render(w, r, handlers.ErrBadRequest)
			return
		}

		if !helper.IsEmailValid(req.Sender) {
			log.Println("Email request is invalid")
			_ = render.Render(w, r, handlers.ErrBadRequest)
			return
		}

		response, err := service.CreateReceiveUpdate(req.Sender)
		if err != nil {
			render.Render(w, r, handlers.ServerErrorRenderer(err))
			return
		}

		json.NewEncoder(w).Encode(response)
	}
}
