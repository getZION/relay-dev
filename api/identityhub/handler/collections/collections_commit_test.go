package collections

import (
	"fmt"
	"testing"

	"github.com/getzion/relay/api/identityhub/handler"
	hub "github.com/getzion/relay/gen/proto/identityhub/v1"
	"github.com/stretchr/testify/require"
)

func Test_CollectionCommit_ValidationFailed(t *testing.T) {
	tests := []struct {
		name                 string
		message              *hub.Message
		expectedStatusCode   int64
		expectedErrorMessage string
	}{
		{
			name: "missing objectId",
			message: &hub.Message{
				Descriptor_: &hub.MessageDescriptor{},
			},
			expectedStatusCode:   400,
			expectedErrorMessage: "invalid objectId: ",
		},
		{
			name: "invalid objectId",
			message: &hub.Message{
				Descriptor_: &hub.MessageDescriptor{
					ObjectId: INVALID,
				},
			},
			expectedStatusCode:   400,
			expectedErrorMessage: fmt.Sprintf("invalid objectId: %s", INVALID),
		},
		{
			name: "invalid schema",
			message: &hub.Message{
				Descriptor_: &hub.MessageDescriptor{
					ObjectId: OBJECT_ID,
					Schema:   INVALID,
				},
			},
			expectedStatusCode:   400,
			expectedErrorMessage: fmt.Sprintf("invalid schema: %s", INVALID),
		},
		{
			name: "missing dateCreated",
			message: &hub.Message{
				Descriptor_: &hub.MessageDescriptor{
					ObjectId: OBJECT_ID,
				},
			},
			expectedStatusCode:   400,
			expectedErrorMessage: fmt.Sprintf("dateCreated cannot be null or empty"),
		},
		{
			name: "invalid dateCreated",
			message: &hub.Message{
				Descriptor_: &hub.MessageDescriptor{
					ObjectId:    OBJECT_ID,
					DateCreated: INVALID,
				},
			},
			expectedStatusCode:   400,
			expectedErrorMessage: fmt.Sprintf("invalid dateCreated: %s", INVALID),
		},
		{
			name: "invalid datePublished",
			message: &hub.Message{
				Descriptor_: &hub.MessageDescriptor{
					ObjectId:      OBJECT_ID,
					DateCreated:   DATE_CREATED,
					DatePublished: INVALID,
				},
			},
			expectedStatusCode:   400,
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
