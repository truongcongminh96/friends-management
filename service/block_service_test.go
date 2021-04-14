package service

import (
	"errors"
	"github.com/friends-management/models"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBlockService_CreateBlock(t *testing.T) {
	testCases := []struct {
		name          string
		input         *models.Block
		expectedError error
		mockRepoInput *models.Block
		mockRepoError error
	}{
		{
			name: "create Block failed",
			input: &models.Block{
				Requestor: 1,
				Target:    2,
			},
			expectedError: errors.New("insert Block failed"),
			mockRepoInput: &models.Block{
				Requestor: 1,
				Target:    2,
			},
			mockRepoError: errors.New("insert Block failed"),
		},
		{
			name: "create Block successfully",
			input: &models.Block{
				Requestor: 3,
				Target:    4,
			},
			expectedError: nil,
			mockRepoInput: &models.Block{
				Requestor: 3,
				Target:    4,
			},
			mockRepoError: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockBlock := new(MockBlock)
			mockBlock.On("CreateBlock", testCase.mockRepoInput).
				Return(testCase.mockRepoError)

			service := BlockService{
				IBlockRepo: mockBlock,
			}

			err := service.CreateBlock(testCase.input)

			if testCase.expectedError != nil {
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestBlockService_IsExistedBlock(t *testing.T) {
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
			name:           "check is existed block failed",
			input:          []int{1, 2},
			expectedValue:  false,
			expectedError:  errors.New("query database failed"),
			mockRepoInput:  []int{1, 2},
			mockRepoResult: false,
			mockRepoError:  errors.New("query database failed"),
		},
		{
			name:           "check is existed block successfully",
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
			mockBlock := new(MockBlock)
			mockBlock.On("IsExistedBlock", testCase.mockRepoInput[0], testCase.mockRepoInput[1]).
				Return(testCase.mockRepoResult, testCase.mockRepoError)

			service := BlockService{
				IBlockRepo: mockBlock,
			}

			result, err := service.CheckExistedBlock(testCase.input[0], testCase.input[1])

			if testCase.expectedError != nil {
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, result, testCase.expectedValue)
			}
		})
	}
}
