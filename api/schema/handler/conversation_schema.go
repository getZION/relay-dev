package handler

import (
	"encoding/json"
	"fmt"

	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/constants"
	"github.com/getzion/relay/api/datastore"
	"github.com/getzion/relay/api/validator"
)

type ConversationHandler struct {
	DataStore *datastore.Store
}

func (h *ConversationHandler) Execute(data []byte, method string) (interface{}, error) {
	switch method {
	case constants.COLLECTIONS_QUERY:

		return h.DataStore.ConversationService.GetAll()

	case constants.COLLECTIONS_WRITE:

		var conversation api.Conversation
		err := json.Unmarshal(data, &conversation)
		if err != nil {
			return nil, err
		}

		err = validator.Struct(&conversation)
		if err != nil {
			return nil, err
		}

		return h.DataStore.ConversationService.Insert(conversation)

	default:
		return nil, fmt.Errorf("unimplemented payment method: %s", method)
	}
}
