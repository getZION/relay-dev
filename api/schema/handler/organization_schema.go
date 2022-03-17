package handler

import (
	"fmt"

	"github.com/getzion/relay/api/constants"
	"github.com/getzion/relay/api/datastore"
)

type OrganizationHandler struct {
	DataStore *datastore.Store
}

func (h *OrganizationHandler) Execute(data []byte, method string) (interface{}, error) {
	switch method {

	case constants.COLLECTIONS_QUERY:

		return h.DataStore.CommunityService.GetAll()

	case constants.COLLECTIONS_WRITE:

		return h.DataStore.CommunityService.Insert(data)

	default:
		return nil, fmt.Errorf("unimplemented organization method: %s", method)
	}
}
