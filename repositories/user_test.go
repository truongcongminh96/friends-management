package repositories

import (
	"errors"
	"github.com/friends-management/helper"
	"github.com/friends-management/models"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserRepositories_CreateUser(t *testing.T) {
	testCases := []struct {
		name          string
		input         *models.User
		expectedError error
		fixturePath   string
	}{
		{
			name: "insert user failed",
			input: &models.User{
				Email: "andy@example",
			},
			expectedError: errors.New("pq: duplicate key value violates unique constraint \"unique_user_email\""),
			fixturePath:   "./testdata/user/user.sql",
		},
		{
			name: "insert user successfully",
			input: &models.User{
				Email: "kate@example",
			},
			expectedError: nil,
			fixturePath:   "./testdata/truncate_table.sql",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			dbMock := helper.ConnectDb()
			_ = helper.LoadFixture(dbMock, testCase.fixturePath)

			repo := UserRepo{
				Db: dbMock,
			}

			// When
			err := repo.CreateUser(testCase.input)

			// Then
			if testCase.expectedError != nil {
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
