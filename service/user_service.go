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
}

func (db DbInstance) CreateUser(email string) (*models.ResultResponse, error) {
	response := &models.ResultResponse{}
	if err := db.Db.CreateUser(email); err != nil {
		return response, err
	}
	response.Success = true
	return response, nil
}
