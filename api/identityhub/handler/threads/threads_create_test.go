package threads

import (
	"testing"

	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/identityhub/errors"
	"github.com/getzion/relay/api/identityhub/handler"
	"github.com/stretchr/testify/require"
)

func Test_ThreadsCreate_ValidationFailed(t *testing.T) {
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
			expectedStatusCode:   400,
			expectedErrorMessage: errors.ImproperlyConstructedErrorMessage,
		},
		{
			name: "invalid objectId",
			message: &api.Message{
				Descriptor: &api.MessageDescriptor{
					ObjectId: INVALID,
				},
			},
			expectedStatusCode:   400,
			expectedErrorMessage: errors.ImproperlyConstructedErrorMessage,
		},
		{
			name: "missing schema",
			message: &api.Message{
				Descriptor: &api.MessageDescriptor{
					ObjectId: OBJECT_ID,
				},
			},
			expectedStatusCode:   400,
			expectedErrorMessage: errors.ImproperlyConstructedErrorMessage,
		},
		{
			name: "invalid schema",
			message: &api.Message{
				Descriptor: &api.MessageDescriptor{
					ObjectId: OBJECT_ID,
					Schema:   INVALID,
				},
			},
			expectedStatusCode:   400,
			expectedErrorMessage: errors.ImproperlyConstructedErrorMessage,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entry, err := ThreadsCreate(&handler.RequestContext{Message: tt.message})

			require.Empty(t, entry)
			require.NotNil(t, err)
			require.Equal(t, tt.expectedStatusCode, err.Code)
			require.Equal(t, tt.expectedErrorMessage, errors.ImproperlyConstructedErrorMessage)
		})
	}
}
