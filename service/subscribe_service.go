package service

import (
	"github.com/friends-management/models"
	"github.com/friends-management/repositories"
)

type SubscribeService struct {
	ISubscribeRepo repositories.ISubscribeRepo
}

type ISubscribeService interface {
	CreateSubscribe(subscribe *models.Subscribe) error
	CheckExistedSubscribe(requestorId int, targetId int) (bool, error)
	CheckBlockedByUser(requestorId int, targetId int) (bool, string, error)
}

func (_subscribeService SubscribeService) CheckExistedSubscribe(requestorId int, targetId int) (bool, error) {
	isExist, err := _subscribeService.ISubscribeRepo.IsExistedSubscribe(requestorId, targetId)
	return isExist, err
}

func (_subscribeService SubscribeService) CreateSubscribe(subscribe *models.Subscribe) error {
	err := _subscribeService.ISubscribeRepo.CreateSubscribe(subscribe)
	return err
}

func (_subscribeService SubscribeService) CheckBlockedByUser(requestorId int, targetId int) (bool, string, error) {
	isBlocked, message, err := _subscribeService.ISubscribeRepo.IsBlockedByUser(requestorId, targetId)
	return isBlocked, message, err
}
