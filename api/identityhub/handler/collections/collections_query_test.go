package collections

import (
	"fmt"
	"testing"

	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/constants"
	"github.com/getzion/relay/api/identityhub/handler"
	"github.com/getzion/relay/api/schema"
	"github.com/getzion/relay/api/storage"
	hub "github.com/getzion/relay/gen/proto/identityhub/v1"
	"github.com/golang/mock/gomock"
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
			name: "unknown schema",
			message: &hub.Message{
				Descriptor_: &hub.MessageDescriptor{
					ObjectId: OBJECT_ID,
					Schema:   SCHEMA_UNKNOWN,
				},
			},
			expectedStatusCode:   400,
			expectedErrorMessage: fmt.Sprintf("unknown schema: %s", SCHEMA_UNKNOWN),
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
			expectedErrorMessage: fmt.Sprintf("invalid dataFormat: %s", INVALID),
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
			expectedErrorMessage: fmt.Sprintf("invalid dateSort: %s", INVALID),
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := storage.NewMockStorage(ctrl)
	schemaManager := schema.NewSchemaManager(storage)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entry, err := CollectionsQuery(&handler.RequestContext{Message: tt.message, SchemaManager: schemaManager})

			require.Empty(t, entry)
			require.NotNil(t, err)
			require.Equal(t, tt.expectedStatusCode, err.Code)
			require.Equal(t, tt.expectedErrorMessage, err.Message)
		})
	}
}

func Test_Communities_Get(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := storage.NewMockStorage(ctrl)
	storage.EXPECT().GetCommunities().Times(1).Return([]api.Community{
		{Id: 1, Zid: "zid1", OwnerDid: "did1", OwnerUsername: "user1"},
		{Id: 2, Zid: "zid2", OwnerDid: "did1", OwnerUsername: "user1"},
	}, nil)

	schemaManager := schema.NewSchemaManager(storage)

	tests := []struct {
		name                 string
		message              *hub.Message
		expectedStatusCode   int64
		expectedErrorMessage string
	}{
		{
			name: "should return 2 community",
			message: &hub.Message{
				Descriptor_: &hub.MessageDescriptor{
					ObjectId: OBJECT_ID,
					Schema:   constants.SCHEMA_COMMUNITY,
					Method:   constants.COLLECTIONS_QUERY,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entries, err := CollectionsQuery(&handler.RequestContext{SchemaManager: schemaManager, Message: tt.message})

			require.Nil(t, err)
			require.NotNil(t, entries)
			require.NotEmpty(t, entries)
			require.Len(t, entries, 2)
		})
	}
}

func Test_Conversation_Get(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := storage.NewMockStorage(ctrl)
	storage.EXPECT().GetConversations().Times(1).Return([]api.Conversation{
		{Id: 1, Zid: "zid1", CommunityZid: "c_zid1"},
		{Id: 2, Zid: "zid2", CommunityZid: "c_zid1"},
	}, nil)

	schemaManager := schema.NewSchemaManager(storage)

	tests := []struct {
		name                 string
		message              *hub.Message
		expectedStatusCode   int64
		expectedErrorMessage string
	}{
		{
			name: "should return 2 conversation",
			message: &hub.Message{
				Descriptor_: &hub.MessageDescriptor{
					ObjectId: OBJECT_ID,
					Schema:   constants.SCHEMA_CONVERSATION,
					Method:   constants.COLLECTIONS_QUERY,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entries, err := CollectionsQuery(&handler.RequestContext{SchemaManager: schemaManager, Message: tt.message})

			require.Nil(t, err)
			require.NotNil(t, entries)
			require.NotEmpty(t, entries)
			require.Len(t, entries, 2)
		})
	}
}

func Test_User_Get(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := storage.NewMockStorage(ctrl)
	storage.EXPECT().GetUsers().Times(1).Return([]api.User{
		{Id: 1, Did: "did1"},
		{Id: 2, Did: "did2"},
	}, nil)
	schemaManager := schema.NewSchemaManager(storage)

	tests := []struct {
		name                 string
		message              *hub.Message
		expectedStatusCode   int64
		expectedErrorMessage string
	}{
		{
			name: "should return 2 user",
			message: &hub.Message{
				Descriptor_: &hub.MessageDescriptor{
					ObjectId: OBJECT_ID,
					Schema:   constants.SCHEMA_PERSON,
					Method:   constants.COLLECTIONS_QUERY,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entries, err := CollectionsQuery(&handler.RequestContext{SchemaManager: schemaManager, Message: tt.message})

			require.Nil(t, err)
			require.NotNil(t, entries)
			require.NotEmpty(t, entries)
			require.Len(t, entries, 2)
		})
	}
}

func Test_Payment_Get(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := storage.NewMockStorage(ctrl)
	storage.EXPECT().GetPayments().Times(1).Return([]api.Payment{
		{Id: 1, Zid: "zid1"},
		{Id: 2, Zid: "zid2"},
	}, nil)
	schemaManager := schema.NewSchemaManager(storage)

	tests := []struct {
		name                 string
		message              *hub.Message
		expectedStatusCode   int64
		expectedErrorMessage string
	}{
		{
			name: "should return 2 payment",
			message: &hub.Message{
				Descriptor_: &hub.MessageDescriptor{
					ObjectId: OBJECT_ID,
					Schema:   constants.SCHEMA_PAYMENT,
					Method:   constants.COLLECTIONS_QUERY,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entries, err := CollectionsQuery(&handler.RequestContext{SchemaManager: schemaManager, Message: tt.message})

			require.Nil(t, err)
			require.NotNil(t, entries)
			require.NotEmpty(t, entries)
			require.Len(t, entries, 2)
		})
	}
}
