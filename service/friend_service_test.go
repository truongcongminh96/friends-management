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
	type mockGetFriendsList struct {
		input  int
		result []int
		err    error
	}
	type mockGetBlockedByTarget struct {
		input  int
		result []int
		err    error
	}
	type mockGetBlockedUsers struct {
		input  int
		result []int
		err    error
	}
	type mockGetEmailsFromIds struct {
		input  []int
		result []string
		err    error
	}

	testCases := []struct {
		name                   string
		input                  int
		expectedValue          []string
		expectedError          error
		mockGetFriendsList     mockGetFriendsList
		mockGetBlockedByTarget mockGetBlockedByTarget
		mockGetBlockedUsers    mockGetBlockedUsers
		mockGetEmailsFromIds   mockGetEmailsFromIds
	}{
		{
			name:          "get friends list failed",
			input:         1,
			expectedValue: nil,
			expectedError: errors.New("get friends list failed"),
			mockGetFriendsList: mockGetFriendsList{
				input:  1,
				result: nil,
				err:    errors.New("get friends list failed"),
			},
		},
		{
			name:          "get blocked users failed",
			input:         1,
			expectedValue: nil,
			expectedError: errors.New("get blocked users failed"),
			mockGetFriendsList: mockGetFriendsList{
				input:  1,
				result: []int{2},
				err:    nil,
			},
			mockGetBlockedByTarget: mockGetBlockedByTarget{
				input:  1,
				result: nil,
				err:    errors.New("get blocked users failed"),
			},
		},
		{
			name:          "get blocking users failed",
			input:         1,
			expectedValue: nil,
			expectedError: errors.New("get blocking users failed"),
			mockGetFriendsList: mockGetFriendsList{
				input:  1,
				result: []int{2},
				err:    nil,
			},
			mockGetBlockedByTarget: mockGetBlockedByTarget{
				input:  1,
				result: []int{3},
				err:    nil,
			},
			mockGetBlockedUsers: mockGetBlockedUsers{
				input:  1,
				result: nil,
				err:    errors.New("get blocking users failed"),
			},
		},
		{
			name:          "get emails from userIds failed",
			input:         1,
			expectedValue: nil,
			expectedError: errors.New("get emails from userIds failed"),
			mockGetFriendsList: mockGetFriendsList{
				input:  1,
				result: []int{2, 3, 4, 5},
				err:    nil,
			},
			mockGetBlockedByTarget: mockGetBlockedByTarget{
				input:  1,
				result: []int{3},
				err:    nil,
			},
			mockGetBlockedUsers: mockGetBlockedUsers{
				input:  1,
				result: []int{4},
				err:    nil,
			},
			mockGetEmailsFromIds: mockGetEmailsFromIds{
				input:  []int{2, 5},
				result: nil,
				err:    errors.New("get emails from userIds failed"),
			},
		},
		{
			name:          "get friends list successfully",
			input:         1,
			expectedValue: []string{"andy@example.com", "john@example.com"},
			expectedError: nil,
			mockGetFriendsList: mockGetFriendsList{
				input:  1,
				result: []int{2, 3, 4, 5},
				err:    nil,
			},
			mockGetBlockedByTarget: mockGetBlockedByTarget{
				input:  1,
				result: []int{3},
				err:    nil,
			},
			mockGetBlockedUsers: mockGetBlockedUsers{
				input:  1,
				result: []int{4},
				err:    nil,
			}, mockGetEmailsFromIds: mockGetEmailsFromIds{
			input:  []int{2, 5},
			result: []string{"andy@example.com", "john@example.com"},
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
			mockFriendRepo.On("GetIdsBlockedByTarget", testCase.mockGetBlockedByTarget.input).
				Return(testCase.mockGetBlockedByTarget.result, testCase.mockGetBlockedByTarget.err)
			mockFriendRepo.On("GetIdsBlockedUsers", testCase.mockGetBlockedUsers.input).
				Return(testCase.mockGetBlockedUsers.result, testCase.mockGetBlockedUsers.err)
			mockUserRepo.On("GetEmailsByIDs", testCase.mockGetEmailsFromIds.input).
				Return(testCase.mockGetEmailsFromIds.result, testCase.mockGetEmailsFromIds.err)

			services := FriendService{
				IFriendRepo: mockFriendRepo,
				IUserRepo:   mockUserRepo,
			}

			// When
			result, err := services.GetFriendsList(testCase.input)

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

func TestFriendServices_GetCommonFriends(t *testing.T) {
	type mockGetFriendsList struct {
		input  int
		result []int
		err    error
	}
	type mockGetBlockedByTarget struct {
		input  int
		result []int
		err    error
	}
	type mockGetBlockedUsers struct {
		input  int
		result []int
		err    error
	}
	type mockGetEmailsFromIds struct {
		input  []int
		result []string
		err    error
	}

	testCases := []struct {
		name                    string
		input                   []int
		expectedValue           []string
		expectedError           error
		mockGetFriendsList1     mockGetFriendsList
		mockGetBlockedByTarget1 mockGetBlockedByTarget
		mockGetBlockedUsers1    mockGetBlockedUsers
		mockGetEmailsFromIds1   mockGetEmailsFromIds
		mockGetFriendsList2     mockGetFriendsList
		mockGetBlockedByTarget2 mockGetBlockedByTarget
		mockGetBlockedUsers2    mockGetBlockedUsers
		mockGetEmailsFromIds2   mockGetEmailsFromIds
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
			mockGetBlockedByTarget1: mockGetBlockedByTarget{
				input:  1,
				result: []int{3},
				err:    nil,
			},
			mockGetBlockedUsers1: mockGetBlockedUsers{
				input:  1,
				result: []int{4},
				err:    nil,
			},
			mockGetEmailsFromIds1: mockGetEmailsFromIds{
				input:  []int{2, 5},
				result: []string{"andy@example.com", "john@example.com"},
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
			expectedValue: []string{"andy@example.com"},
			expectedError: nil,
			mockGetFriendsList1: mockGetFriendsList{
				input:  1,
				result: []int{2, 3, 4, 5},
				err:    nil,
			},
			mockGetBlockedByTarget1: mockGetBlockedByTarget{
				input:  1,
				result: []int{3},
				err:    nil,
			},
			mockGetBlockedUsers1: mockGetBlockedUsers{
				input:  1,
				result: []int{4},
				err:    nil,
			},
			mockGetEmailsFromIds1: mockGetEmailsFromIds{
				input:  []int{2, 5},
				result: []string{"andy@example.com", "john@example.com"},
				err:    nil,
			},

			mockGetFriendsList2: mockGetFriendsList{
				input:  2,
				result: []int{2, 3},
				err:    nil,
			},
			mockGetBlockedByTarget2: mockGetBlockedByTarget{
				input:  2,
				result: []int{},
				err:    nil,
			},
			mockGetBlockedUsers2: mockGetBlockedUsers{
				input:  2,
				result: []int{},
				err:    nil,
			},
			mockGetEmailsFromIds2: mockGetEmailsFromIds{
				input:  []int{2, 3},
				result: []string{"andy@example.com", "kate@example.com"},
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
			mockFriendRepo.On("GetIdsBlockedByTarget", testCase.mockGetBlockedByTarget1.input).
				Return(testCase.mockGetBlockedByTarget1.result, testCase.mockGetBlockedByTarget1.err)
			mockFriendRepo.On("GetIdsBlockedUsers", testCase.mockGetBlockedUsers1.input).
				Return(testCase.mockGetBlockedUsers1.result, testCase.mockGetBlockedUsers1.err)
			mockUserRepo.On("GetEmailsByIDs", testCase.mockGetEmailsFromIds1.input).
				Return(testCase.mockGetEmailsFromIds1.result, testCase.mockGetEmailsFromIds1.err)

			mockFriendRepo.On("GetListFriendId", testCase.mockGetFriendsList2.input).
				Return(testCase.mockGetFriendsList2.result, testCase.mockGetFriendsList2.err)
			mockFriendRepo.On("GetIdsBlockedByTarget", testCase.mockGetBlockedByTarget2.input).
				Return(testCase.mockGetBlockedByTarget2.result, testCase.mockGetBlockedByTarget2.err)
			mockFriendRepo.On("GetIdsBlockedUsers", testCase.mockGetBlockedUsers2.input).
				Return(testCase.mockGetBlockedUsers2.result, testCase.mockGetBlockedUsers2.err)
			mockUserRepo.On("GetEmailsByIDs", testCase.mockGetEmailsFromIds2.input).
				Return(testCase.mockGetEmailsFromIds2.result, testCase.mockGetEmailsFromIds2.err)

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

func TestFriendServices_GetEmailsReceiveUpdate(t *testing.T) {
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
	type mockGetSubscribers struct {
		input  int
		result []int
		err    error
	}
	type mockGetEmailsFromIds struct {
		input  []int
		result []string
		err    error
	}

	testCases := []struct {
		name                  string
		inputSender           int
		inputText             string
		expectedValue         []string
		expectedError         error
		mockGetBlockedUsers   mockGetBlockedUsers
		mockGetEmailsFromIds1 mockGetEmailsFromIds
		mockGetFriendsList    mockGetFriendsList
		mockGetEmailsFromIds2 mockGetEmailsFromIds
		mockGetSubscribers    mockGetSubscribers
		mockGetEmailsFromIds3 mockGetEmailsFromIds
	}{
		{
			name:          "get blocked users failed",
			inputSender:   1,
			expectedValue: nil,
			expectedError: errors.New("query database failed"),
			mockGetBlockedUsers: mockGetBlockedUsers{
				input:  1,
				result: nil,
				err:    errors.New("query database failed"),
			},
		},
		{
			name:          "get emails from blocked user Ids failed",
			inputSender:   1,
			expectedValue: nil,
			expectedError: errors.New("query database failed"),
			mockGetBlockedUsers: mockGetBlockedUsers{
				input:  1,
				result: []int{2},
				err:    nil,
			},
			mockGetEmailsFromIds1: mockGetEmailsFromIds{
				input:  []int{2},
				result: nil,
				err:    errors.New("query database failed"),
			},
		},
		{
			name:          "get friends failed",
			inputSender:   1,
			expectedValue: nil,
			expectedError: errors.New("query database failed"),
			mockGetBlockedUsers: mockGetBlockedUsers{
				input:  1,
				result: []int{2},
				err:    nil,
			},
			mockGetEmailsFromIds1: mockGetEmailsFromIds{
				input:  []int{2},
				result: []string{"andy@example.com"},
				err:    nil,
			},
			mockGetFriendsList: mockGetFriendsList{
				input:  1,
				result: nil,
				err:    errors.New("query database failed"),
			},
		},
		{
			name:          "get emails from friend Ids failed",
			inputSender:   1,
			expectedValue: nil,
			expectedError: errors.New("query database failed"),
			mockGetBlockedUsers: mockGetBlockedUsers{
				input:  1,
				result: []int{2},
				err:    nil,
			},
			mockGetEmailsFromIds1: mockGetEmailsFromIds{
				input:  []int{2},
				result: []string{"andy@example.com"},
				err:    nil,
			},
			mockGetFriendsList: mockGetFriendsList{
				input:  1,
				result: []int{3},
				err:    nil,
			},
			mockGetEmailsFromIds2: mockGetEmailsFromIds{
				input:  []int{3},
				result: nil,
				err:    errors.New("query database failed"),
			},
		},
		{
			name:          "get subscribers failed",
			inputSender:   1,
			expectedValue: nil,
			expectedError: errors.New("query database failed"),
			mockGetBlockedUsers: mockGetBlockedUsers{
				input:  1,
				result: []int{2},
				err:    nil,
			},
			mockGetEmailsFromIds1: mockGetEmailsFromIds{
				input:  []int{2},
				result: []string{"andy@example.com"},
				err:    nil,
			},
			mockGetFriendsList: mockGetFriendsList{
				input:  1,
				result: []int{3},
				err:    nil,
			},
			mockGetEmailsFromIds2: mockGetEmailsFromIds{
				input:  []int{3},
				result: []string{"john@example"},
				err:    nil,
			},
			mockGetSubscribers: mockGetSubscribers{
				input:  1,
				result: nil,
				err:    errors.New("query database failed"),
			},
		},
		{
			name:          "get emails from subscriber Ids failed",
			inputSender:   1,
			expectedValue: nil,
			expectedError: errors.New("query database failed"),
			mockGetBlockedUsers: mockGetBlockedUsers{
				input:  1,
				result: []int{2},
				err:    nil,
			},
			mockGetEmailsFromIds1: mockGetEmailsFromIds{
				input:  []int{2},
				result: []string{"andy@example.com"},
				err:    nil,
			},
			mockGetFriendsList: mockGetFriendsList{
				input:  1,
				result: []int{3},
				err:    nil,
			},
			mockGetEmailsFromIds2: mockGetEmailsFromIds{
				input:  []int{3},
				result: []string{"john@example"},
				err:    nil,
			},
			mockGetSubscribers: mockGetSubscribers{
				input:  1,
				result: []int{4},
				err:    nil,
			},
			mockGetEmailsFromIds3: mockGetEmailsFromIds{
				input:  []int{4},
				result: nil,
				err:    errors.New("query database failed"),
			},
		},
		{
			name:          "get emails receiving updates successfully",
			inputSender:   1,
			inputText:     "hello kate@gmail.com",
			expectedValue: []string{"john@example.com", "lisa@example.com", "kate@gmail.com"},
			expectedError: nil,
			mockGetBlockedUsers: mockGetBlockedUsers{
				input:  1,
				result: []int{2},
				err:    nil,
			},
			mockGetEmailsFromIds1: mockGetEmailsFromIds{
				input:  []int{2},
				result: []string{"andy@example.com"},
				err:    nil,
			},
			mockGetFriendsList: mockGetFriendsList{
				input:  1,
				result: []int{3},
				err:    nil,
			},
			mockGetEmailsFromIds2: mockGetEmailsFromIds{
				input:  []int{3},
				result: []string{"john@example.com"},
				err:    nil,
			},
			mockGetSubscribers: mockGetSubscribers{
				input:  1,
				result: []int{4},
				err:    nil,
			},
			mockGetEmailsFromIds3: mockGetEmailsFromIds{
				input:  []int{4},
				result: []string{"lisa@example.com"},
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
			mockFriendRepo.On("GetIdsBlockedUsers", testCase.mockGetBlockedUsers.input).
				Return(testCase.mockGetBlockedUsers.result, testCase.mockGetBlockedUsers.err)
			mockFriendRepo.On("GetIdsSubscribers", testCase.mockGetSubscribers.input).
				Return(testCase.mockGetSubscribers.result, testCase.mockGetSubscribers.err)
			mockUserRepo.On("GetEmailsByIDs", testCase.mockGetEmailsFromIds1.input).
				Return(testCase.mockGetEmailsFromIds1.result, testCase.mockGetEmailsFromIds1.err)
			mockUserRepo.On("GetEmailsByIDs", testCase.mockGetEmailsFromIds2.input).
				Return(testCase.mockGetEmailsFromIds2.result, testCase.mockGetEmailsFromIds2.err)
			mockUserRepo.On("GetEmailsByIDs", testCase.mockGetEmailsFromIds3.input).
				Return(testCase.mockGetEmailsFromIds3.result, testCase.mockGetEmailsFromIds3.err)

			service := FriendService{
				IFriendRepo: mockFriendRepo,
				IUserRepo:   mockUserRepo,
			}

			// When
			result, err := service.GetEmailsReceiveUpdate(testCase.inputSender, testCase.inputText)

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
