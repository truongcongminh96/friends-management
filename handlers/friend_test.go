package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/friends-management/models"
	"github.com/friends-management/service"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
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
