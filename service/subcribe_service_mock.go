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

func (_m *MockSubscribe) CheckExistedSubscribe(requestorID int, targetID int) (bool, error) {
	returnArgs := _m.Called(requestorID, targetID)
	return returnArgs.Bool(0), returnArgs.Error(1)
}

func (_m *MockSubscribe) CheckBlockedByUser(requestorID int, targetID int) (bool, string, error) {
	returnArgs := _m.Called(requestorID, targetID)
	return returnArgs.Bool(0), "", returnArgs.Error(1)
}
