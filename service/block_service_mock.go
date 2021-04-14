package service

import (
	"github.com/friends-management/models"
	"github.com/stretchr/testify/mock"
)

type MockBlock struct {
	mock.Mock
}

func (_m *MockBlock) CheckExistedBlock(requestorID int, targetID int) (bool, error) {
	returnArgs := _m.Called(requestorID, targetID)
	return returnArgs.Bool(0), returnArgs.Error(1)
}

func (_m *MockBlock) CreateBlock(blocking *models.Block) error {
	returnArgs := _m.Called(blocking)
	return returnArgs.Error(0)
}
