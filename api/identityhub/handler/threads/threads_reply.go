package threads

import (
	"fmt"
	"net/url"

	"github.com/getzion/relay/api/identityhub/errors"
	"github.com/getzion/relay/api/identityhub/handler"
	"github.com/google/uuid"
)

func ThreadsReply(context *handler.RequestContext) ([]string, *errors.MessageLevelError) {

	var err error

	if _, err = uuid.Parse(context.Message.Descriptor.ObjectId); err != nil {
		return nil, errors.NewMessageLevelError(400, fmt.Sprintf("invalid objectId: %s", context.Message.Descriptor.ObjectId), err)
	} else if _, err = url.ParseRequestURI(context.Message.Descriptor.Schema); err != nil {
		return nil, errors.NewMessageLevelError(400, fmt.Sprintf("invalid schema: %s", context.Message.Descriptor.Schema), err)
	} else if _, err = uuid.Parse(context.Message.Descriptor.Root); err != nil {
		return nil, errors.NewMessageLevelError(400, fmt.Sprintf("invalid root: %s", context.Message.Descriptor.Root), err)
	} else if _, err = uuid.Parse(context.Message.Descriptor.Parent); err != nil {
		return nil, errors.NewMessageLevelError(400, fmt.Sprintf("invalid parent: %s", context.Message.Descriptor.Parent), err)
	}

	return nil, nil
}
