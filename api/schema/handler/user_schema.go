package handler

import (
	"encoding/json"
	"fmt"

	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/constants"
	"github.com/getzion/relay/api/datastore"
	"github.com/getzion/relay/api/validator"
)

type UserHandler struct {
	DataStore *datastore.Store
}

func (h *UserHandler) Execute(data []byte, method string) (interface{}, error) {
	switch method {
	case constants.COLLECTIONS_QUERY:

		return h.DataStore.UserService.GetAll()

	case constants.COLLECTIONS_WRITE:

		var user api.User
		err := json.Unmarshal(data, &user)
		if err != nil {
			return nil, err
		}

		err = validator.Struct(&user)
		if err != nil {
			return nil, err
		}

		return h.DataStore.UserService.Insert(user)

	default:
		return nil, fmt.Errorf("unimplemented user method: %s", method)
	}
}
