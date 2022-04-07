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

	if _, err = uuid.Parse(context.Message.Descriptor.ObjectId); err != nil {
		return nil, errors.NewMessageLevelError(400, fmt.Sprintf("invalid objectId: %s", context.Message.Descriptor.ObjectId), err)
	}

	if context.Message.Descriptor.Schema != "" {
		if _, err = url.ParseRequestURI(context.Message.Descriptor.Schema); err != nil {
			return nil, errors.NewMessageLevelError(400, fmt.Sprintf("invalid schema: %s", context.Message.Descriptor.Schema), err)
		}
	}

	if context.Message.Descriptor.DateCreated == "" {
		err = fmt.Errorf("dateCreated cannot be null or empty")
		return nil, errors.NewMessageLevelError(400, err.Error(), err)
	} else if _, err = strconv.ParseInt(context.Message.Descriptor.DateCreated, 10, 64); err != nil {
		return nil, errors.NewMessageLevelError(400, fmt.Sprintf("invalid dateCreated: %s", context.Message.Descriptor.DateCreated), err)
	} else if context.Message.Descriptor.DatePublished != "" {
		_, err = strconv.ParseInt(context.Message.Descriptor.DatePublished, 10, 64)
		if err != nil {
			return nil, errors.NewMessageLevelError(400, fmt.Sprintf("invalid datePublished: %s", context.Message.Descriptor.DatePublished), err)
		}
	}

	return nil, nil
}
