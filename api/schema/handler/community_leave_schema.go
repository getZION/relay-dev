package handler

import (
	"encoding/json"
	"fmt"

	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/constants"
	"github.com/getzion/relay/api/datastore"
	"github.com/getzion/relay/api/validator"
)

type CommunityLeaveHandler struct {
	DataStore *datastore.Store
}

func (h *CommunityLeaveHandler) Execute(data []byte, method string) (interface{}, error) {
	switch method {
	case constants.COLLECTIONS_WRITE:

		var model api.LeaveCommunity
		err := json.Unmarshal(data, &model)
		if err != nil {
			return nil, err
		}

		err = validator.Struct(&model)
		if err != nil {
			return nil, err
		}

		community, err := h.DataStore.CommunityService.GetByZid(model.CommunityZid)
		if err != nil {
			return nil, fmt.Errorf("community not found: %s", model.CommunityZid)
		}

		user, err := h.DataStore.UserService.GetByDid(model.UserDid)
		if err != nil {
			return nil, fmt.Errorf("user not found: %s", model.UserDid)
		}

		err = h.DataStore.CommunityService.RemoveUserToCommunity(community, user)
		if err != nil {
			return nil, err
		}

		return nil, nil

	default:
		return nil, fmt.Errorf("unimplemented leave conversation method: %s", method)
	}
}
