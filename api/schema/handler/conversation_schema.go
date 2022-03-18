package handler

import (
	"encoding/json"
	"fmt"

	"github.com/getzion/relay/api/constants"
	"github.com/getzion/relay/api/datastore"
	v1 "github.com/getzion/relay/gen/proto/zion/v1"
)

type ConversationHandler struct {
	DataStore *datastore.Store
}

func (h *ConversationHandler) Execute(data []byte, method string) (interface{}, error) {
	switch method {
	case constants.COLLECTIONS_QUERY:

		return h.DataStore.ConversationService.GetAll()

	case constants.COLLECTIONS_WRITE:

		var conversation v1.ConversationORM
		err := json.Unmarshal(data, &conversation)
		if err != nil {
			return nil, err
		}

		return h.DataStore.ConversationService.Insert(conversation)

	default:
		return nil, fmt.Errorf("unimplemented payment method: %s", method)
	}
}
