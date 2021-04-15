package service

import (
	"github.com/friends-management/models"
	"github.com/stretchr/testify/mock"
)

type MockUser struct {
	mock.Mock
}

func (_m *MockUser) GetEmailsByIDs(listFriendId []int) ([]string, error) {
	returnArgs := _m.Called(listFriendId)
	return returnArgs.Get(0).([]string), returnArgs.Error(1)
}

func (_m *MockUser) GetUserIDByEmail(email string) (int, error) {
	returnArgs := _m.Called(email)
	return returnArgs.Int(0), returnArgs.Error(1)
}

func (_m *MockUser) CreateUser(userRepoInput *models.User) error {
	args := _m.Called(userRepoInput)
	var r error
	if args.Get(0) != nil {
		r = args.Get(0).(error)
	}
	return r
}

func (_m *MockUser) IsExistedUser(email string) (bool, error) {
	args := _m.Called(email)
	r0 := args.Get(0).(bool)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}
