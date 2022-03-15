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

	if _, err = uuid.Parse(context.Message.Descriptor_.ObjectId); err != nil {
		return nil, errors.NewMessageLevelError(400, fmt.Sprintf("invalid objectId: %s", context.Message.Descriptor_.ObjectId), err)
	} else if _, err = url.ParseRequestURI(context.Message.Descriptor_.Schema); err != nil {
		return nil, errors.NewMessageLevelError(400, fmt.Sprintf("invalid schema: %s", context.Message.Descriptor_.Schema), err)
	} else if _, err = uuid.Parse(context.Message.Descriptor_.Root); err != nil {
		return nil, errors.NewMessageLevelError(400, fmt.Sprintf("invalid root: %s", context.Message.Descriptor_.Root), err)
	} else if _, err = uuid.Parse(context.Message.Descriptor_.Parent); err != nil {
		return nil, errors.NewMessageLevelError(400, fmt.Sprintf("invalid parent: %s", context.Message.Descriptor_.Parent), err)
	}

	return nil, nil
}
