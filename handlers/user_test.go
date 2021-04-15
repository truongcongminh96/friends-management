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

func TestUserHandlers_CreateUser(t *testing.T) {
	type mockIsExistedUser struct {
		input  string
		result bool
		err    error
	}
	type mockCreateUser struct {
		input *models.User
		err   error
	}

	testCases := []struct {
		name                 string
		requestBody          map[string]interface{}
		expectedResponseBody string
		expectedStatus       int
		mockIsExistedUser    mockIsExistedUser
		mockCreateUser       mockCreateUser
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
			name: "check email exist failed",
			requestBody: map[string]interface{}{
				"email": "andy@example",
			},
			expectedResponseBody: "{\"Err\":{},\"StatusCode\":500,\"StatusText\":\"Bad request\",\"Message\":\"query database failed\"}\n",
			expectedStatus:       http.StatusInternalServerError,
			mockIsExistedUser: mockIsExistedUser{
				input:  "andy@example",
				result: false,
				err:    errors.New("query database failed"),
			},
		},
		{
			name: "email exists",
			requestBody: map[string]interface{}{
				"email": "andy@example",
			},
			expectedResponseBody: "{\"Err\":{},\"StatusCode\":208,\"StatusText\":\"Bad request\",\"Message\":\"email address exists\"}\n",
			expectedStatus:       http.StatusAlreadyReported,
			mockIsExistedUser: mockIsExistedUser{
				input:  "andy@example",
				result: true,
				err:    nil,
			},
		},
		{
			name: "create user failed",
			requestBody: map[string]interface{}{
				"email": "andy@example",
			},
			expectedResponseBody: "{\"Err\":{},\"StatusCode\":500,\"StatusText\":\"Internal server error\",\"Message\":\"create user failed\"}\n",
			expectedStatus:       http.StatusInternalServerError,
			mockIsExistedUser: mockIsExistedUser{
				input:  "andy@example",
				result: false,
			},
			mockCreateUser: mockCreateUser{
				input: &models.User{
					Email: "andy@example",
				},
				err: errors.New("create user failed"),
			},
		},
		{
			name: "create user successfully",
			requestBody: map[string]interface{}{
				"email": "andy@example",
			},
			expectedResponseBody: "{\"success\":true}\n",
			expectedStatus:       http.StatusOK,
			mockIsExistedUser: mockIsExistedUser{
				input:  "andy@example",
				result: false,
			},
			mockCreateUser: mockCreateUser{
				input: &models.User{
					Email: "andy@example",
				},
				err: nil,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			mockUserService := new(service.MockUser)

			mockUserService.On("IsExistedUser", testCase.mockIsExistedUser.input).
				Return(testCase.mockIsExistedUser.result, testCase.mockIsExistedUser.err)
			mockUserService.On("CreateUser", testCase.mockCreateUser.input).
				Return(testCase.mockCreateUser.err)

			handlers := UserHandler{
				IUserService: mockUserService,
			}

			requestBody, err := json.Marshal(testCase.requestBody)
			if err != nil {
				t.Error(err)
			}

			// When
			req, err := http.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(requestBody))
			if err != nil {
				t.Error(err)
			}
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(handlers.CreateUser)
			handler.ServeHTTP(rr, req)

			// Then
			require.Equal(t, testCase.expectedStatus, rr.Code)
			require.Equal(t, testCase.expectedResponseBody, rr.Body.String())

		})
	}
}
