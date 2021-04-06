package service

import (
	"github.com/friends-management/database"
	"github.com/friends-management/models"
)

type DbInstance struct {
	Db database.Database
}

type IUserService interface {
	CreateUser(email string) (*models.ResultResponse, error)
	CreateFriendConnection(req []string) (*models.ResultResponse, error)
	RetrieveFriendList(email string) (*models.FriendListResponse, error)
	GetUserList() (*models.UserListResponse, error)
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
		if err != nil {
			response.Success = false
			return response, err
		}
	}
	response.Success = true
	return response, nil
}

func (db DbInstance) RetrieveFriendList(email string) (*models.FriendListResponse, error) {
	response := &models.FriendListResponse{}
	rep, err := db.Db.RetrieveFriendListByEmail(email)
	if err != nil {
		return response, err
	}
	response.Success = true
	response.Friends = rep.Friends
	response.Count = len(rep.Friends)
	return response, nil
}
