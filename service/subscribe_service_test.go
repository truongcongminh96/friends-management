package service

import (
	"errors"
	"github.com/friends-management/models"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSubscribeService_CreateSubscribe(t *testing.T) {
	testCases := []struct {
		name          string
		input         *models.Subscribe
		expectedError error
		mockRepoInput *models.Subscribe
		mockRepoError error
	}{
		{
			name: "create Subscribe failed",
			input: &models.Subscribe{
				Requestor: 1,
				Target:    2,
			},
			expectedError: errors.New("insert Subscribe failed"),
			mockRepoInput: &models.Subscribe{
				Requestor: 1,
				Target:    2,
			},
			mockRepoError: errors.New("insert Subscribe failed"),
		},
		{
			name: "create Subscribe successfully",
			input: &models.Subscribe{
				Requestor: 3,
				Target:    4,
			},
			expectedError: nil,
			mockRepoInput: &models.Subscribe{
				Requestor: 3,
				Target:    4,
			},
			mockRepoError: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockSubscribe := new(MockSubscribe)
			mockSubscribe.On("CreateSubscribe", testCase.mockRepoInput).
				Return(testCase.mockRepoError)

			service := SubscribeService{
				ISubscribeRepo: mockSubscribe,
			}

			err := service.CreateSubscribe(testCase.input)

			if testCase.expectedError != nil {
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestSubscribeService_CheckExistedSubscribe(t *testing.T) {
	testCases := []struct {
		name           string
		input          []int
		expectedValue  bool
		expectedError  error
		mockRepoInput  []int
		mockRepoResult bool
		mockRepoError  error
	}{
		{
			name:           "check is existed subscribe failed",
			input:          []int{1, 2},
			expectedValue:  false,
			expectedError:  errors.New("query database failed"),
			mockRepoInput:  []int{1, 2},
			mockRepoResult: false,
			mockRepoError:  errors.New("query database failed"),
		},
		{
			name:           "check is existed subscribe successfully",
			input:          []int{1, 2},
			expectedError:  nil,
			expectedValue:  true,
			mockRepoInput:  []int{1, 2},
			mockRepoResult: true,
			mockRepoError:  nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockSubscribe := new(MockSubscribe)
			mockSubscribe.On("IsExistedSubscribe", testCase.mockRepoInput[0], testCase.mockRepoInput[1]).
				Return(testCase.mockRepoResult, testCase.mockRepoError)

			service := SubscribeService{
				ISubscribeRepo: mockSubscribe,
			}

			result, err := service.CheckExistedSubscribe(testCase.input[0], testCase.input[1])

			if testCase.expectedError != nil {
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, result, testCase.expectedValue)
			}
		})
	}
}

func TestSubscribeService_CheckBlockedByUser(t *testing.T) {
	testCases := []struct {
		name           string
		input          []int
		expectedValue  bool
		expectedError  error
		mockRepoInput  []int
		mockRepoResult bool
		mockRepoError  error
	}{
		{
			name:           "check is blocked failed",
			input:          []int{1, 2},
			expectedValue:  false,
			expectedError:  errors.New("query database failed"),
			mockRepoInput:  []int{1, 2},
			mockRepoResult: false,
			mockRepoError:  errors.New("query database failed"),
		},
		{
			name:           "check is blocked successfully",
			input:          []int{1, 2},
			expectedError:  nil,
			expectedValue:  true,
			mockRepoInput:  []int{1, 2},
			mockRepoResult: true,
			mockRepoError:  nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			mockSubscribe := new(MockSubscribe)
			mockSubscribe.On("IsBlockedByUser", testCase.mockRepoInput[0], testCase.mockRepoInput[1]).
				Return(testCase.mockRepoResult, testCase.mockRepoError)

			service := SubscribeService{
				ISubscribeRepo: mockSubscribe,
			}

			// When
			result, _, err := service.CheckBlockedByUser(testCase.input[0], testCase.input[1])

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
