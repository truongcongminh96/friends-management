package service

import (
	"errors"
	"github.com/friends-management/models"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFriendService_CreateFriend(t *testing.T) {
	testCases := []struct {
		name          string
		input         *models.Friend
		expectedError error
		mockRepoInput *models.Friend
		mockRepoError error
	}{
		{
			name: "create friend failed",
			input: &models.Friend{
				User1: 1,
				User2: 2,
			},
			expectedError: errors.New("insert friend failed"),
			mockRepoInput: &models.Friend{
				User1: 1,
				User2: 2,
			},
			mockRepoError: errors.New("insert friend failed"),
		},
		{
			name: "create friend successfully",
			input: &models.Friend{
				User1: 3,
				User2: 4,
			},
			expectedError: nil,
			mockRepoInput: &models.Friend{
				User1: 3,
				User2: 4,
			},
			mockRepoError: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			mockFriend := new(MockFriend)
			mockFriend.On("CreateFriend", testCase.mockRepoInput).
				Return(testCase.mockRepoError)

			service := FriendService{
				IFriendRepo: mockFriend,
			}

			// When
			err := service.CreateFriend(testCase.input)

			// Then
			if testCase.expectedError != nil {
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestFriendServices_CheckExistedFriend(t *testing.T) {
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
			name:           "check is existed friend failed",
			input:          []int{1, 2},
			expectedValue:  false,
			expectedError:  errors.New("query database failed"),
			mockRepoInput:  []int{1, 2},
			mockRepoResult: false,
			mockRepoError:  errors.New("query database failed"),
		},
		{
			name:           "check is existed friend successfully",
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
			mockFriend := new(MockFriend)
			mockFriend.On("IsExistedFriend", testCase.mockRepoInput[0], testCase.mockRepoInput[1]).
				Return(testCase.mockRepoResult, testCase.mockRepoError)

			service := FriendService{
				IFriendRepo: mockFriend,
			}

			// When
			result, err := service.CheckExistedFriend(testCase.input[0], testCase.input[1])

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

func TestFriendServices_CheckBlockedByUser(t *testing.T) {
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
			mockFriend := new(MockFriend)
			mockFriend.On("IsBlockedByUser", testCase.mockRepoInput[0], testCase.mockRepoInput[1]).
				Return(testCase.mockRepoResult, testCase.mockRepoError)

			service := FriendService{
				IFriendRepo: mockFriend,
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

func TestFriendServices_GetFriendsList(t *testing.T) {
	type mockGetFriendListByID struct {
		input  int
		result []int
		err    error
	}
	type mockGetBlockedListByID struct {
		input  int
		result []int
		err    error
	}
	type mockGetBlockingListByID struct {
		input  int
		result []int
		err    error
	}
	type mockGetEmailListByIDs struct {
		input  []int
		result []string
		err    error
	}
	testCases := []struct {
		name                string
		input               int
		expectedResult      []string
		expectedErr         error
		mockGetFriendsList  mockGetFriendListByID
		mockGetBlockedList  mockGetBlockedListByID
		mockGetBlockingList mockGetBlockingListByID
		mockGetEmailList    mockGetEmailListByIDs
	}{
		{
			name:           "Get friends list failed with error",
			input:          1,
			expectedResult: nil,
			expectedErr:    errors.New("get friends list failed with error"),
			mockGetFriendsList: mockGetFriendListByID{
				input:  1,
				result: nil,
				err:    errors.New("get friends list failed with error"),
			},
		},
		{
			name:           "get blocked list failed",
			input:          1,
			expectedResult: nil,
			expectedErr:    errors.New("get blocked list failed with error"),
			mockGetFriendsList: mockGetFriendListByID{
				input:  1,
				result: []int{1,2},
				err:    nil,
			},
			mockGetBlockedList: mockGetBlockedListByID{
				input:  1,
				result: []int{3},
				err:    errors.New("get blocked list failed with error"),
			},
		},
		{
			name:           "get blocking list failed",
			input:          1,
			expectedResult: nil,
			expectedErr:    errors.New("get blocking list failed with error"),
			mockGetFriendsList: mockGetFriendListByID{
				input:  1,
				result: []int{2},
				err:    nil,
			},
			mockGetBlockedList: mockGetBlockedListByID{
				input:  1,
				result: []int{3},
				err:    nil,
			},
			mockGetBlockingList: mockGetBlockingListByID{
				input:  1,
				result: nil,
				err:    errors.New("get blocking list failed with error"),
			},
		},
		{
			name:           "Get email list by IDs failed with error",
			input:          1,
			expectedResult: nil,
			expectedErr:    errors.New("get email list by userIDs failed with error"),
			mockGetFriendsList: mockGetFriendListByID{
				input:  1,
				result: []int{2, 3, 4, 5},
				err:    nil,
			},
			mockGetBlockedList: mockGetBlockedListByID{
				input:  1,
				result: []int{3},
				err:    nil,
			},
			mockGetBlockingList: mockGetBlockingListByID{
				input:  1,
				result: []int{4},
				err:    nil,
			},
			mockGetEmailList: mockGetEmailListByIDs{
				input:  []int{2, 5},
				result: nil,
				err:    errors.New("get email list by userIDs failed with error"),
			},
		},
		{
			name:           "Get friend connection list success",
			input:          1,
			expectedResult: []string{"xyz@xyz.com", "xyzk@abc.com"},
			expectedErr:    nil,
			mockGetFriendsList: mockGetFriendListByID{
				input:  1,
				result: []int{2, 3, 4, 5},
				err:    nil,
			},
			mockGetBlockedList: mockGetBlockedListByID{
				input:  1,
				result: []int{3},
				err:    nil,
			},
			mockGetBlockingList: mockGetBlockingListByID{
				input:  1,
				result: []int{4},
				err:    nil,
			},
			mockGetEmailList: mockGetEmailListByIDs{
				input:  []int{2, 5},
				result: []string{"xyz@xyz.com", "xyzk@abc.com"},
				err:    nil,
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			mockFriendRepo := new(MockFriend)
			mockUserRepo := new(MockUser)
			mockFriendRepo.On("GetListFriendId", testCase.mockGetFriendsList.input).
				Return(testCase.mockGetFriendsList.result, testCase.mockGetFriendsList.err)
			mockFriendRepo.On("GetBlockedListByID", testCase.mockGetBlockedList.input).
				Return(testCase.mockGetBlockedList.result, testCase.mockGetBlockedList.err)
			mockFriendRepo.On("GetBlockingListByID", testCase.mockGetBlockingList.input).
				Return(testCase.mockGetBlockingList.result, testCase.mockGetBlockingList.err)
			mockUserRepo.On("GetEmailsByIDs", testCase.mockGetEmailList.input).
				Return(testCase.mockGetEmailList.result, testCase.mockGetEmailList.err)

			service := FriendService{
				IFriendRepo: mockFriendRepo,
				IUserRepo:   mockUserRepo,
			}

			// When
			result, err := service.GetFriendsList(testCase.input)

			// Then
			if testCase.expectedErr != nil {
				require.EqualError(t, err, testCase.expectedErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, testCase.expectedResult, result)
			}
		})
	}
}

func TestFriendServices_GetCommonFriends(t *testing.T) {
	type mockGetFriendsList struct {
		input  int
		result []int
		err    error
	}
	type mockGetBlockedUsers struct {
		input  int
		result []int
		err    error
	}
	type mockGetBlockingUsers struct {
		input  int
		result []int
		err    error
	}
	type mockGetEmailsFromIDs struct {
		input  []int
		result []string
		err    error
	}

	testCases := []struct {
		name                  string
		input                 []int
		expectedValue         []string
		expectedError         error
		mockGetFriendsList1   mockGetFriendsList
		mockGetBlockedUsers1  mockGetBlockedUsers
		mockGetBlockingUsers1 mockGetBlockingUsers
		mockGetEmailsFromIDs1 mockGetEmailsFromIDs
		mockGetFriendsList2   mockGetFriendsList
		mockGetBlockedUsers2  mockGetBlockedUsers
		mockGetBlockingUsers2 mockGetBlockingUsers
		mockGetEmailsFromIDs2 mockGetEmailsFromIDs
	}{
		{
			name:          "get friends list of first user failed",
			input:         []int{1, 2},
			expectedValue: nil,
			expectedError: errors.New("query database failed"),
			mockGetFriendsList1: mockGetFriendsList{
				input:  1,
				result: nil,
				err:    errors.New("query database failed"),
			},
		},
		{
			name:          "get friends list of second user failed",
			input:         []int{1, 2},
			expectedValue: nil,
			expectedError: errors.New("query database failed"),
			mockGetFriendsList1: mockGetFriendsList{
				input:  1,
				result: []int{2, 3, 4, 5},
				err:    nil,
			},
			mockGetBlockedUsers1: mockGetBlockedUsers{
				input:  1,
				result: []int{3},
				err:    nil,
			},
			mockGetBlockingUsers1: mockGetBlockingUsers{
				input:  1,
				result: []int{4},
				err:    nil,
			},
			mockGetEmailsFromIDs1: mockGetEmailsFromIDs{
				input:  []int{2, 5},
				result: []string{"dao@example.com", "mai@example.com"},
				err:    nil,
			},
			mockGetFriendsList2: mockGetFriendsList{
				input:  2,
				result: nil,
				err:    errors.New("query database failed"),
			},
		},
		{
			name:          "get common friends successfully",
			input:         []int{1, 2},
			expectedValue: []string{"dao@example.com"},
			expectedError: nil,
			mockGetFriendsList1: mockGetFriendsList{
				input:  1,
				result: []int{2, 3, 4, 5},
				err:    nil,
			},
			mockGetBlockedUsers1: mockGetBlockedUsers{
				input:  1,
				result: []int{3},
				err:    nil,
			},
			mockGetBlockingUsers1: mockGetBlockingUsers{
				input:  1,
				result: []int{4},
				err:    nil,
			},
			mockGetEmailsFromIDs1: mockGetEmailsFromIDs{
				input:  []int{2, 5},
				result: []string{"dao@example.com", "mai@example.com"},
				err:    nil,
			},

			mockGetFriendsList2: mockGetFriendsList{
				input:  2,
				result: []int{2, 3},
				err:    nil,
			},
			mockGetBlockedUsers2: mockGetBlockedUsers{
				input:  2,
				result: []int{},
				err:    nil,
			},
			mockGetBlockingUsers2: mockGetBlockingUsers{
				input:  2,
				result: []int{},
				err:    nil,
			},
			mockGetEmailsFromIDs2: mockGetEmailsFromIDs{
				input:  []int{2, 3},
				result: []string{"dao@example.com", "thu@example.com"},
				err:    nil,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			mockFriendRepo := new(MockFriend)
			mockUserRepo := new(MockUser)

			mockFriendRepo.On("GetListFriendId", testCase.mockGetFriendsList1.input).
				Return(testCase.mockGetFriendsList1.result, testCase.mockGetFriendsList1.err)
			mockFriendRepo.On("GetBlockedUsers", testCase.mockGetBlockedUsers1.input).
				Return(testCase.mockGetBlockedUsers1.result, testCase.mockGetBlockedUsers1.err)
			mockFriendRepo.On("GetBlockingUsers", testCase.mockGetBlockingUsers1.input).
				Return(testCase.mockGetBlockingUsers1.result, testCase.mockGetBlockingUsers1.err)
			mockUserRepo.On("GetEmailsByIDs", testCase.mockGetEmailsFromIDs1.input).
				Return(testCase.mockGetEmailsFromIDs1.result, testCase.mockGetEmailsFromIDs1.err)

			mockFriendRepo.On("GetListFriendId", testCase.mockGetFriendsList2.input).
				Return(testCase.mockGetFriendsList2.result, testCase.mockGetFriendsList2.err)
			mockFriendRepo.On("GetBlockedUsers", testCase.mockGetBlockedUsers2.input).
				Return(testCase.mockGetBlockedUsers2.result, testCase.mockGetBlockedUsers2.err)
			mockFriendRepo.On("GetBlockingUsers", testCase.mockGetBlockingUsers2.input).
				Return(testCase.mockGetBlockingUsers2.result, testCase.mockGetBlockingUsers2.err)
			mockUserRepo.On("GetEmailsByIDs", testCase.mockGetEmailsFromIDs2.input).
				Return(testCase.mockGetEmailsFromIDs2.result, testCase.mockGetEmailsFromIDs2.err)

			service := FriendService{
				IFriendRepo: mockFriendRepo,
				IUserRepo:   mockUserRepo,
			}

			// When
			result, err := service.GetCommonFriends(testCase.input)

			// Then
			if testCase.expectedError != nil {
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.ElementsMatch(t, result, testCase.expectedValue)
			}
		})
	}
}
