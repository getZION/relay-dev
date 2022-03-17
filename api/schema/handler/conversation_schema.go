package handler

import (
	"fmt"

	"github.com/getzion/relay/api/constants"
	"github.com/getzion/relay/api/datastore"
)

type ConversationHandler struct {
	DataStore *datastore.Store
}

func (h *ConversationHandler) Execute(data []byte, method string) (interface{}, error) {
	switch method {
	case constants.COLLECTIONS_QUERY:

		return h.DataStore.ConversationService.GetAll()

	case constants.COLLECTIONS_WRITE:

		return h.DataStore.ConversationService.Insert(data)

	default:
		return nil, fmt.Errorf("unimplemented payment method: %s", method)
	}
}
