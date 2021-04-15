package repositories

import (
	"database/sql"
	"errors"
	"github.com/friends-management/helper"
	"github.com/friends-management/models"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFriendRepositories_CreateFriend(t *testing.T) {
	testCases := []struct {
		name          string
		input         *models.Friend
		expectedError error
		fixturePath   string
	}{
		{
			name: "insert friend failed",
			input: &models.Friend{
				User1: 1,
				User2: 2,
			},
			expectedError: errors.New("pq: duplicate key value violates unique constraint \"unique_friends_user1_user2\""),
			fixturePath:   "./testdata/friend/friend.sql",
		},
		{
			name: "insert friend successfully",
			input: &models.Friend{
				User1: 3,
				User2: 4,
			},
			expectedError: nil,
			fixturePath:   "./testdata/friend/friend.sql",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			dbMock := helper.ConnectDb()
			_ = helper.LoadFixture(dbMock, testCase.fixturePath)

			repo := FriendRepo{
				Db: dbMock,
			}

			err := repo.CreateFriend(testCase.input)

			if testCase.expectedError != nil {
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestFriendRepositories_IsExistedFriend(t *testing.T) {
	testCases := []struct {
		name          string
		input         []int
		expectedValue bool
		expectedError error
		fixturePath   string
		mockDb        *sql.DB
	}{
		{
			name:          "query database failed",
			input:         []int{1, 2},
			expectedValue: true,
			expectedError: errors.New("pq: password authentication failed for user \"postgres\""),
			mockDb:        helper.ConnectDbFailed(),
			fixturePath:   "",
		},
		{
			name:          "friend connection exists",
			input:         []int{1, 2},
			expectedValue: true,
			expectedError: nil,
			mockDb:        helper.ConnectDb(),
			fixturePath:   "./testdata/friend/friend.sql",
		},
		{
			name:          "friend connection does not exist",
			input:         []int{3, 4},
			expectedValue: false,
			expectedError: nil,
			mockDb:        helper.ConnectDb(),
			fixturePath:   "./testdata/friend/friend.sql",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			_ = helper.LoadFixture(testCase.mockDb, testCase.fixturePath)

			repo := FriendRepo{
				Db: testCase.mockDb,
			}

			result, err := repo.IsExistedFriend(testCase.input[0], testCase.input[1])

			if testCase.expectedError != nil {
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, result, testCase.expectedValue)
			}
		})
	}
}

func TestFriendRepositories_IsBlockedByUser(t *testing.T) {
	testCases := []struct {
		name          string
		input         []int
		expectedValue bool
		expectedError error
		fixturePath   string
		mockDb        *sql.DB
	}{
		{
			name:          "query database failed",
			input:         []int{1, 2},
			expectedValue: true,
			expectedError: errors.New("pq: password authentication failed for user \"postgres\""),
			mockDb:        helper.ConnectDbFailed(),
			fixturePath:   "",
		},
		{
			name:          "is blocked by the other one",
			input:         []int{1, 2},
			expectedValue: true,
			expectedError: nil,
			mockDb:        helper.ConnectDb(),
			fixturePath:   "./testdata/block/block.sql",
		},
		{
			name:          "is not blocked by the other one",
			input:         []int{3, 4},
			expectedValue: false,
			expectedError: nil,
			mockDb:        helper.ConnectDb(),
			fixturePath:   "./testdata/block/block.sql",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			_ = helper.LoadFixture(testCase.mockDb, testCase.fixturePath)

			repo := FriendRepo{
				Db: testCase.mockDb,
			}

			result, _, err := repo.IsBlockedByUser(testCase.input[0], testCase.input[1])

			if testCase.expectedError != nil {
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, result, testCase.expectedValue)
			}
		})
	}
}

func TestFriendRepositories_GetListFriendId(t *testing.T) {
	testCases := []struct {
		name          string
		input         int
		expectedValue []int
		expectedError error
		fixturePath   string
		mockDb        *sql.DB
	}{
		{
			name:          "query database failed",
			input:         1,
			expectedValue: nil,
			expectedError: errors.New("pq: password authentication failed for user \"postgres\""),
			fixturePath:   "",
			mockDb:        helper.ConnectDbFailed(),
		},
		{
			name:          "get friends list successfully",
			input:         2,
			expectedValue: []int{1, 3},
			expectedError: nil,
			fixturePath:   "./testdata/friend/friend.sql",
			mockDb:        helper.ConnectDb(),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			_ = helper.LoadFixture(testCase.mockDb, testCase.fixturePath)

			repo := FriendRepo{
				Db: testCase.mockDb,
			}

			result, err := repo.GetListFriendId(testCase.input)

			if testCase.expectedError != nil {
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, result, testCase.expectedValue)
			}
		})
	}
}

func TestFriendRepositories_GetIdsBlockedUsers(t *testing.T) {
	testCases := []struct {
		name          string
		input         int
		expectedValue []int
		expectedError error
		fixturePath   string
		mockDb        *sql.DB
	}{
		{
			name:          "query database failed",
			input:         1,
			expectedValue: nil,
			expectedError: errors.New("pq: password authentication failed for user \"postgres\""),
			fixturePath:   "",
			mockDb:        helper.ConnectDbFailed(),
		},
		{
			name:          "get blocked emails successfully",
			input:         1,
			expectedValue: []int{},
			expectedError: nil,
			fixturePath:   "./testdata/block/block.sql",
			mockDb:        helper.ConnectDb(),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			_ = helper.LoadFixture(testCase.mockDb, testCase.fixturePath)

			repo := FriendRepo{
				Db: testCase.mockDb,
			}

			// When
			result, err := repo.GetIdsBlockedUsers(testCase.input)

			// Then
			if testCase.expectedError != nil {
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, result, testCase.expectedValue)
			}
		})
	}
}

func TestFriendRepositories_GetIdsSubscribers(t *testing.T) {
	testCases := []struct {
		name          string
		input         int
		expectedValue []int
		expectedError error
		fixturePath   string
		mockDb        *sql.DB
	}{
		{
			name:          "query database failed",
			input:         1,
			expectedValue: nil,
			expectedError: errors.New("pq: password authentication failed for user \"postgres\""),
			fixturePath:   "",
			mockDb:        helper.ConnectDbFailed(),
		},
		{
			name:          "get blocked emails successfully",
			input:         2,
			expectedValue: []int{1},
			expectedError: nil,
			fixturePath:   "./testdata/subscribe/subscribe.sql",
			mockDb:        helper.ConnectDb(),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			_ = helper.LoadFixture(testCase.mockDb, testCase.fixturePath)

			repo := FriendRepo{
				Db: testCase.mockDb,
			}

			result, err := repo.GetIdsSubscribers(testCase.input)

			if testCase.expectedError != nil {
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, result, testCase.expectedValue)
			}
		})
	}
}
