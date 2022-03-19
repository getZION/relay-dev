package schema

import (
	"testing"

	"github.com/getzion/relay/api/constants"
	"github.com/getzion/relay/api/schema/handler"
	"github.com/stretchr/testify/require"
)

func Test_ShouldReturnCorrectHandler(t *testing.T) {
	tests := []struct {
		name            string
		schema          string
		expectedHandler SchemaHandler
	}{
		{
			name:            "should return community service",
			schema:          constants.SCHEMA_COMMUNITY,
			expectedHandler: &handler.CommunityHandler{},
		},
		{
			name:            "should return conversation service",
			schema:          constants.SCHEMA_CONVERSATION,
			expectedHandler: &handler.ConversationHandler{},
		},
		{
			name:            "should return payment service",
			schema:          constants.SCHEMA_PAYMENT,
			expectedHandler: &handler.PaymentHandler{},
		},
		{
			name:            "should return user service",
			schema:          constants.SCHEMA_PERSON,
			expectedHandler: &handler.UserHandler{},
		},
		{
			name:            "should return community join schema",
			schema:          constants.SCHEMA_JOIN_COMMUNITY,
			expectedHandler: &handler.CommunityJoinHandler{},
		},
		{
			name:            "should return community leave schema",
			schema:          constants.SCHEMA_LEAVE_COMMUNITY,
			expectedHandler: &handler.CommunityLeaveHandler{},
		},
	}

	schemaManager := NewSchemaManager(nil)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler, err := schemaManager.GetSchemaHandler(tt.schema)

			require.Nil(t, err)
			require.IsType(t, tt.expectedHandler, handler)
		})
	}
}

func Test_ShouldReturnError(t *testing.T) {

	schemaManager := NewSchemaManager(nil)
	handler, err := schemaManager.GetSchemaHandler("<invalid>")

	require.NotNil(t, err)
	require.EqualError(t, err, "unknown schema: <invalid>")
	require.Nil(t, handler)
}
