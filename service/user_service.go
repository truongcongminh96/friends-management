package service

import (
	"fmt"
	"github.com/friends-management/database"
	"github.com/friends-management/models"
)

type DbInstance struct {
	Db database.Database
}

type IUserService interface {
	GetUserList() (*models.UserListResponse, error)
	CreateUser(email string) (*models.ResultResponse, error)
	CreateFriendConnection(req []string) (*models.ResultResponse, error)
	RetrieveFriendList(email string) (*models.FriendListResponse, error)
	GetCommonFriendsList(req []string) (*models.FriendListResponse, error)
	CreateSubscribe(req *models.SubscriptionRequest) (*models.ResultResponse, error)
	CreateBlockFriend(req *models.BlockRequest) (*models.ResultResponse, error)
	CreateReceiveUpdate(req *models.SendUpdateEmailRequest) (*models.SendUpdateEmailResponse, error)
}

func (db DbInstance) GetUserList() (*models.UserListResponse, error) {
	response := &models.UserListResponse{}

	rep, err := db.Db.GetUserList()

	if err != nil {
		return response, err
	}

	response.Users = rep.Users
	return response, nil
}

func (db DbInstance) CreateUser(email string) (*models.ResultResponse, error) {
	response := &models.ResultResponse{}
	if err := db.Db.CreateUser(email); err != nil {
		return response, err
	}
	response.Success = true
	return response, nil
}

func (db DbInstance) CreateFriendConnection(req []string) (*models.ResultResponse, error) {
	response := &models.ResultResponse{}
	if err := db.Db.CreateFriendConnection(req); err != nil {
		response.Success = false
		return response, err
	}
	response.Success = true
	return response, nil
}

func (db DbInstance) RetrieveFriendList(email string) (*models.FriendListResponse, error) {
	response := &models.FriendListResponse{}
	data, err := db.Db.RetrieveFriendListByEmail(email)
	if err != nil {
		return response, err
	}
	response.Success = true
	response.Friends = data.Friends
	response.Count = len(data.Friends)
	return response, nil
}

func (db DbInstance) GetCommonFriendsList(email []string) (*models.FriendListResponse, error) {
	response := &models.FriendListResponse{}

	friendListEmailOne, err := db.Db.RetrieveFriendListByEmail(email[0])
	if err != nil {
		return response, err
	}

	friendListEmailTwo, err := db.Db.RetrieveFriendListByEmail(email[1])
	if err != nil {
		return response, err
	}

	var mutualFriends []string

	// mutualFriends = append(friendListEmailOne.Friends, friendListEmailTwo.Friends...)
	// need to improve
	for _, friendOne := range friendListEmailOne.Friends {
		for _, friendTwo := range friendListEmailTwo.Friends {
			if friendOne == friendTwo {
				mutualFriends = append(mutualFriends, friendOne)
			}
		}
	}

	response.Success = true
	response.Friends = mutualFriends
	response.Count = len(mutualFriends)

	return response, nil
}

func (db DbInstance) CreateSubscribe(req *models.SubscriptionRequest) (*models.ResultResponse, error) {
	response := &models.ResultResponse{}

	if err := db.Db.CreateSubscribe(req.Requestor, req.Target); err != nil {
		return response, err
	}
	response.Success = true
	return response, nil
}

func (db DbInstance) CreateBlockFriend(req *models.BlockRequest) (*models.ResultResponse, error) {
	response := &models.ResultResponse{}

	if err := db.Db.CreateBlockFriend(req.Requestor, req.Target); err != nil {
		return response, err
	}

	response.Success = true
	return response, nil
}

func (db DbInstance) CreateReceiveUpdate(req *models.SendUpdateEmailRequest) (*models.SendUpdateEmailResponse, error) {
	response := &models.SendUpdateEmailResponse{}

	listUsers, err := db.Db.GetUserList()

	blockedList, err := db.Db.GetAllBlockerByEmail(req.Sender)
	if err != nil {
		return response, err
	}

	friendList, err := db.Db.RetrieveFriendListByEmail(req.Sender)
	if err != nil {
		return response, err
	}

	subscriber, err := db.Db.GetAllSubscriber(req.Sender)
	if err != nil {
		return response, err
	}

	var usersNotBlocked []string
	for _, user := range listUsers.Users {
		for _, userBlock := range blockedList.Blocked {
			if user.Email != userBlock {
				usersNotBlocked = append(usersNotBlocked, user.Email)
			}
		}
	}

	var listFriendsNotBlock []string
	for _, userNotBlocked := range usersNotBlocked {
		for _, friends := range friendList.Friends {
			if userNotBlocked == friends {
				listFriendsNotBlock = append(listFriendsNotBlock, userNotBlocked)
			}
		}
	}

	var recipient []string
	recipient = append(recipient, listFriendsNotBlock...)
	recipient = append(recipient, subscriber.Subscription...)

	fmt.Println(recipient)
	response.Success = true
	response.Recipients = recipient
	return response, nil
}
