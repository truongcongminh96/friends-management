package handlers

import (
	"encoding/json"
	"github.com/friends-management/models"
	"github.com/friends-management/service"
	"net/http"
)

type FriendHandlers struct {
	IFriendService service.IFriendService
	IUserService   service.IUserService
}

func (_friendHandlers FriendHandlers) CreateFriend(w http.ResponseWriter, r *http.Request)  {
	err := json.NewEncoder(w).Encode(&models.SuccessResponse{
		Success: true,
	})
	if err != nil {
		return
	}
}