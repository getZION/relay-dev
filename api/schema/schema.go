package schema

import (
	"fmt"

	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/constants"
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

func NewSchemaManager(storage api.Storage) *SchemaManager {
	schemaHandler := &SchemaManager{
		handlers: map[string]SchemaHandler{
			constants.SCHEMA_COMMUNITY:       handler.InitCommunityHandler(storage),
			constants.SCHEMA_CONVERSATION:    handler.InitConversationHandler(storage),
			constants.SCHEMA_PAYMENT:         handler.InitPaymentHandler(storage),
			constants.SCHEMA_PERSON:          handler.InitUserHandler(storage),
			constants.SCHEMA_JOIN_COMMUNITY:  handler.InitCommunityJoinHandler(storage),
			constants.SCHEMA_LEAVE_COMMUNITY: handler.InitCommunityLeaveHandler(storage),
			constants.SCHEMA_COMMENT:         handler.InitCommentHandler(storage),
			constants.SCHEMA_KICK_USER:       handler.InitCommunityKickUserHandler(storage),
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
