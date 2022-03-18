package handler

import (
	"encoding/json"
	"fmt"

	"github.com/getzion/relay/api/constants"
	"github.com/getzion/relay/api/datastore"
	v1 "github.com/getzion/relay/gen/proto/zion/v1"
)

type CommunityHandler struct {
	DataStore *datastore.Store
}

func (h *CommunityHandler) Execute(data []byte, method string) (interface{}, error) {
	switch method {

	case constants.COLLECTIONS_QUERY:

		return h.DataStore.CommunityService.GetAll()

	case constants.COLLECTIONS_WRITE:

		var community v1.CommunityORM
		err := json.Unmarshal(data, &community)
		if err != nil {
			return nil, err
		}

		return h.DataStore.CommunityService.Insert(community)

	default:
		return nil, fmt.Errorf("unimplemented organization method: %s", method)
	}
}
