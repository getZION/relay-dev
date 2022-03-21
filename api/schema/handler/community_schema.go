package handler

import (
	"encoding/json"
	"fmt"

	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/constants"
	"github.com/getzion/relay/api/datastore"
	"github.com/getzion/relay/api/validator"
)

type CommunityHandler struct {
	DataStore *datastore.Store
}

func (h *CommunityHandler) Execute(data []byte, method string) (interface{}, error) {
	switch method {

	case constants.COLLECTIONS_QUERY:

		return h.DataStore.CommunityService.GetAll()

	case constants.COLLECTIONS_WRITE:

		var community api.Community
		err := json.Unmarshal(data, &community)
		if err != nil {
			return nil, err
		}

		err = validator.Struct(&community)
		if err != nil {
			return nil, err
		}

		return h.DataStore.CommunityService.Insert(community)

	default:
		return nil, fmt.Errorf("unimplemented organization method: %s", method)
	}
}
