package service

import (
	"github.com/friends-management/models"
	"github.com/friends-management/repositories"
)

type IFriendService interface {
	CreateFriend(friend *models.Friend) error
	CheckExistedFriend(userID1 int, userID2 int) (bool, error)
	CheckBlockedByUser(requestorId int, targetId int) (bool, string, error)
	GetFriendsList(userId int) ([]string, error)
	GetCommonFriends(Ids []int) ([]string, error)
}

type FriendService struct {
	IFriendRepo repositories.IFriendRepo
	IUserRepo   repositories.IUserRepo
}

func (_friendService FriendService) CreateFriend(friend *models.Friend) error {
	err := _friendService.IFriendRepo.CreateFriend(friend)
	return err
}

func (_friendService FriendService) CheckExistedFriend(userID1 int, userID2 int) (bool, error) {
	isExist, err := _friendService.IFriendRepo.IsExistedFriend(userID1, userID2)
	return isExist, err
}

func (_friendService FriendService) GetFriendsList(userId int) ([]string, error) {
	listFriendId, err := _friendService.IFriendRepo.GetListFriendId(userId)
	if err != nil {
		return nil, err
	}

	result, err := _friendService.IUserRepo.GetEmailsByIDs(listFriendId)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (_friendService FriendService) CheckBlockedByUser(requestorId int, targetId int) (bool, string, error) {
	isBlocked, message, err := _friendService.IFriendRepo.IsBlockedByUser(requestorId, targetId)
	return isBlocked, message, err
}

func (_friendService FriendService) GetCommonFriends(Ids []int) ([]string, error) {
	// Get friends list of each email address
	friends1, err := _friendService.GetFriendsList(Ids[0])
	if err != nil {
		return nil, err
	}
	friends2, err := _friendService.GetFriendsList(Ids[1])
	if err != nil {
		return nil, err
	}

	// Find common friends in 2 friends list
	var commonFriends []string
	hashMap := make(map[string]bool)
	for _, email := range friends1 {
		hashMap[email] = true

	}
	for _, email := range friends2 {
		if _, ok := hashMap[email]; ok {
			commonFriends = append(commonFriends, email)
		}
	}
	return commonFriends, nil
}
