package handler

import (
	"encoding/json"
	"fmt"

	"github.com/getzion/relay/api/constants"
	"github.com/getzion/relay/api/datastore"
)

type CommunityJoinHandler struct {
	DataStore *datastore.Store
}

type joinCommunityRequestModel struct {
	UserId      int64 `json:"user_id"`
	CommunityId int64 `json:"community_id"`
}

func (h *CommunityJoinHandler) Execute(data []byte, method string) (interface{}, error) {
	switch method {
	case constants.COLLECTIONS_WRITE:

		var model joinCommunityRequestModel
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

		err = h.DataStore.CommunityService.AddUserToCommunity(community, user.Id)
		if err != nil {
			return nil, err
		}

		return nil, nil

	default:
		return nil, fmt.Errorf("unimplemented join conversation method: %s", method)
	}
}
