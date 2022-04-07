package collections

import (
	"fmt"
	"testing"

	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/identityhub/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
)

func Test_CollectionCommit_ValidationFailed(t *testing.T) {
	tests := []struct {
		name                 string
		message              *api.Message
		expectedStatusCode   int
		expectedErrorMessage string
	}{
		{
			name: "missing objectId",
			message: &api.Message{
				Descriptor: &api.MessageDescriptor{},
			},
			expectedStatusCode:   fiber.StatusBadRequest,
			expectedErrorMessage: "invalid objectId: ",
		},
		{
			name: "invalid objectId",
			message: &api.Message{
				Descriptor: &api.MessageDescriptor{
					ObjectId: INVALID,
				},
			},
			expectedStatusCode:   fiber.StatusBadRequest,
			expectedErrorMessage: fmt.Sprintf("invalid objectId: %s", INVALID),
		},
		{
			name: "invalid schema",
			message: &api.Message{
				Descriptor: &api.MessageDescriptor{
					ObjectId: OBJECT_ID,
					Schema:   INVALID,
				},
			},
			expectedStatusCode:   fiber.StatusBadRequest,
			expectedErrorMessage: fmt.Sprintf("invalid schema: %s", INVALID),
		},
		{
			name: "missing dateCreated",
			message: &api.Message{
				Descriptor: &api.MessageDescriptor{
					ObjectId: OBJECT_ID,
				},
			},
			expectedStatusCode:   fiber.StatusBadRequest,
			expectedErrorMessage: fmt.Sprintf("dateCreated cannot be null or empty"),
		},
		{
			name: "invalid dateCreated",
			message: &api.Message{
				Descriptor: &api.MessageDescriptor{
					ObjectId:    OBJECT_ID,
					DateCreated: INVALID,
				},
			},
			expectedStatusCode:   fiber.StatusBadRequest,
			expectedErrorMessage: fmt.Sprintf("invalid dateCreated: %s", INVALID),
		},
		{
			name: "invalid datePublished",
			message: &api.Message{
				Descriptor: &api.MessageDescriptor{
					ObjectId:      OBJECT_ID,
					DateCreated:   DATE_CREATED,
					DatePublished: INVALID,
				},
			},
			expectedStatusCode:   fiber.StatusBadRequest,
			expectedErrorMessage: fmt.Sprintf("invalid datePublished: %s", INVALID),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entry, err := CollectionsCommit(&handler.RequestContext{Message: tt.message})

			require.Empty(t, entry)
			require.NotNil(t, err)
			require.Equal(t, tt.expectedStatusCode, err.Code)
			require.Equal(t, tt.expectedErrorMessage, err.Message)
		})
	}
}
