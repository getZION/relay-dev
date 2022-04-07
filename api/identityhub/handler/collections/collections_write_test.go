package collections

import (
	"errors"
	"fmt"
	"testing"

	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/constants"
	"github.com/getzion/relay/api/identityhub/handler"
	"github.com/getzion/relay/api/schema"
	"github.com/getzion/relay/api/storage"
	"github.com/getzion/relay/api/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func Test_CollectionWrite_ValidationFailed(t *testing.T) {
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
			name: "invalid schema",
			message: &api.Message{
				Descriptor: &api.MessageDescriptor{
					ObjectId:    OBJECT_ID,
					DateCreated: DATE_CREATED,
					Schema:      INVALID,
				},
			},
			expectedStatusCode:   fiber.StatusBadRequest,
			expectedErrorMessage: fmt.Sprintf("invalid schema: %s", INVALID),
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
		{
			name: "data cannot be empty",
			message: &api.Message{
				Data: "",
				Descriptor: &api.MessageDescriptor{
					ObjectId:      OBJECT_ID,
					DateCreated:   DATE_CREATED,
					DatePublished: DATE_PUBLISHED,
				},
			},
			expectedStatusCode:   fiber.StatusBadRequest,
			expectedErrorMessage: "data cannot be empty",
		},
		{
			name: "unknown schema",
			message: &api.Message{
				Data: `{ "Name": "test", "Description": "test", "OwnerUsername": "test_username", "OwnerDid": "test_did", "EscrowAmount": 10, "OwnerAlias": "test", "OwnerPubkey": "test", "PricePerMessage": 10, "PriceToJoin": 10 }`,
				Descriptor: &api.MessageDescriptor{
					ObjectId:    OBJECT_ID,
					DateCreated: DATE_CREATED,
					Schema:      SCHEMA_UNKNOWN,
				},
			},
			expectedStatusCode:   fiber.StatusBadRequest,
			expectedErrorMessage: fmt.Sprintf("unknown schema: %s", SCHEMA_UNKNOWN),
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := storage.NewMockStorage(ctrl)
	schemaManager := schema.NewSchemaManager(storage)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entry, err := CollectionsWrite(&handler.RequestContext{Message: tt.message, SchemaManager: schemaManager})

			require.Empty(t, entry)
			require.NotNil(t, err)
			require.Equal(t, tt.expectedStatusCode, err.Code)
			require.Equal(t, tt.expectedErrorMessage, err.Message)
		})
	}
}

func Test_CommunityCreate(t *testing.T) {
	validator.InitValidator()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := storage.NewMockStorage(ctrl)
	schemaManager := schema.NewSchemaManager(storage)

	storage.EXPECT().InsertCommunity(gomock.Any()).Times(1).Return(nil)

	entries, err := CollectionsWrite(&handler.RequestContext{
		SchemaManager: schemaManager,
		Message: &api.Message{
			Data: `{ "Name": "test", "Description": "test", "OwnerUsername": "test_username", "OwnerDid": "test_did", "EscrowAmount": 10, "OwnerAlias": "test", "OwnerPubkey": "test", "PricePerMessage": 10, "PriceToJoin": 10 }`,
			Descriptor: &api.MessageDescriptor{
				ObjectId:    OBJECT_ID,
				Schema:      constants.SCHEMA_COMMUNITY,
				DateCreated: DATE_CREATED,
				Method:      constants.COLLECTIONS_WRITE,
			},
		},
	})

	require.Nil(t, err)
	require.Len(t, entries, 1)
}

func Test_CommunityCreate_AlreadyExist(t *testing.T) {
	validator.InitValidator()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := storage.NewMockStorage(ctrl)
	schemaManager := schema.NewSchemaManager(storage)

	storage.EXPECT().InsertCommunity(gomock.Any()).Times(1).Return(errors.New("the specified community already exist: test"))

	entries, err := CollectionsWrite(&handler.RequestContext{
		SchemaManager: schemaManager,
		Message: &api.Message{
			Data: `{ "Name": "test", "Description": "test", "OwnerUsername": "test_username", "OwnerDid": "test_did", "EscrowAmount": 10, "OwnerAlias": "test", "OwnerPubkey": "test", "PricePerMessage": 10, "PriceToJoin": 10 }`,
			Descriptor: &api.MessageDescriptor{
				ObjectId:    OBJECT_ID,
				Schema:      constants.SCHEMA_COMMUNITY,
				DateCreated: DATE_CREATED,
				Method:      constants.COLLECTIONS_WRITE,
			},
		},
	})

	require.NotNil(t, err)
	require.Equal(t, "the specified community already exist: test", err.Message)
	require.Equal(t, fiber.StatusBadRequest, err.Code)
	require.Len(t, entries, 0)
}

func Test_ConversationCreate(t *testing.T) {
	validator.InitValidator()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := storage.NewMockStorage(ctrl)
	schemaManager := schema.NewSchemaManager(storage)

	storage.EXPECT().InsertConversation(gomock.Any()).Times(1).Return(nil)

	entries, err := CollectionsWrite(&handler.RequestContext{
		SchemaManager: schemaManager,
		Message: &api.Message{
			Data: `{ "CommunityZid": "test_zid", "UserDid": "test_did", "Text": "test" }`,
			Descriptor: &api.MessageDescriptor{
				ObjectId:    OBJECT_ID,
				Schema:      constants.SCHEMA_CONVERSATION,
				DateCreated: DATE_CREATED,
				Method:      constants.COLLECTIONS_WRITE,
			},
		},
	})

	require.Nil(t, err)
	require.Len(t, entries, 1)
}

func Test_UserCreate(t *testing.T) {
	validator.InitValidator()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := storage.NewMockStorage(ctrl)
	schemaManager := schema.NewSchemaManager(storage)

	storage.EXPECT().InsertUser(gomock.Any()).Times(1).Return(nil)

	entries, err := CollectionsWrite(&handler.RequestContext{
		SchemaManager: schemaManager,
		Message: &api.Message{
			Data: `{ "Name": "test_name", "Username": "test_username", "Email": "test@test.org", "Did": "did" }`,
			Descriptor: &api.MessageDescriptor{
				ObjectId:    OBJECT_ID,
				Schema:      constants.SCHEMA_PERSON,
				DateCreated: DATE_CREATED,
				Method:      constants.COLLECTIONS_WRITE,
			},
		},
	})

	require.Nil(t, err)
	require.Len(t, entries, 1)
}

func Test_UserCreate_AlreadyExist(t *testing.T) {
	validator.InitValidator()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := storage.NewMockStorage(ctrl)
	schemaManager := schema.NewSchemaManager(storage)

	storage.EXPECT().InsertUser(gomock.Any()).Times(1).Return(errors.New("the specified username already exist: test_username"))

	entries, err := CollectionsWrite(&handler.RequestContext{
		SchemaManager: schemaManager,
		Message: &api.Message{
			Data: `{ "Name": "test_name", "Username": "test_username", "Email": "test@test.org", "Did": "did" }`,
			Descriptor: &api.MessageDescriptor{
				ObjectId:    OBJECT_ID,
				Schema:      constants.SCHEMA_PERSON,
				DateCreated: DATE_CREATED,
				Method:      constants.COLLECTIONS_WRITE,
			},
		},
	})

	require.NotNil(t, err)
	require.Equal(t, "the specified username already exist: test_username", err.Message)
	require.Equal(t, fiber.StatusBadRequest, err.Code)
	require.Len(t, entries, 0)
}

func Test_JoinCommunity(t *testing.T) {
	validator.InitValidator()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := storage.NewMockStorage(ctrl)
	schemaManager := schema.NewSchemaManager(storage)

	community := &api.Community{Zid: "zid"}
	user := &api.User{Did: "did"}
	storage.EXPECT().GetCommunityByZid("zid").Times(1).Return(community, nil)
	storage.EXPECT().GetUserByDid("did").Times(1).Return(user, nil)
	storage.EXPECT().AddUserToCommunity(community, user).Times(1).Return(nil)

	entries, err := CollectionsWrite(&handler.RequestContext{
		SchemaManager: schemaManager,
		Message: &api.Message{
			Data: `{ "communityZid": "zid", "userDid": "did" }`,
			Descriptor: &api.MessageDescriptor{
				ObjectId:    OBJECT_ID,
				Schema:      constants.SCHEMA_JOIN_COMMUNITY,
				DateCreated: DATE_CREATED,
				Method:      constants.COLLECTIONS_WRITE,
			},
		},
	})

	require.Nil(t, err)
	require.Nil(t, entries)
}

func Test_LeaveCommunity(t *testing.T) {
	validator.InitValidator()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := storage.NewMockStorage(ctrl)
	schemaManager := schema.NewSchemaManager(storage)

	storage.EXPECT().GetCommunityByZid("zid").Times(1).Return(&api.Community{Zid: "zid"}, nil)
	storage.EXPECT().GetUserByDid("did").Times(1).Return(&api.User{Did: "did"}, nil)
	storage.EXPECT().RemoveUserToCommunity("zid", "did", "").Times(1).Return(nil)

	entries, err := CollectionsWrite(&handler.RequestContext{
		SchemaManager: schemaManager,
		Message: &api.Message{
			Data: `{ "communityZid": "zid", "userDid": "did" }`,
			Descriptor: &api.MessageDescriptor{
				ObjectId:    OBJECT_ID,
				Schema:      constants.SCHEMA_LEAVE_COMMUNITY,
				DateCreated: DATE_CREATED,
				Method:      constants.COLLECTIONS_WRITE,
			},
		},
	})

	require.Nil(t, err)
	require.Nil(t, entries)
}

func Test_KickUserCommunity(t *testing.T) {
	validator.InitValidator()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := storage.NewMockStorage(ctrl)
	schemaManager := schema.NewSchemaManager(storage)

	storage.EXPECT().GetCommunityByZid("zid").Times(1).Return(&api.Community{Zid: "zid"}, nil)
	storage.EXPECT().GetUserByDid("did").Times(1).Return(&api.User{Did: "did"}, nil)
	storage.EXPECT().RemoveUserToCommunity("zid", "did", "Kicked by Owner").Times(1).Return(nil)

	entries, err := CollectionsWrite(&handler.RequestContext{
		SchemaManager: schemaManager,
		Message: &api.Message{
			Data: `{ "communityZid": "zid", "userDid": "did" }`,
			Descriptor: &api.MessageDescriptor{
				ObjectId:    OBJECT_ID,
				Schema:      constants.SCHEMA_KICK_USER,
				DateCreated: DATE_CREATED,
				Method:      constants.COLLECTIONS_WRITE,
			},
		},
	})

	require.Nil(t, err)
	require.Nil(t, entries)
}

// todo: user already joined test
