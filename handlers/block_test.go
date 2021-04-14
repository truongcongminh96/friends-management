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

func TestBlockHandlers_CreateBlock(t *testing.T) {
	type mockGetUserIDFromEmail struct {
		input  string
		result int
		err    error
	}
	type mockIsExistedBlock struct {
		input  []int
		result bool
		err    error
	}
	type mockCreateBlock struct {
		input *models.Block
		err   error
	}
	testCases := []struct {
		name                 string
		requestBody          map[string]interface{}
		expectedResponseBody string
		expectedStatus       int
		mockGetRequestorID   mockGetUserIDFromEmail
		mockGetTargetID      mockGetUserIDFromEmail
		mockIsExistedBlock   mockIsExistedBlock
		mockCreateBlock      mockCreateBlock
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
				"target":    "lisa@example.com",
			},
			expectedResponseBody: "{\"Err\":{},\"StatusCode\":500,\"StatusText\":\"Bad request\",\"Message\":\"get user id failed\"}\n",
			expectedStatus:       http.StatusInternalServerError,
			mockGetRequestorID: mockGetUserIDFromEmail{
				input:  "andy@example.com",
				result: 0,
				err:    errors.New("get user id failed"),
			},
		},
		{
			name: "requestor does not exist",
			requestBody: map[string]interface{}{
				"requestor": "andy@example.com",
				"target":    "lisa@example.com",
			},
			expectedResponseBody: "{\"Err\":{},\"StatusCode\":400,\"StatusText\":\"Bad request\",\"Message\":\"the requestor does not exist\"}\n",
			expectedStatus:       http.StatusBadRequest,
			mockGetRequestorID: mockGetUserIDFromEmail{
				input:  "andy@example.com",
				result: 0,
				err:    nil,
			},
		},
		{
			name: "get requestor id failed",
			requestBody: map[string]interface{}{
				"requestor": "andy@example.com",
				"target":    "lisa@example.com",
			},
			expectedResponseBody: "{\"Err\":{},\"StatusCode\":500,\"StatusText\":\"Bad request\",\"Message\":\"get user id failed\"}\n",
			expectedStatus:       http.StatusInternalServerError,
			mockGetRequestorID: mockGetUserIDFromEmail{
				input:  "andy@example.com",
				result: 1,
				err:    nil,
			},
			mockGetTargetID: mockGetUserIDFromEmail{
				input:  "lisa@example.com",
				result: 0,
				err:    errors.New("get user id failed"),
			},
		},
		{
			name: "target does not exist",
			requestBody: map[string]interface{}{
				"requestor": "andy@example.com",
				"target":    "lisa@example.com",
			},
			expectedResponseBody: "{\"Err\":{},\"StatusCode\":400,\"StatusText\":\"Bad request\",\"Message\":\"the target does not exist\"}\n",
			expectedStatus:       http.StatusBadRequest,
			mockGetRequestorID: mockGetUserIDFromEmail{
				input:  "andy@example.com",
				result: 1,
				err:    nil,
			},
			mockGetTargetID: mockGetUserIDFromEmail{
				input:  "lisa@example.com",
				result: 0,
				err:    nil,
			},
		},
		{
			name: "check Block exist failed",
			requestBody: map[string]interface{}{
				"requestor": "andy@example.com",
				"target":    "lisa@example.com",
			},
			expectedResponseBody: "{\"Err\":{},\"StatusCode\":500,\"StatusText\":\"Bad request\",\"Message\":\"query database failed\"}\n",
			expectedStatus:       http.StatusInternalServerError,
			mockGetRequestorID: mockGetUserIDFromEmail{
				input:  "andy@example.com",
				result: 1,
				err:    nil,
			},
			mockGetTargetID: mockGetUserIDFromEmail{
				input:  "lisa@example.com",
				result: 2,
				err:    nil,
			},
			mockIsExistedBlock: mockIsExistedBlock{
				input:  []int{1, 2},
				result: false,
				err:    errors.New("query database failed"),
			},
		},
		{
			name: "Block exists",
			requestBody: map[string]interface{}{
				"requestor": "andy@example.com",
				"target":    "lisa@example.com",
			},
			expectedResponseBody: "{\"Err\":{},\"StatusCode\":208,\"StatusText\":\"Bad request\",\"Message\":\"you are block the target\"}\n",
			expectedStatus:       http.StatusAlreadyReported,
			mockGetRequestorID: mockGetUserIDFromEmail{
				input:  "andy@example.com",
				result: 1,
				err:    nil,
			},
			mockGetTargetID: mockGetUserIDFromEmail{
				input:  "lisa@example.com",
				result: 2,
				err:    nil,
			},
			mockIsExistedBlock: mockIsExistedBlock{
				input:  []int{1, 2},
				result: true,
				err:    nil,
			},
		},
		{
			name: "create Block failed",
			requestBody: map[string]interface{}{
				"requestor": "andy@example.com",
				"target":    "lisa@example.com",
			},
			expectedResponseBody: "{\"Err\":{},\"StatusCode\":500,\"StatusText\":\"Internal server error\",\"Message\":\"create Block failed\"}\n",
			expectedStatus:       http.StatusInternalServerError,
			mockGetRequestorID: mockGetUserIDFromEmail{
				input:  "andy@example.com",
				result: 1,
				err:    nil,
			},
			mockGetTargetID: mockGetUserIDFromEmail{
				input:  "lisa@example.com",
				result: 2,
				err:    nil,
			},
			mockIsExistedBlock: mockIsExistedBlock{
				input:  []int{1, 2},
				result: false,
				err:    nil,
			},
			mockCreateBlock: mockCreateBlock{
				input: &models.Block{
					Requestor: 1,
					Target:    2,
				},
				err: errors.New("create Block failed"),
			},
		},
		{
			name: "create Block successfully",
			requestBody: map[string]interface{}{
				"requestor": "andy@example.com",
				"target":    "lisa@example.com",
			},
			expectedResponseBody: "{\"success\":true}\n",
			expectedStatus:       http.StatusOK,
			mockGetRequestorID: mockGetUserIDFromEmail{
				input:  "andy@example.com",
				result: 1,
			},
			mockGetTargetID: mockGetUserIDFromEmail{
				input:  "lisa@example.com",
				result: 2,
			},
			mockIsExistedBlock: mockIsExistedBlock{
				input:  []int{1, 2},
				result: false,
			},
			mockCreateBlock: mockCreateBlock{
				input: &models.Block{
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
			mockBlockService := new(service.MockBlock)
			mockUserService := new(service.MockUser)

			mockUserService.On("GetUserIDByEmail", testCase.mockGetRequestorID.input).
				Return(testCase.mockGetRequestorID.result, testCase.mockGetRequestorID.err)
			mockUserService.On("GetUserIDByEmail", testCase.mockGetTargetID.input).
				Return(testCase.mockGetTargetID.result, testCase.mockGetTargetID.err)

			mockBlockService.On("CreateBlock", testCase.mockCreateBlock.input).
				Return(testCase.mockCreateBlock.err)

			if testCase.mockIsExistedBlock.input != nil {
				mockBlockService.On("CheckExistedBlock",
					testCase.mockIsExistedBlock.input[0], testCase.mockIsExistedBlock.input[1]).
					Return(testCase.mockIsExistedBlock.result, testCase.mockIsExistedBlock.err)
			}

			handlers := BlockHandlers{
				IBlockService: mockBlockService,
				IUserService:  mockUserService,
			}

			requestBody, err := json.Marshal(testCase.requestBody)
			if err != nil {
				t.Error(err)
			}

			// When
			req, err := http.NewRequest(http.MethodPost, "/Blocks", bytes.NewBuffer(requestBody))
			if err != nil {
				t.Error(err)
			}
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(handlers.CreateBlock)
			handler.ServeHTTP(rr, req)

			// Then
			require.Equal(t, testCase.expectedStatus, rr.Code)
			require.Equal(t, testCase.expectedResponseBody, rr.Body.String())

		})
	}
}
