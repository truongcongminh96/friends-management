package handlers

import (
	"encoding/json"
	"github.com/friends-management/models"
	"github.com/friends-management/service"
	"github.com/go-chi/render"
	"net/http"
)

type UserHandler struct {
	IUserService service.IUserService
}

func (userHandler UserHandler) GetAllUser(w http.ResponseWriter, r *http.Request) {

	response, err := userHandler.IUserService.GetUserList()

	if err != nil {
		_ = render.Render(w, r, ServerErrorRenderer(err))
		return
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

func (userHandler UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	userRequest := models.UserRequest{}
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		_ = render.Render(w, r, ErrBadRequest)
		return
	}

	if err := userRequest.Validate(); err != nil {
		_ = render.Render(w, r, ErrorRenderer(err))
		return
	}

	if err := userHandler.IUserService.CreateUser(userRequest); err != nil {
		_ = render.Render(w, r, ServerErrorRenderer(err))
		return
	}

	err := json.NewEncoder(w).Encode(&models.ResultResponse{
		Success: true,
	})
	if err != nil {
		return
	}
}

//func  CreateUser() http.HandleFunc{
//	return func(w http.ResponseWriter, r *http.Request) {
//		userRequest := models.UserRequest{}
//
//		if err := render.Bind(r, userRequest); err != nil {
//			render.Render(w, r, ErrBadRequest)
//			return
//		}
//
//		err := userHandler.IUserService.CreateUser(userRequest)
//				if err != nil {
//					_ = render.Render(w, r, ServerErrorRenderer(err))
//					return
//				}
//		_ = json.NewEncoder(w).Encode(&models.ResultResponse{
//			Success: true,
//		})
//	})
//}

// https://golang.org/src/net/http/example_test.go
// https://github.com/golang/go/issues/20803
//func (userHandler UserHandler) CreateUser(service service.IUserService) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		req := &models.UserRequest{}
//
//		if err := render.Bind(r, req); err != nil {
//			_ = render.Render(w, r, ErrBadRequest)
//			return
//		}
//
//		if !helper.IsEmailValid(req.Email) {
//			log.Println("Email address is invalid")
//			_ = render.Render(w, r, ErrBadRequest)
//			return
//		}
//
//		err := service.CreateUser(*req)
//		if err != nil {
//			_ = render.Render(w, r, ServerErrorRenderer(err))
//			return
//		}
//
//		_ = json.NewEncoder(w).Encode(&models.ResultResponse{
//			Success: true,
//		})
//	}
//}

//func CreateFriendConnection(service service.IUserService) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		req := &models.FriendConnectionRequest{}
//
//		if err := render.Bind(r, req); err != nil {
//			render.Render(w, r, ErrBadRequest)
//			return
//		}
//
//		response, err := service.CreateFriendConnection(req.Friends)
//
//		if err != nil {
//			render.Render(w, r, ServerErrorRenderer(err))
//			return
//		}
//		json.NewEncoder(w).Encode(response)
//	}
//}
//
//func RetrieveFriendList(service service.IUserService) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		req := &models.FriendListRequest{}
//
//		if err := render.Bind(r, req); err != nil {
//			render.Render(w, r, ErrBadRequest)
//			return
//		}
//
//		if !helper.IsEmailValid(req.Email) {
//			log.Println("Email address is invalid")
//			_ = render.Render(w, r, ErrBadRequest)
//			return
//		}
//
//		response, err := service.RetrieveFriendList(req.Email)
//		if err != nil {
//			render.Render(w, r, ServerErrorRenderer(err))
//			return
//		}
//		json.NewEncoder(w).Encode(response)
//	}
//}
//
//func GetCommonFriendsList(service service.IUserService) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		req := &models.CommonFriendsListRequest{}
//		if err := render.Bind(r, req); err != nil {
//			render.Render(w, r, ErrBadRequest)
//			return
//		}
//
//		if !helper.IsEmailValid(req.Friends[0]) || !helper.IsEmailValid(req.Friends[1]) {
//			log.Println("One Email request is invalid")
//			_ = render.Render(w, r, ErrBadRequest)
//			return
//		}
//
//		response, err := service.GetCommonFriendsList(req.Friends)
//		if err != nil {
//			render.Render(w, r, ServerErrorRenderer(err))
//			return
//		}
//
//		json.NewEncoder(w).Encode(response)
//	}
//}
//
//func CreateSubscribe(service service.IUserService) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		req := &models.SubscriptionRequest{}
//		if err := render.Bind(r, req); err != nil {
//			render.Render(w, r, ErrBadRequest)
//			return
//		}
//
//		if !helper.IsEmailValid(req.Requestor) || !helper.IsEmailValid(req.Target) {
//			log.Println("One Email request is invalid")
//			_ = render.Render(w, r, ErrBadRequest)
//			return
//		}
//
//		response, err := service.CreateSubscribe(req)
//		if err != nil {
//			render.Render(w, r, ServerErrorRenderer(err))
//			return
//		}
//
//		json.NewEncoder(w).Encode(response)
//	}
//}
//
//func CreateBlockFriend(service service.IUserService) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		req := &models.BlockRequest{}
//		if err := render.Bind(r, req); err != nil {
//			render.Render(w, r, ErrBadRequest)
//			return
//		}
//
//		if !helper.IsEmailValid(req.Requestor) || !helper.IsEmailValid(req.Target) {
//			log.Println("One Email request is invalid")
//			_ = render.Render(w, r, ErrBadRequest)
//			return
//		}
//
//		response, err := service.CreateBlockFriend(req)
//		if err != nil {
//			render.Render(w, r, ServerErrorRenderer(err))
//			return
//		}
//
//		json.NewEncoder(w).Encode(response)
//	}
//}
//
//func ReceiveUpdate(service service.IUserService) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		req := &models.SendUpdateEmailRequest{}
//		if err := render.Bind(r, req); err != nil {
//			render.Render(w, r, ErrBadRequest)
//			return
//		}
//
//		if !helper.IsEmailValid(req.Sender) {
//			log.Println("Email request is invalid")
//			_ = render.Render(w, r, ErrBadRequest)
//			return
//		}
//
//		response, err := service.CreateReceiveUpdate(req.Sender)
//		if err != nil {
//			render.Render(w, r, ServerErrorRenderer(err))
//			return
//		}
//
//		json.NewEncoder(w).Encode(response)
//	}
//}
