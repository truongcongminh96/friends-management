package service

import (
	"errors"
	"testing"

	"github.com/friends-management/models"
	"github.com/stretchr/testify/require"
)

func TestUserServices_CreateUser(t *testing.T) {
	testCases := []struct {
		name          string
		input         *models.User
		expectedError error
		mockRepoInput *models.User
		mockRepoError error
	}{
		{
			name: "create user failed",
			input: &models.User{
				Email: "andy@example",
			},
			expectedError: errors.New("insert user failed"),
			mockRepoInput: &models.User{
				Email: "andy@example",
			},
			mockRepoError: errors.New("insert user failed"),
		},
		{
			name: "create user successfully",
			input: &models.User{
				Email: "kate@example.com",
			},
			expectedError: nil,
			mockRepoInput: &models.User{
				Email: "kate@example.com",
			},
			mockRepoError: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			mockUser := new(MockUser)
			mockUser.On("CreateUser", testCase.mockRepoInput).
				Return(testCase.mockRepoError)

			services := UserService{
				IUserRepo: mockUser,
			}

			// When
			err := services.CreateUser(testCase.input)

			// Then
			if testCase.expectedError != nil {
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestUserServices_IsExistedUser(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		expectedValue  bool
		expectedError  error
		mockRepoInput  string
		mockRepoResult bool
		mockRepoError  error
	}{
		{
			name:           "check is existed user failed",
			input:          "andy@example",
			expectedValue:  false,
			expectedError:  errors.New("query database failed"),
			mockRepoInput:  "andy@example",
			mockRepoResult: false,
			mockRepoError:  errors.New("query database failed"),
		},
		{
			name:           "check is existed user successfully",
			input:          "kate@example",
			expectedError:  nil,
			expectedValue:  true,
			mockRepoInput:  "kate@example",
			mockRepoResult: true,
			mockRepoError:  nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockUser := new(MockUser)
			mockUser.On("IsExistedUser", testCase.mockRepoInput).
				Return(testCase.mockRepoResult, testCase.mockRepoError)

			services := UserService{
				IUserRepo: mockUser,
			}

			result, err := services.IsExistedUser(testCase.input)

			if testCase.expectedError != nil {
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, result, testCase.expectedValue)
			}
		})
	}
}
