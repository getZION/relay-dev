package collections

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/getzion/relay/api/identityhub/errors"
	"github.com/getzion/relay/api/identityhub/handler"
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
			entry, err := CollectionsQuery(&handler.RequestContext{Message: tt.message})

			require.Empty(t, entry)
			require.NotNil(t, err)
			require.Equal(t, tt.expectedStatusCode, err.Code)
			require.Equal(t, tt.expectedErrorMessage, errors.ImproperlyConstructedErrorMessage)
		})
	}
}

func Test_Communities_Get(t *testing.T) {

	store, mock := initTestDataStore()
	mock.ExpectQuery("SELECT[a-zA-Z *]*").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description", "escrowAmount", "owner_alias", "owner_pubkey", "price_per_message", "price_to_join"}).
			AddRow(1, "test", "desc", 0, "alias", "pubkey", 10, 10).
			AddRow(2, "test2", "desc2", 0, "alias2", "pubkey2", 20, 20))

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
					Schema:   SCHEMA_ORGANIZATION,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entries, err := CollectionsQuery(&handler.RequestContext{Store: store, Message: tt.message})

			require.Nil(t, err)
			require.NotNil(t, entries)
			require.NotEmpty(t, entries)
			require.Len(t, entries, 2)

			//todo: check datas
		})
	}

}

func Test_Conversation_Get(t *testing.T) {

	store, mock := initTestDataStore()
	mock.ExpectQuery("SELECT[a-zA-Z *]*").
		WillReturnRows(sqlmock.NewRows([]string{"id", "community_zid", "public", "public_price", "zid"}).
			AddRow(1, "c_zid1", false, 10, "zid1").
			AddRow(2, "c_zid2", true, 20, "zid2"))

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
					Schema:   SCHEMA_CONVERSATION,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entries, err := CollectionsQuery(&handler.RequestContext{Store: store, Message: tt.message})

			require.Nil(t, err)
			require.NotNil(t, entries)
			require.NotEmpty(t, entries)
			require.Len(t, entries, 2)

			//todo: check datas
		})
	}

}

func Test_User_Get(t *testing.T) {

	store, mock := initTestDataStore()
	mock.ExpectQuery("SELECT[a-zA-Z *]*").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "test_user1").
			AddRow(2, "test_user2"))

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
					Schema:   SCHEMA_PERSON,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entries, err := CollectionsQuery(&handler.RequestContext{Store: store, Message: tt.message})

			require.Nil(t, err)
			require.NotNil(t, entries)
			require.NotEmpty(t, entries)
			require.Len(t, entries, 2)

			//todo: check datas
		})
	}

}

func Test_Payment_Get(t *testing.T) {

	store, mock := initTestDataStore()
	mock.ExpectQuery("SELECT[a-zA-Z *]*").
		WillReturnRows(sqlmock.NewRows([]string{"id", "status"}).
			AddRow(1, "sended").
			AddRow(2, "sending"))

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
					Schema:   SCHEMA_PAYMENT,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entries, err := CollectionsQuery(&handler.RequestContext{Store: store, Message: tt.message})

			require.Nil(t, err)
			require.NotNil(t, entries)
			require.NotEmpty(t, entries)
			require.Len(t, entries, 2)

			//todo: check datas
		})
	}

}
