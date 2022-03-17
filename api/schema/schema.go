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

func NewSchemaManager(store *datastore.Store) *SchemaManager {
	schemaHandler := &SchemaManager{
		handlers: map[string]SchemaHandler{
			constants.SCHEMA_ORGANIZATION: &handler.OrganizationHandler{
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
