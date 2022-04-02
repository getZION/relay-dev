package handler

import (
	"encoding/json"
	"fmt"

	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/constants"
	"github.com/getzion/relay/api/validator"
)

type CommunityKickUserHandler struct {
	storage api.Storage
}

func InitCommunityKickUserHandler(storage api.Storage) *CommunityKickUserHandler {
	return &CommunityKickUserHandler{
		storage: storage,
	}
}

func (h *CommunityKickUserHandler) Execute(data []byte, method string) (interface{}, error) {
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

		community, err := h.storage.GetCommunityByZid(model.CommunityZid)
		if err != nil {
			return nil, fmt.Errorf("community not found: %s", model.CommunityZid)
		}

		user, err := h.storage.GetUserByDid(model.UserDid)
		if err != nil {
			return nil, fmt.Errorf("user not found: %s", model.UserDid)
		}

		err = h.storage.RemoveUserToCommunity(community.Zid, user.Did, "Kicked by Owner")
		if err != nil {
			return nil, err
		}

		return nil, nil

	default:
		return nil, fmt.Errorf("unimplemented commuity kick user method: %s", method)
	}
}
