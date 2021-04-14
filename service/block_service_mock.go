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

func (_m *MockBlock) CreateBlock(block *models.Block) error {
	returnArgs := _m.Called(block)
	return returnArgs.Error(0)
}

func (_m *MockBlock) IsExistedBlock(reqestorID int, targetID int) (bool, error) {
	returnArgs := _m.Called(reqestorID, targetID)
	return returnArgs.Bool(0), returnArgs.Error(1)
}
