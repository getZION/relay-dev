package handler

import (
	"encoding/json"
	"fmt"

	"github.com/getzion/relay/api/constants"
	"github.com/getzion/relay/api/datastore"
	v1 "github.com/getzion/relay/gen/proto/zion/v1"
)

type UserHandler struct {
	DataStore *datastore.Store
}

func (h *UserHandler) Execute(data []byte, method string) (interface{}, error) {
	switch method {
	case constants.COLLECTIONS_QUERY:

		return h.DataStore.UserService.GetAll()

	case constants.COLLECTIONS_WRITE:

		var user v1.UserORM
		err := json.Unmarshal(data, &user)
		if err != nil {
			return nil, err
		}

		return h.DataStore.UserService.Insert(user)

	default:
		return nil, fmt.Errorf("unimplemented user method: %s", method)
	}
}
