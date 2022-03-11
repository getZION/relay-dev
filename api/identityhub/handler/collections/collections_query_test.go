package collections

import (
	"testing"

	"github.com/getzion/relay/api/identityhub/errors"
	hub "github.com/getzion/relay/gen/proto/identityhub/v1"
	"github.com/stretchr/testify/require"
)

func Test_CollectionQuery_ValidationFailed(t *testing.T) {
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
			name: "invalid dataFormat",
			message: &hub.Message{
				Descriptor_: &hub.MessageDescriptor{
					ObjectId:   OBJECT_ID,
					DataFormat: INVALID,
				},
			},
			expectedStatusCode:   400,
			expectedErrorMessage: errors.ImproperlyConstructedErrorMessage,
		},
		{
			name: "invalid dateSort",
			message: &hub.Message{
				Descriptor_: &hub.MessageDescriptor{
					ObjectId: OBJECT_ID,
					DateSort: INVALID,
				},
			},
			expectedStatusCode:   400,
			expectedErrorMessage: errors.ImproperlyConstructedErrorMessage,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entry, err := CollectionsQuery(nil, tt.message)

			require.Empty(t, entry)
			require.NotNil(t, err)
			require.Equal(t, tt.expectedStatusCode, err.Code)
			require.Equal(t, tt.expectedErrorMessage, errors.ImproperlyConstructedErrorMessage)
		})
	}
}
