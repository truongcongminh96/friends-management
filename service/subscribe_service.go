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
}

func (_subscribeService SubscribeService) CheckExistedSubscribe(requestorId int, targetId int) (bool, error) {
	exist, err := _subscribeService.ISubscribeRepo.IsExistedSubscribe(requestorId, targetId)
	return exist, err
}

func (_subscribeService SubscribeService) CreateSubscribe(subscribe *models.Subscribe) error {
	err := _subscribeService.ISubscribeRepo.CreateSubscribe(subscribe)
	return err
}
