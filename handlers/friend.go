package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/friends-management/models"
	"github.com/friends-management/service"
	"github.com/go-chi/render"
)

type FriendHandlers struct {
	IFriendService service.IFriendService
	IUserService   service.IUserService
}

func (_friendHandlers FriendHandlers) CreateFriend(w http.ResponseWriter, r *http.Request) {
	// Decode request body
	friendRequest := models.FriendRequest{}
	if err := json.NewDecoder(r.Body).Decode(&friendRequest); err != nil {
		_ = render.Render(w, r, ErrBadRequest)
		return
	}

	// Validate, check email valid request body
	if err := friendRequest.Validate(); err != nil {
		_ = render.Render(w, r, ErrorRenderer(err, 400))
		return
	}

	// Check user and friend's
	Ids, statusCode, err := _friendHandlers.checkFriendRelationship(friendRequest.Friends)
	if err != nil {
		_ = render.Render(w, r, ErrorRenderer(err, statusCode))
		return
	}

	friendModel := &models.Friend{
		User1: Ids[0],
		User2: Ids[1],
	}

	// Call service to create friend
	if err := _friendHandlers.IFriendService.CreateFriend(friendModel); err != nil {
		_ = render.Render(w, r, ServerErrorRenderer(err))
		return
	}

	err = json.NewEncoder(w).Encode(&models.SuccessResponse{
		Success: true,
	})
	if err != nil {
		return
	}
}

func (_friendHandlers FriendHandlers) checkFriendRelationship(friendRequest []string) ([]int, int, error) {
	// Check first email exists
	userId1, err := _friendHandlers.IUserService.GetUserIDByEmail(friendRequest[0])
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if userId1 == 0 {
		return nil, http.StatusBadRequest, errors.New("your email does not exist")
	}

	// Check second email exists
	userId2, err := _friendHandlers.IUserService.GetUserIDByEmail(friendRequest[1])
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if userId2 == 0 {
		return nil, http.StatusBadRequest, errors.New("your friend email does not exist")
	}

	// Check friend connection exists
	isExists, err := _friendHandlers.IFriendService.CheckExistedFriend(userId1, userId2)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if isExists {
		return nil, http.StatusAlreadyReported, errors.New("you are friends")
	}

	// Check blocking
	isBlocked, message, err := _friendHandlers.IFriendService.CheckBlockedByUser(userId1, userId2)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if isBlocked {
		return nil, http.StatusPreconditionFailed, errors.New(message)
	}

	return []int{userId1, userId2}, 0, nil
}

func (_friendHandlers *FriendHandlers) GetFriendsList(w http.ResponseWriter, r *http.Request) {
	var emailRequest models.FriendsListRequest
	if err := json.NewDecoder(r.Body).Decode(&emailRequest); err != nil {
		_ = render.Render(w, r, ErrBadRequest)
		return
	}

	// Validate email request
	if err := emailRequest.Validate(); err != nil {
		_ = render.Render(w, r, ErrorRenderer(err, 400))
		return
	}

	// Check user exists and get userID
	userId, statusCode, err := _friendHandlers.getFriendsListValidation(emailRequest.Email)
	if err != nil {
		_ = render.Render(w, r, ErrorRenderer(err, statusCode))
		return
	}

	// Call service to get friends list
	friendsList, err := _friendHandlers.IFriendService.GetFriendsList(userId)
	if err != nil {
		_ = render.Render(w, r, ServerErrorRenderer(err))
		return
	}

	// Response
	err = json.NewEncoder(w).Encode(models.FriendsResponse{
		Success: true,
		Friends: friendsList,
		Count:   len(friendsList),
	})
	if err != nil {
		return
	}
	return
}

func (_friendHandlers *FriendHandlers) GetCommonFriends(w http.ResponseWriter, r *http.Request) {
	// Decode request body
	var emailRequest models.CommonFriendsRequest
	if err := json.NewDecoder(r.Body).Decode(&emailRequest); err != nil {
		_ = render.Render(w, r, ErrBadRequest)
		return
	}

	// Validate request body
	if err := emailRequest.Validate(); err != nil {
		_ = render.Render(w, r, ErrorRenderer(err, 400))
		return
	}

	// Check users exists and get userIDs
	userIds, statusCode, err := _friendHandlers.checkCommonFriends(emailRequest.Friends)
	if err != nil {
		_ = render.Render(w, r, ErrorRenderer(err, statusCode))
		return
	}

	// Call service to get common friends
	friendsList, err := _friendHandlers.IFriendService.GetCommonFriends(userIds)
	if err != nil {
		_ = render.Render(w, r, ServerErrorRenderer(err))
		return
	}

	// Response
	err = json.NewEncoder(w).Encode(models.FriendsResponse{
		Success: true,
		Friends: friendsList,
		Count:   len(friendsList),
	})
	if err != nil {
		return
	}

	return
}

func (_friendHandlers *FriendHandlers) getFriendsListValidation(email string) (int, int, error) {
	userID, err := _friendHandlers.IUserService.GetUserIDByEmail(email)
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}
	if userID == 0 {
		return 0, http.StatusBadRequest, errors.New("email does not exist")
	}
	return userID, 0, nil
}

func (_friendHandlers *FriendHandlers) checkCommonFriends(email []string) ([]int, int, error) {
	// Check first email
	userId1, err := _friendHandlers.IUserService.GetUserIDByEmail(email[0])
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if userId1 == 0 {
		return nil, http.StatusBadRequest, errors.New("the first email does not exist")
	}

	// Check second email
	userId2, err := _friendHandlers.IUserService.GetUserIDByEmail(email[1])
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if userId2 == 0 {
		return nil, http.StatusBadRequest, errors.New("the second email does not exist")
	}

	return []int{userId1, userId2}, 0, nil
}

func (_friendHandlers *FriendHandlers) ReceiveUpdate(w http.ResponseWriter, r *http.Request) {
	// Decode request body
	var emailRequest models.ReceiveUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&emailRequest); err != nil {
		_ = render.Render(w, r, ErrBadRequest)
		return
	}

	// Validate request
	if err := emailRequest.Validate(); err != nil {
		_ = render.Render(w, r, ErrBadRequest)
		return
	}

	// Check user exists and get userID
	senderId, statusCode, err := _friendHandlers.checkEmailReceiveUpdate(emailRequest.Sender)
	if err != nil {
		_ = render.Render(w, r, ErrorRenderer(err, statusCode))
		return
	}

	// Call service to get emails receiving updates
	recipients, err := _friendHandlers.IFriendService.GetEmailsReceiveUpdate(senderId, emailRequest.Text)
	if err != nil {
		_ = render.Render(w, r, ServerErrorRenderer(err))
		return
	}

	// Response
	err = json.NewEncoder(w).Encode(models.ReceiveUpdateResponse{
		Success:    true,
		Recipients: recipients,
	})
	if err != nil {
		return
	}
	return
}

func (_friendHandlers *FriendHandlers) checkEmailReceiveUpdate(email string) (int, int, error) {
	userId, err := _friendHandlers.IUserService.GetUserIDByEmail(email)
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}
	if userId == 0 {
		return 0, http.StatusBadRequest, errors.New("the sender does not exist")
	}
	return userId, 0, nil
}
