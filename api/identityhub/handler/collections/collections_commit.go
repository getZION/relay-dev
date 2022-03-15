package collections

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/getzion/relay/api/identityhub/errors"
	"github.com/getzion/relay/api/identityhub/handler"
	"github.com/google/uuid"
)

func CollectionsCommit(context *handler.RequestContext) ([]string, *errors.MessageLevelError) {

	var err error

	if _, err = uuid.Parse(context.Message.Descriptor_.ObjectId); err != nil {
		return nil, errors.NewMessageLevelError(400, fmt.Sprintf("invalid objectId: %s", context.Message.Descriptor_.ObjectId), err)
	}

	if context.Message.Descriptor_.Schema != "" {
		if _, err = url.ParseRequestURI(context.Message.Descriptor_.Schema); err != nil {
			return nil, errors.NewMessageLevelError(400, fmt.Sprintf("invalid schema: %s", context.Message.Descriptor_.Schema), err)
		}
	}

	if context.Message.Descriptor_.DateCreated == "" {
		err = fmt.Errorf("dateCreated cannot be null or empty")
		return nil, errors.NewMessageLevelError(400, err.Error(), err)
	} else if _, err = strconv.ParseInt(context.Message.Descriptor_.DateCreated, 10, 64); err != nil {
		return nil, errors.NewMessageLevelError(400, fmt.Sprintf("invalid dateCreated: %s", context.Message.Descriptor_.DateCreated), err)
	} else if context.Message.Descriptor_.DatePublished != "" {
		_, err = strconv.ParseInt(context.Message.Descriptor_.DatePublished, 10, 64)
		if err != nil {
			return nil, errors.NewMessageLevelError(400, fmt.Sprintf("invalid datePublished: %s", context.Message.Descriptor_.DatePublished), err)
		}
	}

	//todo: process the request
	return nil, nil
}
