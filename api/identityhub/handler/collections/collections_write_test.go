package collections

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/getzion/relay/api/constants"
	"github.com/getzion/relay/api/datastore"
	"github.com/getzion/relay/api/identityhub/errors"
	"github.com/getzion/relay/api/identityhub/handler"
	"github.com/getzion/relay/api/schema"
	hub "github.com/getzion/relay/gen/proto/identityhub/v1"
	"github.com/stretchr/testify/require"
)

func Test_CollectionWrite_ValidationFailed(t *testing.T) {
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
			name: "missing dateCreated",
			message: &hub.Message{
				Descriptor_: &hub.MessageDescriptor{
					ObjectId: OBJECT_ID,
				},
			},
			expectedStatusCode:   400,
			expectedErrorMessage: errors.ImproperlyConstructedErrorMessage,
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
			expectedErrorMessage: errors.ImproperlyConstructedErrorMessage,
		},
		{
			name: "invalid schema",
			message: &hub.Message{
				Descriptor_: &hub.MessageDescriptor{
					ObjectId:    OBJECT_ID,
					DateCreated: DATE_CREATED,
					Schema:      INVALID,
				},
			},
			expectedStatusCode:   400,
			expectedErrorMessage: errors.ImproperlyConstructedErrorMessage,
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
			expectedErrorMessage: errors.ImproperlyConstructedErrorMessage,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entry, err := CollectionsWrite(&handler.RequestContext{Message: tt.message})

			require.Empty(t, entry)
			require.NotNil(t, err)
			require.Equal(t, tt.expectedStatusCode, err.Code)
			require.Equal(t, tt.expectedErrorMessage, errors.ImproperlyConstructedErrorMessage)
		})
	}
}

func Test_CommunityCreate(t *testing.T) {
	store, mock := datastore.NewTestStore()
	schemaManager := schema.NewSchemaManager(store)

	mock.ExpectQuery("SELECT count(.*) FROM `communities`[a-zA-Z *]*").
		WillReturnRows(sqlmock.NewRows([]string{"Count"}).
			AddRow(0))

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `communities`[a-zA-Z *]*").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	entries, err := CollectionsWrite(&handler.RequestContext{
		SchemaManager: schemaManager,
		Message: &hub.Message{
			Data: `{ "Name": "test", "Description": "test", "OwnerUsername": "test_username", "OwnerDid": "test_did", "EscrowAmount": 10, "OwnerAlias": "test", "OwnerPubkey": "test", "PricePerMessage": 10, "PriceToJoin": 10 }`,
			Descriptor_: &hub.MessageDescriptor{
				ObjectId:    OBJECT_ID,
				Schema:      constants.SCHEMA_ORGANIZATION,
				DateCreated: DATE_CREATED,
				Method:      constants.COLLECTIONS_WRITE,
			},
		},
	})

	require.Nil(t, err)
	require.Len(t, entries, 1)
	require.Nil(t, mock.ExpectationsWereMet())
}

func Test_CommunityCreate_AlreadyExist(t *testing.T) {
	store, mock := datastore.NewTestStore()
	schemaManager := schema.NewSchemaManager(store)

	mock.ExpectQuery("SELECT count(.*) FROM `communities`[a-zA-Z *]*").
		WillReturnRows(sqlmock.NewRows([]string{"Count"}).
			AddRow(1))

	entries, err := CollectionsWrite(&handler.RequestContext{
		SchemaManager: schemaManager,
		Message: &hub.Message{
			Data: `{ "Name": "test", "Description": "test", "OwnerUsername": "test_username", "OwnerDid": "test_did", "EscrowAmount": 10, "OwnerAlias": "test", "OwnerPubkey": "test", "PricePerMessage": 10, "PriceToJoin": 10 }`,
			Descriptor_: &hub.MessageDescriptor{
				ObjectId:    OBJECT_ID,
				Schema:      constants.SCHEMA_ORGANIZATION,
				DateCreated: DATE_CREATED,
				Method:      constants.COLLECTIONS_WRITE,
			},
		},
	})

	require.NotNil(t, err)
	require.Equal(t, "the specified community already exist: test", err.Message)
	require.Equal(t, int64(400), err.Code)
	require.Len(t, entries, 0)
	require.Nil(t, mock.ExpectationsWereMet())
}

func Test_ConversationCreate(t *testing.T) {
	store, mock := datastore.NewTestStore()
	schemaManager := schema.NewSchemaManager(store)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `conversations`[a-zA-Z *]*").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	entries, err := CollectionsWrite(&handler.RequestContext{
		SchemaManager: schemaManager,
		Message: &hub.Message{
			Data: `{ "CommunityZid": "test_zid" }`,
			Descriptor_: &hub.MessageDescriptor{
				ObjectId:    OBJECT_ID,
				Schema:      constants.SCHEMA_CONVERSATION,
				DateCreated: DATE_CREATED,
				Method:      constants.COLLECTIONS_WRITE,
			},
		},
	})

	require.Nil(t, err)
	require.Len(t, entries, 1)
	require.Nil(t, mock.ExpectationsWereMet())
}

func Test_UserCreate(t *testing.T) {
	store, mock := datastore.NewTestStore()
	schemaManager := schema.NewSchemaManager(store)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `users`[a-zA-Z *]*").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	entries, err := CollectionsWrite(&handler.RequestContext{
		SchemaManager: schemaManager,
		Message: &hub.Message{
			Data: `{ "Name": "test_name", "Username": "test_username", "Email": "test@test.org" }`,
			Descriptor_: &hub.MessageDescriptor{
				ObjectId:    OBJECT_ID,
				Schema:      constants.SCHEMA_PERSON,
				DateCreated: DATE_CREATED,
				Method:      constants.COLLECTIONS_WRITE,
			},
		},
	})

	require.Nil(t, err)
	require.Len(t, entries, 1)
	require.Nil(t, mock.ExpectationsWereMet())
}
