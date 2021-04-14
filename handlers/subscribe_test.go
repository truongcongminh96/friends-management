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

func TestSubscribeHandlers_CreateSubscribe(t *testing.T) {
	type mockGetUserIDByEmail struct {
		input  string
		result int
		err    error
	}
	type mockCheckExistedSubscribe struct {
		input  []int
		result bool
		err    error
	}
	type mockIsBlocked struct {
		input  []int
		result bool
		err    error
	}
	type mockCreateSubscribe struct {
		input *models.Subscribe
		err   error
	}
	testCases := []struct {
		name                      string
		requestBody               map[string]interface{}
		expectedResponseBody      string
		expectedStatus            int
		mockGetRequestorId        mockGetUserIDByEmail
		mockGetTargetId           mockGetUserIDByEmail
		mockCheckExistedSubscribe mockCheckExistedSubscribe
		mockIsBlocked             mockIsBlocked
		mockCreateSubscribe       mockCreateSubscribe
	}{
		{
			name: "decode request body failed",
			requestBody: map[string]interface{}{
				"requestor": 1,
			},
			expectedResponseBody: "{\"Err\":null,\"StatusCode\":400,\"StatusText\":\"\",\"Message\":\"Bad request\"}\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "validate request body failed - missing target",
			requestBody: map[string]interface{}{
				"requestor": "andy@example.com",
			},
			expectedResponseBody: "{\"Err\":null,\"StatusCode\":400,\"StatusText\":\"\",\"Message\":\"Bad request\"}\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "validate request body failed - missing requestor",
			requestBody: map[string]interface{}{
				"target": "andy@example.com",
			},
			expectedResponseBody: "{\"Err\":null,\"StatusCode\":400,\"StatusText\":\"\",\"Message\":\"Bad request\"}\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "get requestor id failed",
			requestBody: map[string]interface{}{
				"requestor": "andy@example.com",
				"target":    "john@example.com",
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
			name: "requestor does not exist",
			requestBody: map[string]interface{}{
				"requestor": "andy@example.com",
				"target":    "john@example.com",
			},
			expectedResponseBody: "{\"Err\":{},\"StatusCode\":400,\"StatusText\":\"Bad request\",\"Message\":\"email requestor does not exist\"}\n",
			expectedStatus:       http.StatusBadRequest,
			mockGetRequestorId: mockGetUserIDByEmail{
				input:  "andy@example.com",
				result: 0,
				err:    nil,
			},
		},
		{
			name: "get requestor id failed",
			requestBody: map[string]interface{}{
				"requestor": "andy@example.com",
				"target":    "john@example.com",
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
			name: "target does not exist",
			requestBody: map[string]interface{}{
				"requestor": "andy@example.com",
				"target":    "john@example.com",
			},
			expectedResponseBody: "{\"Err\":{},\"StatusCode\":400,\"StatusText\":\"Bad request\",\"Message\":\"email target does not exist\"}\n",
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
			name: "check Subscribe exist failed",
			requestBody: map[string]interface{}{
				"requestor": "andy@example.com",
				"target":    "john@example.com",
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
			mockCheckExistedSubscribe: mockCheckExistedSubscribe{
				input:  []int{1, 2},
				result: false,
				err:    errors.New("query database failed"),
			},
		},
		{
			name: "Subscribe exists",
			requestBody: map[string]interface{}{
				"requestor": "andy@example.com",
				"target":    "john@example.com",
			},
			expectedResponseBody: "{\"Err\":{},\"StatusCode\":208,\"StatusText\":\"Bad request\",\"Message\":\"you are subscribed the target\"}\n",
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
			mockCheckExistedSubscribe: mockCheckExistedSubscribe{
				input:  []int{1, 2},
				result: true,
				err:    nil,
			},
		},
		{
			name: "check is blocked failed",
			requestBody: map[string]interface{}{
				"requestor": "andy@example.com",
				"target":    "john@example.com",
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
			mockCheckExistedSubscribe: mockCheckExistedSubscribe{
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
				"requestor": "andy@example.com",
				"target":    "john@example.com",
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
			mockCheckExistedSubscribe: mockCheckExistedSubscribe{
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
			name: "create Subscribe failed",
			requestBody: map[string]interface{}{
				"requestor": "andy@example.com",
				"target":    "john@example.com",
			},
			expectedResponseBody: "{\"Err\":{},\"StatusCode\":500,\"StatusText\":\"Internal server error\",\"Message\":\"create Subscribe failed\"}\n",
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
			mockCheckExistedSubscribe: mockCheckExistedSubscribe{
				input:  []int{1, 2},
				result: false,
				err:    nil,
			},
			mockIsBlocked: mockIsBlocked{
				input:  []int{1, 2},
				result: false,
				err:    nil,
			},
			mockCreateSubscribe: mockCreateSubscribe{
				input: &models.Subscribe{
					Requestor: 1,
					Target:    2,
				},
				err: errors.New("create Subscribe failed"),
			},
		},
		{
			name: "create Subscribe successfully",
			requestBody: map[string]interface{}{
				"requestor": "andy@example.com",
				"target":    "john@example.com",
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
			mockCheckExistedSubscribe: mockCheckExistedSubscribe{
				input:  []int{1, 2},
				result: false,
			},
			mockIsBlocked: mockIsBlocked{
				input:  []int{1, 2},
				result: false,
			},
			mockCreateSubscribe: mockCreateSubscribe{
				input: &models.Subscribe{
					Requestor: 1,
					Target:    2,
				},
				err: nil,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			mockSubscribeService := new(service.MockSubscribe)
			mockUserService := new(service.MockUser)

			mockUserService.On("GetUserIDByEmail", testCase.mockGetRequestorId.input).
				Return(testCase.mockGetRequestorId.result, testCase.mockGetRequestorId.err)
			mockUserService.On("GetUserIDByEmail", testCase.mockGetTargetId.input).
				Return(testCase.mockGetTargetId.result, testCase.mockGetTargetId.err)

			mockSubscribeService.On("CreateSubscribe", testCase.mockCreateSubscribe.input).
				Return(testCase.mockCreateSubscribe.err)

			if testCase.mockCheckExistedSubscribe.input != nil {
				mockSubscribeService.On("CheckExistedSubscribe",
					testCase.mockCheckExistedSubscribe.input[0], testCase.mockCheckExistedSubscribe.input[1]).
					Return(testCase.mockCheckExistedSubscribe.result, testCase.mockCheckExistedSubscribe.err)
			}
			if testCase.mockIsBlocked.input != nil {
				mockSubscribeService.On("CheckBlockedByUser",
					testCase.mockIsBlocked.input[0], testCase.mockIsBlocked.input[1]).
					Return(testCase.mockIsBlocked.result, testCase.mockIsBlocked.err)
			}

			handlers := SubscribeHandlers{
				ISubscribeService: mockSubscribeService,
				IUserService:      mockUserService,
			}

			requestBody, err := json.Marshal(testCase.requestBody)
			if err != nil {
				t.Error(err)
			}

			// When
			req, err := http.NewRequest(http.MethodPost, "/subscribe", bytes.NewBuffer(requestBody))
			if err != nil {
				t.Error(err)
			}
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(handlers.CreateSubscribe)
			handler.ServeHTTP(rr, req)

			// Then
			require.Equal(t, testCase.expectedStatus, rr.Code)
			require.Equal(t, testCase.expectedResponseBody, rr.Body.String())

		})
	}
}
