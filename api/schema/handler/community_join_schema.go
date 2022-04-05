package handler

import (
	"encoding/json"
	"fmt"

	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/constants"
	"github.com/getzion/relay/api/validator"
)

type CommunityJoinHandler struct {
	storage api.Storage
}

func InitCommunityJoinHandler(storage api.Storage) *CommunityJoinHandler {
	return &CommunityJoinHandler{
		storage: storage,
	}
}

func (h *CommunityJoinHandler) Execute(data []byte, method string) (interface{}, error) {
	switch method {
	case constants.COLLECTIONS_WRITE:

		var model api.JoinCommunity
		err := json.Unmarshal(data, &model)
		if err != nil {
			return nil, err
		}

		err = validator.Struct(&model)
		if err != nil {
			return nil, err
		}

		// todo: community can be stored inside other relay?
		community, err := h.storage.GetCommunityByZid(model.CommunityZid)
		if err != nil {
			return nil, fmt.Errorf("community not found: %s", model.CommunityZid)
		}

		user, err := h.storage.GetUserByDid(model.UserDid)
		if err != nil {
			return nil, fmt.Errorf("user not found: %s", model.UserDid)
		}

		if community.PriceToJoin > 0 {
			if user.Amount-float64(community.PriceToJoin) < 0 {
				return nil, fmt.Errorf("user balance is not enought. current balance: %f", user.Amount)
			}
		}

		err = h.storage.AddUserToCommunity(community, user)
		if err != nil {
			return nil, err
		}

		return nil, nil

	default:
		return nil, fmt.Errorf("unimplemented join community method: %s", method)
	}
}
