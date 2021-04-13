package service

import (
	"github.com/friends-management/models"
	"github.com/friends-management/repositories"
)

type IFriendService interface {
	CreateFriend(friend *models.Friend) error
	CheckExistedFriend(userID1 int, userID2 int) (bool, error)
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
