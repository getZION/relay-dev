package handler

import (
	"encoding/json"
	"fmt"

	"github.com/getzion/relay/api/constants"
	"github.com/getzion/relay/api/datastore"
)

type CommunityLeaveHandler struct {
	DataStore *datastore.Store
}

type leaveCommunityRequestModel struct {
	UserId      int64 `json:"user_id"`
	CommunityId int64 `json:"community_id"`
}

func (h *CommunityLeaveHandler) Execute(data []byte, method string) (interface{}, error) {
	switch method {
	case constants.COLLECTIONS_WRITE:

		var model leaveCommunityRequestModel
		err := json.Unmarshal(data, &model)
		if err != nil {
			return nil, err
		}

		community, err := h.DataStore.CommunityService.GetById(model.CommunityId)
		if err != nil {
			return nil, fmt.Errorf("community not found: %d", model.CommunityId)
		}

		user, err := h.DataStore.UserService.GetById(model.UserId)
		if err != nil {
			return nil, fmt.Errorf("user not found: %d", model.UserId)
		}

		err = h.DataStore.CommunityService.RemoveUserToCommunity(community, user.Id)
		if err != nil {
			return nil, err
		}

		return nil, nil

	default:
		return nil, fmt.Errorf("unimplemented leave conversation method: %s", method)
	}
}
