package handler

import (
	"encoding/json"
	"fmt"

	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/constants"
	"github.com/getzion/relay/api/validator"
)

type CommentHandler struct {
	storage api.Storage
}

func InitCommentHandler(connection api.Storage) *CommentHandler {
	return &CommentHandler{
		storage: connection,
	}
}

func (h *CommentHandler) Execute(data []byte, method string) (interface{}, error) {
	switch method {
	case constants.COLLECTIONS_QUERY:

		return h.storage.GetComments()

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

		err = h.storage.InsertComment(&comment)
		if err != nil {
			return nil, err
		}

		return comment, nil

	default:
		return nil, fmt.Errorf("unimplemented comment method: %s", method)
	}
}
