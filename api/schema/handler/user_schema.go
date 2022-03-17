package handler

import (
	"fmt"

	"github.com/getzion/relay/api/constants"
	"github.com/getzion/relay/api/datastore"
)

type UserHandler struct {
	DataStore *datastore.Store
}

func (h *UserHandler) Execute(data []byte, method string) (interface{}, error) {
	switch method {
	case constants.COLLECTIONS_QUERY:

		return h.DataStore.UserService.GetAll()

	case constants.COLLECTIONS_WRITE:

		return h.DataStore.UserService.Insert(data)

	default:
		return nil, fmt.Errorf("unimplemented user method: %s", method)
	}
}
