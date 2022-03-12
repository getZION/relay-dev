package threads

import (
	"testing"

	"github.com/getzion/relay/api/identityhub/errors"
	"github.com/getzion/relay/api/identityhub/handler"
	hub "github.com/getzion/relay/gen/proto/identityhub/v1"
	"github.com/stretchr/testify/require"
)

func Test_ThreadsReply_ValidationFailed(t *testing.T) {
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
			expectedErrorMessage: errors.ImproperlyConstructedErrorMessage,
		},
		{
			name: "invalid objectId",
			message: &hub.Message{
				Descriptor_: &hub.MessageDescriptor{
					ObjectId: INVALID,
				},
			},
			expectedStatusCode:   400,
			expectedErrorMessage: errors.ImproperlyConstructedErrorMessage,
		},
		{
			name: "missing schema",
			message: &hub.Message{
				Descriptor_: &hub.MessageDescriptor{
					ObjectId: OBJECT_ID,
				},
			},
			expectedStatusCode:   400,
			expectedErrorMessage: errors.ImproperlyConstructedErrorMessage,
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
			expectedErrorMessage: errors.ImproperlyConstructedErrorMessage,
		},
		{
			name: "missing root",
			message: &hub.Message{
				Descriptor_: &hub.MessageDescriptor{
					ObjectId: OBJECT_ID,
					Schema:   SCHEMA,
				},
			},
			expectedStatusCode:   400,
			expectedErrorMessage: errors.ImproperlyConstructedErrorMessage,
		},
		{
			name: "invalid root",
			message: &hub.Message{
				Descriptor_: &hub.MessageDescriptor{
					ObjectId: OBJECT_ID,
					Schema:   SCHEMA,
					Root:     INVALID,
				},
			},
			expectedStatusCode:   400,
			expectedErrorMessage: errors.ImproperlyConstructedErrorMessage,
		},
		{
			name: "missing parent",
			message: &hub.Message{
				Descriptor_: &hub.MessageDescriptor{
					ObjectId: OBJECT_ID,
					Schema:   SCHEMA,
					Root:     ROOT,
				},
			},
			expectedStatusCode:   400,
			expectedErrorMessage: errors.ImproperlyConstructedErrorMessage,
		},
		{
			name: "invalid parent",
			message: &hub.Message{
				Descriptor_: &hub.MessageDescriptor{
					ObjectId: OBJECT_ID,
					Schema:   SCHEMA,
					Root:     ROOT,
					Parent:   INVALID,
				},
			},
			expectedStatusCode:   400,
			expectedErrorMessage: errors.ImproperlyConstructedErrorMessage,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entry, err := ThreadsReply(&handler.RequestContext{Message: tt.message})

			require.Empty(t, entry)
			require.NotNil(t, err)
			require.Equal(t, tt.expectedStatusCode, err.Code)
			require.Equal(t, tt.expectedErrorMessage, errors.ImproperlyConstructedErrorMessage)
		})
	}
}
