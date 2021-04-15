package service

import (
	"github.com/friends-management/models"
	"github.com/stretchr/testify/mock"
)

type MockSubscribe struct {
	mock.Mock
}

func (_m *MockSubscribe) CreateSubscribe(subscribe *models.Subscribe) error {
	returnArgs := _m.Called(subscribe)
	return returnArgs.Error(0)
}

func (_m *MockSubscribe) CheckExistedSubscribe(requestorId int, targetId int) (bool, error) {
	returnArgs := _m.Called(requestorId, targetId)
	return returnArgs.Bool(0), returnArgs.Error(1)
}

func (_m *MockSubscribe) CheckBlockedByUser(requestorId int, targetId int) (bool, string, error) {
	returnArgs := _m.Called(requestorId, targetId)
	return returnArgs.Bool(0), "", returnArgs.Error(1)
}

func (_m *MockSubscribe) IsBlockedByUser(requestorId int, targetId int) (bool, string, error) {
	returnArgs := _m.Called(requestorId, targetId)
	return returnArgs.Bool(0), "", returnArgs.Error(1)
}

func (_m *MockSubscribe) IsExistedSubscribe(requestorId int, targetId int) (bool, error) {
	returnArgs := _m.Called(requestorId, targetId)
	return returnArgs.Bool(0), returnArgs.Error(1)
}
