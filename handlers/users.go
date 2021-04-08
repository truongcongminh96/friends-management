package handlers

import (
	"encoding/json"
	"github.com/friends-management/models"
	"github.com/friends-management/service"
	"github.com/go-chi/render"
	"log"
	"net/http"
)

func getUserList(service service.IUserService) http.HandlerFunc {
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

func createFriendConnection(service service.IUserService) http.HandlerFunc {
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

func retrieveFriendList(service service.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &models.FriendListRequest{}

		if err := render.Bind(r, req); err != nil {
			render.Render(w, r, ErrBadRequest)
			return
		}

		if !isEmailValid(req.Email) {
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

func getCommonFriendsList(service service.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &models.CommonFriendsListRequest{}
		if err := render.Bind(r, req); err != nil {
			render.Render(w, r, ErrBadRequest)
			return
		}

		if !isEmailValid(req.Friends[0]) || !isEmailValid(req.Friends[1]) {
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

func createSubscribe(service service.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &models.SubscriptionRequest{}
		if err := render.Bind(r, req); err != nil {
			render.Render(w, r, ErrBadRequest)
			return
		}

		if !isEmailValid(req.Requestor) || !isEmailValid(req.Target) {
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

func createBlockFriend(service service.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &models.BlockRequest{}
		if err := render.Bind(r, req); err != nil {
			render.Render(w, r, ErrBadRequest)
			return
		}

		if !isEmailValid(req.Requestor) || !isEmailValid(req.Target) {
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

func receiveUpdate(service service.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &models.SendUpdateEmailRequest{}
		if err := render.Bind(r, req); err != nil {
			render.Render(w, r, ErrBadRequest)
			return
		}

		if !isEmailValid(req.Sender) {
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
