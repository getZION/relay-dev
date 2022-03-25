package handler

import (
	"encoding/json"
	"fmt"

	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/constants"
	"github.com/getzion/relay/api/datastore"
	"github.com/getzion/relay/api/validator"
)

type CommentHandler struct {
	DataStore *datastore.Store
}

func (h *CommentHandler) Execute(data []byte, method string) (interface{}, error) {
	switch method {
	case constants.COLLECTIONS_QUERY:

		return h.DataStore.CommentService.GetAll()

	case constants.COLLECTIONS_WRITE:

		var comment api.Comment
		err := json.Unmarshal(data, &comment)
		if err != nil {
			return nil, err
		}

		err = validator.Struct(&comment)
		if err != nil {
			return nil, err
		}

		return h.DataStore.CommentService.Insert(comment)

	default:
		return nil, fmt.Errorf("unimplemented comment method: %s", method)
	}
}
