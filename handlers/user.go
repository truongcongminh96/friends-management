package handlers

import (
	"encoding/json"
	"errors"
	"github.com/friends-management/models"
	"github.com/friends-management/service"
	"github.com/go-chi/render"
	"net/http"
)

type UserHandler struct {
	IUserService service.IUserService
}

func (_userHandler UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Decode request body
	userRequest := models.UserRequest{}
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		_ = render.Render(w, r, ErrBadRequest)
		return
	}

	// Validate, check email valid request body
	if err := userRequest.Validate(); err != nil {
		_ = render.Render(w, r, ErrorRenderer(err, 400))
		return
	}

	// Check user exist
	if statusCode, err := _userHandler.checkExistedUser(userRequest.Email); err != nil {
		_ = render.Render(w, r, ErrorRenderer(err, statusCode))
		return
	}

	userModel := &models.User{
		Email: userRequest.Email,
	}

	// Call service to create user
	if err := _userHandler.IUserService.CreateUser(userModel); err != nil {
		_ = render.Render(w, r, ServerErrorRenderer(err))
		return
	}

	err := json.NewEncoder(w).Encode(&models.SuccessResponse{
		Success: true,
	})
	if err != nil {
		return
	}
}

func (_userHandler *UserHandler) checkExistedUser(email string) (int, error) {
	isExist, err := _userHandler.IUserService.IsExistedUser(email)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if isExist {
		return http.StatusAlreadyReported, errors.New("email address exists")
	}
	return 0, nil
}
