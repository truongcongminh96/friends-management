package service

import (
	"github.com/friends-management/database"
	"github.com/friends-management/models"
	"github.com/friends-management/repositories"
)

type DbInstance struct {
	Db database.Database
}

type IUserService interface {
	CreateUser(user *models.User) error
	IsExistedUser(email string) (bool, error)
}

type UserService struct {
	IUserRepo repositories.IUserRepo
}

func (_userService UserService) CreateUser(user *models.User) error {
	err := _userService.IUserRepo.CreateUser(user)
	return err
}

func (_userService UserService) IsExistedUser(email string) (bool, error) {
	exist, err := _userService.IUserRepo.IsExistedUser(email)
	return exist, err
}
