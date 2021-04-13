package handlers

import (
	"encoding/json"
	"errors"
	"github.com/friends-management/models"
	"github.com/friends-management/service"
	"github.com/go-chi/render"
	"net/http"
)

type SubscribeHandlers struct {
	ISubscribeService service.ISubscribeService
	IUserService      service.IUserService
}

func (_subscribeHandlers *SubscribeHandlers) CreateSubscribe(w http.ResponseWriter, r *http.Request) {
	// Decode request body
	var subscribeRequest models.SubscribeRequest
	if err := json.NewDecoder(r.Body).Decode(&subscribeRequest); err != nil {
		_ = render.Render(w, r, ErrBadRequest)
		return
	}

	// Validate request
	if err := subscribeRequest.Validate(); err != nil {
		_ = render.Render(w, r, ErrBadRequest)
		return
	}

	// Check subscribe
	Ids, statusCode, err := _subscribeHandlers.checkSubscribeEmail(subscribeRequest)
	if err != nil {
		_ = render.Render(w, r, ErrorRenderer(err, statusCode))
		return
	}

	subscribeModel := &models.Subscribe{
		Requestor: Ids[0],
		Target:    Ids[1],
	}

	// Call service to create subscribe
	if err := _subscribeHandlers.ISubscribeService.CreateSubscribe(subscribeModel); err != nil {
		_ = render.Render(w, r, ServerErrorRenderer(err))
		return
	}

	// Response
	err = json.NewEncoder(w).Encode(models.SuccessResponse{
		Success: true,
	})
	if err != nil {
		return
	}
	return
}

func (_subscribeHandlers *SubscribeHandlers) checkSubscribeEmail(subscribe models.SubscribeRequest) ([]int, int, error) {
	// Check user id of the requestor
	requestorId, err := _subscribeHandlers.IUserService.GetUserIDByEmail(subscribe.Requestor)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if requestorId == 0 {
		return nil, http.StatusBadRequest, errors.New("email requestor does not exist")
	}

	// Check user id of the target
	targetId, err := _subscribeHandlers.IUserService.GetUserIDByEmail(subscribe.Target)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if targetId == 0 {
		return nil, http.StatusBadRequest, errors.New("email target does not exist")
	}

	// Check exists subscribe
	isExist, err := _subscribeHandlers.ISubscribeService.CheckExistedSubscribe(requestorId, targetId)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if isExist {
		return nil, http.StatusAlreadyReported, errors.New("you are subscribed the target")
	}

	return []int{requestorId, targetId}, 0, nil
}
