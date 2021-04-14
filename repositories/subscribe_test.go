package repositories

import (
	"database/sql"
	"errors"
	"github.com/friends-management/helper"
	"github.com/friends-management/models"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSubscribeRepositories_CreateSubscribe(t *testing.T) {
	testCases := []struct {
		name          string
		input         *models.Subscribe
		expectedError error
		fixturePath   string
	}{
		{
			name: "insert subscribe failed",
			input: &models.Subscribe{
				Requestor: 1,
				Target:    5,
			},
			expectedError: errors.New("pq: insert or update on table \"subscribe\" violates foreign key constraint \"subscribe_target_fkey\""),
			fixturePath:   "./testdata/subscribe/subscribe.sql",
		},
		{
			name: "insert subscribe successfully",
			input: &models.Subscribe{
				Requestor: 3,
				Target:    4,
			},
			expectedError: nil,
			fixturePath:   "./testdata/subscribe/subscribe.sql",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			dbMock := helper.ConnectDb()
			_ = helper.LoadFixture(dbMock, testCase.fixturePath)

			repo := SubscribeRepo{
				Db: dbMock,
			}

			err := repo.CreateSubscribe(testCase.input)

			if testCase.expectedError != nil {
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestSubscribeRepositories_IsExistedSubscribe(t *testing.T) {
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
			name:          "subscribe exists",
			input:         []int{1, 2},
			expectedValue: true,
			expectedError: nil,
			mockDb:        helper.ConnectDb(),
			fixturePath:   "./testdata/subscribe/subscribe.sql",
		},
		{
			name:          "subscribe does not exist",
			input:         []int{3, 4},
			expectedValue: false,
			expectedError: nil,
			mockDb:        helper.ConnectDb(),
			fixturePath:   "./testdata/subscribe/subscribe.sql",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			_ = helper.LoadFixture(testCase.mockDb, testCase.fixturePath)

			repo := SubscribeRepo{
				Db: testCase.mockDb,
			}

			result, err := repo.IsExistedSubscribe(testCase.input[0], testCase.input[1])

			if testCase.expectedError != nil {
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, result, testCase.expectedValue)
			}
		})
	}
}

func TestSubscribeRepositories_IsBlockedByUser(t *testing.T) {
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
			name:          "is blocked",
			input:         []int{1, 2},
			expectedValue: true,
			expectedError: nil,
			mockDb:        helper.ConnectDb(),
			fixturePath:   "./testdata/block/block.sql",
		},
		{
			name:          "is not blocked",
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

			repo := SubscribeRepo{
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
