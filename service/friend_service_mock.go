package service

import (
	"github.com/friends-management/models"
	"github.com/stretchr/testify/mock"
)

type MockFriend struct {
	mock.Mock
}

func (_m *MockFriend) IsExistedFriend(userId1 int, userId2 int) (bool, error) {
	returnArgs := _m.Called(userId1, userId2)
	return returnArgs.Bool(0), returnArgs.Error(1)
}

func (_m *MockFriend) GetListFriendId(userId int) ([]int, error) {
	returnArgs := _m.Called(userId)
	return returnArgs.Get(0).([]int), returnArgs.Error(1)
}

func (_m *MockFriend) IsBlockedByUser(requestorId int, targetId int) (bool, string, error) {
	returnArgs := _m.Called(requestorId, targetId)
	return returnArgs.Bool(0), "", returnArgs.Error(1)
}

func (_m *MockFriend) GetIdsBlockedUsers(userId int) ([]int, error) {
	returnArgs := _m.Called(userId)
	return returnArgs.Get(0).([]int), returnArgs.Error(1)
}

func (_m *MockFriend) GetIdsSubscribers(userId int) ([]int, error) {
	returnArgs := _m.Called(userId)
	return returnArgs.Get(0).([]int), returnArgs.Error(1)
}

func (_m *MockFriend) CreateFriend(friend *models.Friend) error {
	returnArgs := _m.Called(friend)
	return returnArgs.Error(0)
}

func (_m *MockFriend) CheckExistedFriend(userId1 int, userId2 int) (bool, error) {
	returnArgs := _m.Called(userId1, userId2)
	return returnArgs.Bool(0), returnArgs.Error(1)
}

func (_m *MockFriend) CheckBlockedByUser(requestorId int, targetId int) (bool, string, error) {
	returnArgs := _m.Called(requestorId, targetId)
	return returnArgs.Bool(0), "", returnArgs.Error(1)
}

func (_m *MockFriend) GetFriendsList(userId int) ([]string, error) {
	returnArgs := _m.Called(userId)
	return returnArgs.Get(0).([]string), returnArgs.Error(1)
}

func (_m *MockFriend) GetCommonFriends(Ids []int) ([]string, error) {
	returnArgs := _m.Called(Ids)
	return returnArgs.Get(0).([]string), returnArgs.Error(1)
}

func (_m *MockFriend) GetEmailsReceiveUpdate(senderId int, text string) ([]string, error) {
	returnArgs := _m.Called(senderId, text)
	return returnArgs.Get(0).([]string), returnArgs.Error(1)
}
