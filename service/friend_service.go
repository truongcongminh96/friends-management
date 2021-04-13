package service

import "github.com/friends-management/repositories"

type IFriendService interface {

}

type FriendService struct {
	IFriendRepo repositories.IFriendRepo
	IUserRepo repositories.IUserRepo
}
