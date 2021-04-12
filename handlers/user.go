package handlers

import (
	"encoding/json"
	"github.com/friends-management/helper"
	"github.com/friends-management/models"
	"github.com/friends-management/service"
	"github.com/go-chi/render"
	"log"
	"net/http"
)

type UserHandler struct {
	IUserService service.IUserService
}

func GetUserList(service service.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response, err := service.GetUserList()

		if err != nil {
			_ = render.Render(w, r, ServerErrorRenderer(err))
			return
		}
		json.NewEncoder(w).Encode(response)
	}
}

// https://golang.org/src/net/http/example_test.go
// https://github.com/golang/go/issues/20803
func CreateUser(service service.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &models.UserRequest{}

		if err := render.Bind(r, req); err != nil {
			_ = render.Render(w, r, ErrBadRequest)
			return
		}

		if !helper.IsEmailValid(req.Email) {
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

func CreateFriendConnection(service service.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &models.FriendConnectionRequest{}

		if err := render.Bind(r, req); err != nil {
			render.Render(w, r, ErrBadRequest)
			return
		}

		response, err := service.CreateFriendConnection(req.Friends)

		if err != nil {
			render.Render(w, r, ServerErrorRenderer(err))
			return
		}
		json.NewEncoder(w).Encode(response)
	}
}

func RetrieveFriendList(service service.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &models.FriendListRequest{}

		if err := render.Bind(r, req); err != nil {
			render.Render(w, r, ErrBadRequest)
			return
		}

		if !helper.IsEmailValid(req.Email) {
			log.Println("Email address is invalid")
			_ = render.Render(w, r, ErrBadRequest)
			return
		}

		response, err := service.RetrieveFriendList(req.Email)
		if err != nil {
			render.Render(w, r, ServerErrorRenderer(err))
			return
		}
		json.NewEncoder(w).Encode(response)
	}
}

func GetCommonFriendsList(service service.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &models.CommonFriendsListRequest{}
		if err := render.Bind(r, req); err != nil {
			render.Render(w, r, ErrBadRequest)
			return
		}

		if !helper.IsEmailValid(req.Friends[0]) || !helper.IsEmailValid(req.Friends[1]) {
			log.Println("One Email request is invalid")
			_ = render.Render(w, r, ErrBadRequest)
			return
		}

		response, err := service.GetCommonFriendsList(req.Friends)
		if err != nil {
			render.Render(w, r, ServerErrorRenderer(err))
			return
		}

		json.NewEncoder(w).Encode(response)
	}
}

func CreateSubscribe(service service.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &models.SubscriptionRequest{}
		if err := render.Bind(r, req); err != nil {
			render.Render(w, r, ErrBadRequest)
			return
		}

		if !helper.IsEmailValid(req.Requestor) || !helper.IsEmailValid(req.Target) {
			log.Println("One Email request is invalid")
			_ = render.Render(w, r, ErrBadRequest)
			return
		}

		response, err := service.CreateSubscribe(req)
		if err != nil {
			render.Render(w, r, ServerErrorRenderer(err))
			return
		}

		json.NewEncoder(w).Encode(response)
	}
}

func CreateBlockFriend(service service.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &models.BlockRequest{}
		if err := render.Bind(r, req); err != nil {
			render.Render(w, r, ErrBadRequest)
			return
		}

		if !helper.IsEmailValid(req.Requestor) || !helper.IsEmailValid(req.Target) {
			log.Println("One Email request is invalid")
			_ = render.Render(w, r, ErrBadRequest)
			return
		}

		response, err := service.CreateBlockFriend(req)
		if err != nil {
			render.Render(w, r, ServerErrorRenderer(err))
			return
		}

		json.NewEncoder(w).Encode(response)
	}
}

func ReceiveUpdate(service service.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &models.SendUpdateEmailRequest{}
		if err := render.Bind(r, req); err != nil {
			render.Render(w, r, ErrBadRequest)
			return
		}

		if !helper.IsEmailValid(req.Sender) {
			log.Println("Email request is invalid")
			_ = render.Render(w, r, ErrBadRequest)
			return
		}

		response, err := service.CreateReceiveUpdate(req.Sender)
		if err != nil {
			render.Render(w, r, ServerErrorRenderer(err))
			return
		}

		json.NewEncoder(w).Encode(response)
	}
}
