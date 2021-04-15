package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/friends-management/models"
	"github.com/friends-management/service"
	"github.com/stretchr/testify/require"
)

func TestFriendHandlers_CreateFriend(t *testing.T) {
	type mockGetUserIDByEmail struct {
		input  string
		result int
		err    error
	}
	type mockCheckExistedFriend struct {
		input  []int
		result bool
		err    error
	}
	type mockIsBlocked struct {
		input  []int
		result bool
		err    error
	}
	type mockCreateFriend struct {
		input *models.Friend
		err   error
	}
	testCases := []struct {
		name                   string
		requestBody            map[string]interface{}
		expectedResponseBody   string
		expectedStatus         int
		mockGetRequestorId     mockGetUserIDByEmail
		mockGetTargetId        mockGetUserIDByEmail
		mockCheckExistedFriend mockCheckExistedFriend
		mockIsBlocked          mockIsBlocked
		mockCreateFriend       mockCreateFriend
	}{
		{
			name: "decode request body failed",
			requestBody: map[string]interface{}{
				"friends": 1,
			},
			expectedResponseBody: "{\"Err\":null,\"StatusCode\":400,\"StatusText\":\"\",\"Message\":\"Bad request\"}\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "validate request body failed - missing friends",
			requestBody: map[string]interface{}{
				"friends": 0,
			},
			expectedResponseBody: "{\"Err\":null,\"StatusCode\":400,\"StatusText\":\"\",\"Message\":\"Bad request\"}\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "validate request body failed - not enough email address",
			requestBody: map[string]interface{}{
				"friends": []string{"andy@example.com", ""},
			},
			expectedResponseBody: "{\"Err\":{},\"StatusCode\":400,\"StatusText\":\"Bad request\",\"Message\":\"your friend is required\"}\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "get first user id failed",
			requestBody: map[string]interface{}{
				"friends": []string{
					"andy@example.com",
					"john@example.com",
				},
			},
			expectedResponseBody: "{\"Err\":{},\"StatusCode\":500,\"StatusText\":\"Bad request\",\"Message\":\"get user id failed\"}\n",
			expectedStatus:       http.StatusInternalServerError,
			mockGetRequestorId: mockGetUserIDByEmail{
				input:  "andy@example.com",
				result: 0,
				err:    errors.New("get user id failed"),
			},
		},
		{
			name: "first email does not exist",
			requestBody: map[string]interface{}{
				"friends": []string{
					"andy@example.com",
					"john@example.com",
				},
			},
			expectedResponseBody: "{\"Err\":{},\"StatusCode\":400,\"StatusText\":\"Bad request\",\"Message\":\"your email does not exist\"}\n",
			expectedStatus:       http.StatusBadRequest,
			mockGetRequestorId: mockGetUserIDByEmail{
				input:  "andy@example.com",
				result: 0,
				err:    nil,
			},
		},
		{
			name: "get second user failed",
			requestBody: map[string]interface{}{
				"friends": []string{
					"andy@example.com",
					"john@example.com",
				},
			},
			expectedResponseBody: "{\"Err\":{},\"StatusCode\":500,\"StatusText\":\"Bad request\",\"Message\":\"get user id failed\"}\n",
			expectedStatus:       http.StatusInternalServerError,
			mockGetRequestorId: mockGetUserIDByEmail{
				input:  "andy@example.com",
				result: 1,
				err:    nil,
			},
			mockGetTargetId: mockGetUserIDByEmail{
				input:  "john@example.com",
				result: 0,
				err:    errors.New("get user id failed"),
			},
		},
		{
			name: "second email does not exist",
			requestBody: map[string]interface{}{
				"friends": []string{
					"andy@example.com",
					"john@example.com",
				},
			},
			expectedResponseBody: "{\"Err\":{},\"StatusCode\":400,\"StatusText\":\"Bad request\",\"Message\":\"your friend email does not exist\"}\n",
			expectedStatus:       http.StatusBadRequest,
			mockGetRequestorId: mockGetUserIDByEmail{
				input:  "andy@example.com",
				result: 1,
				err:    nil,
			},
			mockGetTargetId: mockGetUserIDByEmail{
				input:  "john@example.com",
				result: 0,
				err:    nil,
			},
		},
		{
			name: "check friend exist failed",
			requestBody: map[string]interface{}{
				"friends": []string{
					"andy@example.com",
					"john@example.com",
				},
			},
			expectedResponseBody: "{\"Err\":{},\"StatusCode\":500,\"StatusText\":\"Bad request\",\"Message\":\"query database failed\"}\n",
			expectedStatus:       http.StatusInternalServerError,
			mockGetRequestorId: mockGetUserIDByEmail{
				input:  "andy@example.com",
				result: 1,
				err:    nil,
			},
			mockGetTargetId: mockGetUserIDByEmail{
				input:  "john@example.com",
				result: 2,
				err:    nil,
			},
			mockCheckExistedFriend: mockCheckExistedFriend{
				input:  []int{1, 2},
				result: false,
				err:    errors.New("query database failed"),
			},
		},
		{
			name: "friend exists",
			requestBody: map[string]interface{}{
				"friends": []string{
					"andy@example.com",
					"john@example.com",
				},
			},
			expectedResponseBody: "{\"Err\":{},\"StatusCode\":208,\"StatusText\":\"Bad request\",\"Message\":\"you are friends\"}\n",
			expectedStatus:       http.StatusAlreadyReported,
			mockGetRequestorId: mockGetUserIDByEmail{
				input:  "andy@example.com",
				result: 1,
				err:    nil,
			},
			mockGetTargetId: mockGetUserIDByEmail{
				input:  "john@example.com",
				result: 2,
				err:    nil,
			},
			mockCheckExistedFriend: mockCheckExistedFriend{
				input:  []int{1, 2},
				result: true,
				err:    nil,
			},
		},
		{
			name: "check is blocked failed",
			requestBody: map[string]interface{}{
				"friends": []string{
					"andy@example.com",
					"john@example.com",
				},
			},
			expectedResponseBody: "{\"Err\":{},\"StatusCode\":500,\"StatusText\":\"Bad request\",\"Message\":\"query database failed\"}\n",
			expectedStatus:       http.StatusInternalServerError,
			mockGetRequestorId: mockGetUserIDByEmail{
				input:  "andy@example.com",
				result: 1,
				err:    nil,
			},
			mockGetTargetId: mockGetUserIDByEmail{
				input:  "john@example.com",
				result: 2,
				err:    nil,
			},
			mockCheckExistedFriend: mockCheckExistedFriend{
				input:  []int{1, 2},
				result: false,
				err:    nil,
			},
			mockIsBlocked: mockIsBlocked{
				input:  []int{1, 2},
				result: false,
				err:    errors.New("query database failed"),
			},
		},
		{
			name: "is blocked",
			requestBody: map[string]interface{}{
				"friends": []string{
					"andy@example.com",
					"john@example.com",
				},
			},
			expectedResponseBody: "{\"Err\":{},\"StatusCode\":412,\"StatusText\":\"Bad request\",\"Message\":\"\"}\n",
			expectedStatus:       http.StatusPreconditionFailed,
			mockGetRequestorId: mockGetUserIDByEmail{
				input:  "andy@example.com",
				result: 1,
				err:    nil,
			},
			mockGetTargetId: mockGetUserIDByEmail{
				input:  "john@example.com",
				result: 2,
				err:    nil,
			},
			mockCheckExistedFriend: mockCheckExistedFriend{
				input:  []int{1, 2},
				result: false,
				err:    nil,
			},
			mockIsBlocked: mockIsBlocked{
				input:  []int{1, 2},
				result: true,
				err:    nil,
			},
		},
		{
			name: "create friend failed",
			requestBody: map[string]interface{}{
				"friends": []string{
					"andy@example.com",
					"john@example.com",
				},
			},
			expectedResponseBody: "{\"Err\":{},\"StatusCode\":500,\"StatusText\":\"Internal server error\",\"Message\":\"create friend failed\"}\n",
			expectedStatus:       http.StatusInternalServerError,
			mockGetRequestorId: mockGetUserIDByEmail{
				input:  "andy@example.com",
				result: 1,
				err:    nil,
			},
			mockGetTargetId: mockGetUserIDByEmail{
				input:  "john@example.com",
				result: 2,
				err:    nil,
			},
			mockCheckExistedFriend: mockCheckExistedFriend{
				input:  []int{1, 2},
				result: false,
				err:    nil,
			},
			mockIsBlocked: mockIsBlocked{
				input:  []int{1, 2},
				result: false,
				err:    nil,
			},
			mockCreateFriend: mockCreateFriend{
				input: &models.Friend{
					User1: 1,
					User2: 2,
				},
				err: errors.New("create friend failed"),
			},
		},
		{
			name: "create friend successfully",
			requestBody: map[string]interface{}{
				"friends": []string{
					"andy@example.com",
					"john@example.com",
				},
			},
			expectedResponseBody: "{\"success\":true}\n",
			expectedStatus:       http.StatusOK,
			mockGetRequestorId: mockGetUserIDByEmail{
				input:  "andy@example.com",
				result: 1,
			},
			mockGetTargetId: mockGetUserIDByEmail{
				input:  "john@example.com",
				result: 2,
			},
			mockCheckExistedFriend: mockCheckExistedFriend{
				input:  []int{1, 2},
				result: false,
			},
			mockIsBlocked: mockIsBlocked{
				input:  []int{1, 2},
				result: false,
			},
			mockCreateFriend: mockCreateFriend{
				input: &models.Friend{
					User1: 1,
					User2: 2,
				},
				err: nil,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			mockFriend := new(service.MockFriend)
			mockUser := new(service.MockUser)

			mockUser.On("GetUserIDByEmail", testCase.mockGetRequestorId.input).
				Return(testCase.mockGetRequestorId.result, testCase.mockGetRequestorId.err)
			mockUser.On("GetUserIDByEmail", testCase.mockGetTargetId.input).
				Return(testCase.mockGetTargetId.result, testCase.mockGetTargetId.err)

			mockFriend.On("CreateFriend", testCase.mockCreateFriend.input).
				Return(testCase.mockCreateFriend.err)

			if testCase.mockCheckExistedFriend.input != nil {
				mockFriend.On("CheckExistedFriend", testCase.mockCheckExistedFriend.input[0], testCase.mockCheckExistedFriend.input[1]).
					Return(testCase.mockCheckExistedFriend.result, testCase.mockCheckExistedFriend.err)
			}
			if testCase.mockIsBlocked.input != nil {
				mockFriend.On("CheckBlockedByUser", testCase.mockIsBlocked.input[0], testCase.mockIsBlocked.input[1]).
					Return(testCase.mockIsBlocked.result, testCase.mockIsBlocked.err)
			}

			handlers := FriendHandlers{
				IFriendService: mockFriend,
				IUserService:   mockUser,
			}

			requestBody, err := json.Marshal(testCase.requestBody)
			if err != nil {
				t.Error(err)
			}

			// When
			req, err := http.NewRequest(http.MethodPost, "/friend", bytes.NewBuffer(requestBody))
			if err != nil {
				t.Error(err)
			}
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(handlers.CreateFriend)
			handler.ServeHTTP(rr, req)

			// Then
			require.Equal(t, testCase.expectedStatus, rr.Code)
			require.Equal(t, testCase.expectedResponseBody, rr.Body.String())

		})
	}
}

func TestFriendHandlers_GetFriendsList(t *testing.T) {
	type mockGetUserIDByEmail struct {
		input  string
		result int
		err    error
	}
	type mockGetFriendsList struct {
		input  int
		result []string
		err    error
	}
	testCases := []struct {
		name                 string
		requestBody          map[string]interface{}
		expectedResponseBody string
		expectedStatus       int
		mockGetUserIDByEmail mockGetUserIDByEmail
		mockGetFriendsList   mockGetFriendsList
	}{
		{
			name: "decode request body failed",
			requestBody: map[string]interface{}{
				"email": 1,
			},
			expectedResponseBody: "{\"Err\":null,\"StatusCode\":400,\"StatusText\":\"\",\"Message\":\"Bad request\"}\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "validate request body failed",
			requestBody: map[string]interface{}{
				"email": "",
			},
			expectedResponseBody: "{\"Err\":{},\"StatusCode\":400,\"StatusText\":\"Bad request\",\"Message\":\"\\\"email\\\" is required\"}\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "get user id failed",
			requestBody: map[string]interface{}{
				"email": "andy@example.com",
			},
			expectedResponseBody: "{\"Err\":{},\"StatusCode\":500,\"StatusText\":\"Bad request\",\"Message\":\"get user id failed\"}\n",
			expectedStatus:       http.StatusInternalServerError,
			mockGetUserIDByEmail: mockGetUserIDByEmail{
				input:  "andy@example.com",
				result: 0,
				err:    errors.New("get user id failed"),
			},
		},
		{
			name: "user does not exist",
			requestBody: map[string]interface{}{
				"email": "andy@example.com",
			},
			expectedResponseBody: "{\"Err\":{},\"StatusCode\":400,\"StatusText\":\"Bad request\",\"Message\":\"email does not exist\"}\n",
			expectedStatus:       http.StatusBadRequest,
			mockGetUserIDByEmail: mockGetUserIDByEmail{
				input:  "andy@example.com",
				result: 0,
				err:    nil,
			},
		},
		{
			name: "get friends list failed",
			requestBody: map[string]interface{}{
				"email": "andy@example.com",
			},
			expectedResponseBody: "{\"Err\":{},\"StatusCode\":500,\"StatusText\":\"Internal server error\",\"Message\":\"get friends list failed\"}\n",
			expectedStatus:       http.StatusInternalServerError,
			mockGetUserIDByEmail: mockGetUserIDByEmail{
				input:  "andy@example.com",
				result: 1,
				err:    nil,
			},
			mockGetFriendsList: mockGetFriendsList{
				input:  1,
				result: nil,
				err:    errors.New("get friends list failed"),
			},
		},
		{
			name: "get friends list successfully",
			requestBody: map[string]interface{}{
				"email": "andy@example.com",
			},
			expectedResponseBody: "{\"success\":true,\"friends\":[\"john@example.com\",\"kate@example.com\"],\"count\":2}\n",
			expectedStatus:       http.StatusOK,
			mockGetUserIDByEmail: mockGetUserIDByEmail{
				input:  "andy@example.com",
				result: 1,
				err:    nil,
			},
			mockGetFriendsList: mockGetFriendsList{
				input:  1,
				result: []string{"john@example.com", "kate@example.com"},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			mockFriend := new(service.MockFriend)
			mockUser := new(service.MockUser)

			mockUser.On("GetUserIDByEmail", testCase.mockGetUserIDByEmail.input).
				Return(testCase.mockGetUserIDByEmail.result, testCase.mockGetUserIDByEmail.err)

			mockFriend.On("GetFriendsList", testCase.mockGetFriendsList.input).
				Return(testCase.mockGetFriendsList.result, testCase.mockGetFriendsList.err)

			handlers := FriendHandlers{
				IFriendService: mockFriend,
				IUserService:   mockUser,
			}

			requestBody, err := json.Marshal(testCase.requestBody)
			if err != nil {
				t.Error(err)
			}

			// When
			req, err := http.NewRequest(http.MethodGet, "/friends/friends-list", bytes.NewBuffer(requestBody))
			if err != nil {
				t.Error(err)
			}
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(handlers.GetFriendsList)
			handler.ServeHTTP(rr, req)

			// Then
			require.Equal(t, testCase.expectedStatus, rr.Code)
			require.Equal(t, testCase.expectedResponseBody, rr.Body.String())

		})
	}
}

func TestFriendHandlers_GetCommonFriends(t *testing.T) {
	type mockGetUserIDByEmail struct {
		input  string
		result int
		err    error
	}
	type mockGetCommonFriends struct {
		input  []int
		result []string
		err    error
	}
	testCases := []struct {
		name                 string
		requestBody          map[string]interface{}
		expectedResponseBody string
		expectedStatus       int
		mockGetUserId1       mockGetUserIDByEmail
		mockGetUserId2       mockGetUserIDByEmail
		mockGetCommonFriends mockGetCommonFriends
	}{
		{
			name: "decode request body failed",
			requestBody: map[string]interface{}{
				"friends": 1,
			},
			expectedResponseBody: "{\"Err\":null,\"StatusCode\":400,\"StatusText\":\"\",\"Message\":\"Bad request\"}\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "validate request body failed - missing friends",
			requestBody: map[string]interface{}{
				"friends": 0,
			},
			expectedResponseBody: "{\"Err\":null,\"StatusCode\":400,\"StatusText\":\"\",\"Message\":\"Bad request\"}\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "validate request body failed - not enough email address",
			requestBody: map[string]interface{}{
				"friends": []string{"andy@example.com", ""},
			},
			expectedResponseBody: "{\"Err\":{},\"StatusCode\":400,\"StatusText\":\"Bad request\",\"Message\":\"your friend is required\"}\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "get first user id failed",
			requestBody: map[string]interface{}{
				"friends": []string{
					"andy@example.com",
					"john@example.com",
				},
			},
			expectedResponseBody: "{\"Err\":{},\"StatusCode\":500,\"StatusText\":\"Bad request\",\"Message\":\"get user id failed\"}\n",
			expectedStatus:       http.StatusInternalServerError,
			mockGetUserId1: mockGetUserIDByEmail{
				input:  "andy@example.com",
				result: 0,
				err:    errors.New("get user id failed"),
			},
		},
		{
			name: "first email does not exist",
			requestBody: map[string]interface{}{
				"friends": []string{
					"andy@example.com",
					"john@example.com",
				},
			},
			expectedResponseBody: "{\"Err\":{},\"StatusCode\":400,\"StatusText\":\"Bad request\",\"Message\":\"the first email does not exist\"}\n",
			expectedStatus:       http.StatusBadRequest,
			mockGetUserId1: mockGetUserIDByEmail{
				input:  "andy@example.com",
				result: 0,
				err:    nil,
			},
		},
		{
			name: "get second user failed",
			requestBody: map[string]interface{}{
				"friends": []string{
					"andy@example.com",
					"john@example.com",
				},
			},
			expectedResponseBody: "{\"Err\":{},\"StatusCode\":500,\"StatusText\":\"Bad request\",\"Message\":\"get user id failed\"}\n",
			expectedStatus:       http.StatusInternalServerError,
			mockGetUserId1: mockGetUserIDByEmail{
				input:  "andy@example.com",
				result: 1,
				err:    nil,
			},
			mockGetUserId2: mockGetUserIDByEmail{
				input:  "john@example.com",
				result: 0,
				err:    errors.New("get user id failed"),
			},
		},
		{
			name: "second email does not exist",
			requestBody: map[string]interface{}{
				"friends": []string{
					"andy@example.com",
					"john@example.com",
				},
			},
			expectedResponseBody: "{\"Err\":{},\"StatusCode\":400,\"StatusText\":\"Bad request\",\"Message\":\"the second email does not exist\"}\n",
			expectedStatus:       http.StatusBadRequest,
			mockGetUserId1: mockGetUserIDByEmail{
				input:  "andy@example.com",
				result: 1,
				err:    nil,
			},
			mockGetUserId2: mockGetUserIDByEmail{
				input:  "john@example.com",
				result: 0,
				err:    nil,
			},
		},
		{
			name: "get common friends failed",
			requestBody: map[string]interface{}{
				"friends": []string{
					"andy@example.com",
					"john@example.com",
				},
			},
			expectedResponseBody: "{\"Err\":{},\"StatusCode\":500,\"StatusText\":\"Internal server error\",\"Message\":\"get common friends failed\"}\n",
			expectedStatus:       http.StatusInternalServerError,
			mockGetUserId1: mockGetUserIDByEmail{
				input:  "andy@example.com",
				result: 1,
				err:    nil,
			},
			mockGetUserId2: mockGetUserIDByEmail{
				input:  "john@example.com",
				result: 2,
				err:    nil,
			},
			mockGetCommonFriends: mockGetCommonFriends{
				input:  []int{1, 2},
				result: nil,
				err:    errors.New("get common friends failed"),
			},
		},
		{
			name: "get common friends successfully",
			requestBody: map[string]interface{}{
				"friends": []string{
					"andy@example.com",
					"john@example.com",
				},
			},
			expectedResponseBody: "{\"success\":true,\"friends\":[\"lisa@example.com\",\"john@example.com\"],\"count\":2}\n",
			expectedStatus:       http.StatusOK,
			mockGetUserId1: mockGetUserIDByEmail{
				input:  "andy@example.com",
				result: 1,
			},
			mockGetUserId2: mockGetUserIDByEmail{
				input:  "john@example.com",
				result: 2,
			},
			mockGetCommonFriends: mockGetCommonFriends{
				input:  []int{1, 2},
				result: []string{"lisa@example.com", "john@example.com"},
				err:    nil,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			mockFriend := new(service.MockFriend)
			mockUser := new(service.MockUser)

			mockUser.On("GetUserIDByEmail", testCase.mockGetUserId1.input).
				Return(testCase.mockGetUserId1.result, testCase.mockGetUserId1.err)
			mockUser.On("GetUserIDByEmail", testCase.mockGetUserId2.input).
				Return(testCase.mockGetUserId2.result, testCase.mockGetUserId2.err)

			mockFriend.On("GetCommonFriends", testCase.mockGetCommonFriends.input).
				Return(testCase.mockGetCommonFriends.result, testCase.mockGetCommonFriends.err)

			handlers := FriendHandlers{
				IFriendService: mockFriend,
				IUserService:   mockUser,
			}

			requestBody, err := json.Marshal(testCase.requestBody)
			if err != nil {
				t.Error(err)
			}

			// When
			req, err := http.NewRequest(http.MethodGet, "/friends/common-friends", bytes.NewBuffer(requestBody))
			if err != nil {
				t.Error(err)
			}
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(handlers.GetCommonFriends)
			handler.ServeHTTP(rr, req)

			// Then
			require.Equal(t, testCase.expectedStatus, rr.Code)
			require.Equal(t, testCase.expectedResponseBody, rr.Body.String())

		})

	}
}

func TestFriendHandlers_ReceiveUpdates(t *testing.T) {
	type mockGetUserIDFromEmail struct {
		input  string
		result int
		err    error
	}
	type mockGetEmailsReceivingUpdates struct {
		inputSender int
		inputText   string
		result      []string
		err         error
	}
	testCases := []struct {
		name                          string
		requestBody                   map[string]interface{}
		expectedResponseBody          string
		expectedStatus                int
		mockGetSenderID               mockGetUserIDFromEmail
		mockGetEmailsReceivingUpdates mockGetEmailsReceivingUpdates
	}{
		{
			name: "decode request body failed",
			requestBody: map[string]interface{}{
				"sender": 1,
			},
			expectedResponseBody: "{\"Err\":null,\"StatusCode\":400,\"StatusText\":\"\",\"Message\":\"Bad request\"}\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "validate request body failed",
			requestBody: map[string]interface{}{
				"sender": "",
			},
			expectedResponseBody: "{\"Err\":null,\"StatusCode\":400,\"StatusText\":\"\",\"Message\":\"Bad request\"}\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "get sender id failed",
			requestBody: map[string]interface{}{
				"sender": "andy@example.com",
			},
			expectedResponseBody: "{\"Err\":{},\"StatusCode\":500,\"StatusText\":\"Bad request\",\"Message\":\"get user id failed\"}\n",
			expectedStatus:       http.StatusInternalServerError,
			mockGetSenderID: mockGetUserIDFromEmail{
				input:  "andy@example.com",
				result: 0,
				err:    errors.New("get user id failed"),
			},
		},
		{
			name: "sender does not exist",
			requestBody: map[string]interface{}{
				"sender": "andy@example.com",
			},
			expectedResponseBody: "{\"Err\":{},\"StatusCode\":400,\"StatusText\":\"Bad request\",\"Message\":\"the sender does not exist\"}\n",
			expectedStatus:       http.StatusBadRequest,
			mockGetSenderID: mockGetUserIDFromEmail{
				input:  "andy@example.com",
				result: 0,
				err:    nil,
			},
		},
		{
			name: "get emails receiving updates failed",
			requestBody: map[string]interface{}{
				"sender": "andy@example.com",
			},
			expectedResponseBody: "{\"Err\":{},\"StatusCode\":500,\"StatusText\":\"Internal server error\",\"Message\":\"get emails receiving updates failed\"}\n",
			expectedStatus:       http.StatusInternalServerError,
			mockGetSenderID: mockGetUserIDFromEmail{
				input:  "andy@example.com",
				result: 1,
				err:    nil,
			},
			mockGetEmailsReceivingUpdates: mockGetEmailsReceivingUpdates{
				inputSender: 1,
				result:      nil,
				err:         errors.New("get emails receiving updates failed"),
			},
		},
		{
			name: "get common friends successfully",
			requestBody: map[string]interface{}{
				"sender": "andy@example.com",
			},
			expectedResponseBody: "{\"success\":true,\"recipients\":[\"lisa@example.com\",\"john@example.com\"]}\n",
			expectedStatus:       http.StatusOK,
			mockGetSenderID: mockGetUserIDFromEmail{
				input:  "andy@example.com",
				result: 1,
			},
			mockGetEmailsReceivingUpdates: mockGetEmailsReceivingUpdates{
				inputSender: 1,
				result:      []string{"lisa@example.com", "john@example.com"},
				err:         nil,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			mockFriendService := new(service.MockFriend)
			mockUserService := new(service.MockUser)

			mockUserService.On("GetUserIDByEmail", testCase.mockGetSenderID.input).
				Return(testCase.mockGetSenderID.result, testCase.mockGetSenderID.err)

			mockFriendService.On("GetEmailsReceiveUpdate",
				testCase.mockGetEmailsReceivingUpdates.inputSender, testCase.mockGetEmailsReceivingUpdates.inputText).
				Return(testCase.mockGetEmailsReceivingUpdates.result, testCase.mockGetEmailsReceivingUpdates.err)

			handlers := FriendHandlers{
				IFriendService: mockFriendService,
				IUserService:   mockUserService,
			}

			requestBody, err := json.Marshal(testCase.requestBody)
			if err != nil {
				t.Error(err)
			}

			// When
			req, err := http.NewRequest(http.MethodGet, "/friends/receive-updates", bytes.NewBuffer(requestBody))
			if err != nil {
				t.Error(err)
			}
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(handlers.ReceiveUpdate)
			handler.ServeHTTP(rr, req)

			// Then
			require.Equal(t, testCase.expectedStatus, rr.Code)
			require.Equal(t, testCase.expectedResponseBody, rr.Body.String())

		})

	}
}
