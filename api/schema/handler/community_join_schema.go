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

type requestData struct {
	UserId         int64 `json:"user_id"`
	OrganizationId int64 `json:"organization_id"`
}

func (h *CommunityJoinHandler) Execute(data []byte, method string) (interface{}, error) {
	switch method {
	case constants.COLLECTIONS_WRITE:

		var request requestData
		err := json.Unmarshal(data, &request)
		if err != nil {
			return nil, err
		}

		community, err := h.DataStore.CommunityService.GetById(request.OrganizationId)
		if err != nil {
			return nil, err
		}

		user, err := h.DataStore.UserService.GetById(request.UserId)
		if err != nil {
			return nil, err
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
