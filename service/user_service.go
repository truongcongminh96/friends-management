package service

import (
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
