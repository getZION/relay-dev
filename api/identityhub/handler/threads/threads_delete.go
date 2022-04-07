package threads

import (
	"fmt"

	"github.com/getzion/relay/api/identityhub/errors"
	"github.com/getzion/relay/api/identityhub/handler"
	"github.com/google/uuid"
)

func ThreadsDelete(context *handler.RequestContext) ([]string, *errors.MessageLevelError) {

	var err error

	if _, err = uuid.Parse(context.Message.Descriptor.Root); err != nil {
		return nil, errors.NewMessageLevelError(400, fmt.Sprintf("invalid root: %s", context.Message.Descriptor.Root), err)
	}

	return nil, nil
}
