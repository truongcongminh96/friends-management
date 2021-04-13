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
	GetUserIDByEmail(email string) (int, error)
}

type UserService struct {
	IUserRepo repositories.IUserRepo
}

func (_userService UserService) CreateUser(user *models.User) error {
	err := _userService.IUserRepo.CreateUser(user)
	return err
}

func (_userService UserService) IsExistedUser(email string) (bool, error) {
	isExist, err := _userService.IUserRepo.IsExistedUser(email)
	return isExist, err
}

func (_userService UserService) GetUserIDByEmail(email string) (int, error) {
	userId, err := _userService.IUserRepo.GetUserIDByEmail(email)
	return userId, err
}
