package identityhub

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/schema"
	"github.com/getzion/relay/api/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	REQUEST_ID     = "3eb8ea70-7ea5-4069-a153-cfb0ea682df9"
	TARGET         = "TheTarget"
	OBJECT_ID      = "d82c0026-ed42-4b26-81f3-94805958a75c"
	DATE_CREATED   = "1645917431"
	DATE_PUBLISHED = "1645917431"
	SCHEMA         = "https://test.com"
	DATA_FORMAT    = "application/json"
	ROOT           = "e23ea8cf-5e64-42d0-b3c6-54e5ab1dcf25"
	PARENT         = "654a4593-4c01-4a6c-9cd9-6bf04bd3d441"
	INVALID        = "<invalid>"
)

func prepareApp(t *testing.T) *fiber.App {
	gomockController := gomock.NewController(t)
	st := storage.NewMockStorage(gomockController)
	schemaManager := schema.NewSchemaManager(st)
	server := InitIdentityHubServer(schemaManager, st)
	return server.app
}

func process(app *fiber.App, body *api.Request) (*api.Response, error) {

	data, _ := json.Marshal(body)
	reader := bytes.NewReader(data)

	req := httptest.NewRequest("POST", "/process", reader)
	req.Header.Add("Content-Type", "application/json")
	resp, err := app.Test(req, -1)
	if err != nil {
		return nil, err
	}

	responseByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response api.Response
	err = json.Unmarshal(responseByte, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func Test_RequestLevelResponses(t *testing.T) {
	app := prepareApp(t)

	tests := []struct {
		description       string
		body              *api.Request
		expectedCode      int
		expectedRequestId string
	}{
		{
			description:       "receives an error if Request is missing",
			body:              nil,
			expectedCode:      fiber.StatusBadRequest,
			expectedRequestId: "",
		},
		{
			description: "receives an error if RequestID is missing",
			body: &api.Request{
				Target:   TARGET,
				Messages: make([]*api.Message, 1),
			},
			expectedCode:      fiber.StatusBadRequest,
			expectedRequestId: "",
		},
		{
			description: "receives an error if RequestID is not len 36 (uuid v4)",
			body: &api.Request{
				Target:    TARGET,
				RequestId: INVALID,
				Messages:  make([]*api.Message, 1),
			},
			expectedCode:      fiber.StatusBadRequest,
			expectedRequestId: INVALID,
		},
		{
			description: "receives an error if Target is missing",
			body: &api.Request{
				RequestId: REQUEST_ID,
				Messages:  make([]*api.Message, 1),
			},
			expectedCode:      fiber.StatusBadRequest,
			expectedRequestId: REQUEST_ID,
		},
		{
			description: "receives an error if Messages are missing",
			body: &api.Request{
				RequestId: REQUEST_ID,
				Target:    TARGET,
			},
			expectedCode:      fiber.StatusBadRequest,
			expectedRequestId: REQUEST_ID,
		},
	}

	for _, test := range tests {

		var reader io.Reader
		if test.body != nil {
			data, _ := json.Marshal(test.body)
			reader = bytes.NewReader(data)
		}
		req := httptest.NewRequest("POST", "/process", reader)
		req.Header.Add("Content-Type", "application/json")
		resp, err := app.Test(req, -1)

		require.Nil(t, err)
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)

		body, err := ioutil.ReadAll(resp.Body)
		require.Nil(t, err)

		var response api.Response
		err = json.Unmarshal(body, &response)
		require.Nil(t, err)

		require.NotNil(t, response, "response cannot be nil")
		require.Equal(t, test.expectedRequestId, response.RequestId, "response requestId must be expected value")
		require.NotNil(t, response.Status, "response status cannot be nil")
		require.Equal(t, test.expectedCode, response.Status.Code, "response status must be expected value")
	}
}

func Test_MessageLevelResponses(t *testing.T) {
	app := prepareApp(t)

	tests := []struct {
		description  string
		body         *api.Request
		expectedCode int
		expectation  func(*testing.T, *api.Response)
	}{
		{
			description: "receives an error if a Message is missing Descriptor",
			body: &api.Request{
				RequestId: REQUEST_ID,
				Target:    TARGET,
				Messages: []*api.Message{
					{},
				},
			},
			expectedCode: fiber.StatusOK,
			expectation: func(t *testing.T, r *api.Response) {
				require.NotNil(t, r.Replies, "replies cannot be nil")
				require.Len(t, r.Replies, 1, "replies must be contains 1 element")
				require.NotNil(t, r.Replies[0].Status, "reply status cannot be nil")
				require.Equal(t, fiber.StatusBadRequest, r.Replies[0].Status.Code, "reply status code must be expected value")
			},
		},
		{
			description: "receives an error if a Message Descriptor is missing Method",
			body: &api.Request{
				RequestId: REQUEST_ID,
				Target:    TARGET,
				Messages: []*api.Message{
					{
						Descriptor: &api.MessageDescriptor{},
					},
				},
			},
			expectedCode: fiber.StatusOK,
			expectation: func(t *testing.T, r *api.Response) {
				require.NotNil(t, r.Replies, "replies cannot be nil")
				require.Len(t, r.Replies, 1, "replies must be contains 1 element")
				require.NotNil(t, r.Replies[0].Status, "reply status cannot be nil")
				require.Equal(t, fiber.StatusBadRequest, r.Replies[0].Status.Code, "reply status code must be expected value")
			},
		},
		{
			description: "receives an error if a Message Descriptor method is not implemented",
			body: &api.Request{
				RequestId: REQUEST_ID,
				Target:    TARGET,
				Messages: []*api.Message{
					{
						Descriptor: &api.MessageDescriptor{
							Method: INVALID,
						},
					},
				},
			},
			expectedCode: fiber.StatusOK,
			expectation: func(t *testing.T, r *api.Response) {
				require.NotNil(t, r.Replies, "replies cannot be nil")
				require.Len(t, r.Replies, 1, "replies must be contains 1 element")
				require.NotNil(t, r.Replies[0].Status, "reply status cannot be nil")
				require.Equal(t, 501, r.Replies[0].Status.Code, "reply status code must be expected value")
			},
		},
	}

	for _, test := range tests {

		var reader io.Reader
		if test.body != nil {
			data, _ := json.Marshal(test.body)
			reader = bytes.NewReader(data)
		}
		req := httptest.NewRequest("POST", "/process", reader)
		req.Header.Add("Content-Type", "application/json")
		resp, err := app.Test(req, -1)
		require.Nil(t, err)
		require.Equal(t, fiber.StatusOK, resp.StatusCode, "response status code must be 200")

		body, err := ioutil.ReadAll(resp.Body)
		require.Nil(t, err)

		var response api.Response
		err = json.Unmarshal(body, &response)
		require.Nil(t, err)

		require.NotNil(t, response, "response cannot be nil")
		require.NotNil(t, response.Status, "response status cannot be nil")
		require.Equal(t, fiber.StatusOK, response.Status.Code, "response status must be 200")

		test.expectation(t, &response)
	}
}
