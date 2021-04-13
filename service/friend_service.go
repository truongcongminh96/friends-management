package service

import "github.com/friends-management/repositories"

type IFriendService interface {
	CheckExistedFriend(userID1 int, userID2 int) (bool, error)
}

type FriendService struct {
	IFriendRepo repositories.IFriendRepo
	IUserRepo   repositories.IUserRepo
}

func (_friendService FriendService) IsExistedFriend(userID1 int, userID2 int) (bool, error) {
	isExist, err := _friendService.IFriendRepo.IsExistedFriend(userID1, userID2)
	return isExist, err
}