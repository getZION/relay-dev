package schema

import (
	"fmt"

	"github.com/getzion/relay/api/constants"
	"github.com/getzion/relay/api/datastore"
	"github.com/getzion/relay/api/schema/handler"
)

type SchemaManager struct {
	handlers map[string]SchemaHandler
}

type SchemaHandler interface {
	Execute(data []byte, method string) (interface{}, error)
}

// todo: can we split the method operations via structs? like this, we don't need to worry about return types ...
// type MethodExecutor interface {
// 	Execute(data []byte) (interface{}, error)
// }

func NewSchemaManager(store *datastore.Store) *SchemaManager {
	schemaHandler := &SchemaManager{
		handlers: map[string]SchemaHandler{
			constants.SCHEMA_COMMUNITY: &handler.CommunityHandler{
				DataStore: store,
			},
			constants.SCHEMA_CONVERSATION: &handler.ConversationHandler{
				DataStore: store,
			},
			constants.SCHEMA_PAYMENT: &handler.PaymentHandler{
				DataStore: store,
			},
			constants.SCHEMA_PERSON: &handler.UserHandler{
				DataStore: store,
			},
			constants.SCHEMA_JOIN_COMMUNITY: &handler.CommunityJoinHandler{
				DataStore: store,
			},
			constants.SCHEMA_LEAVE_COMMUNITY: &handler.CommunityLeaveHandler{
				DataStore: store,
			},
			constants.SCHEMA_COMMENT: &handler.CommentHandler{
				DataStore: store,
			},
		},
	}

	return schemaHandler
}

func (sm *SchemaManager) GetSchemaHandler(schema string) (SchemaHandler, error) {
	if handler, ok := sm.handlers[schema]; ok {
		return handler, nil
	}

	return nil, fmt.Errorf("unknown schema: %s", schema)
}
