package permissions

import (
	"fmt"
	"net/url"

	"github.com/getzion/relay/api/identityhub/errors"
	"github.com/getzion/relay/api/identityhub/handler"
)

func PermissionsQuery(context *handler.RequestContext) ([]string, *errors.MessageLevelError) {

	var err error

	if _, err = url.ParseRequestURI(context.Message.Descriptor.Schema); err != nil {
		return nil, errors.NewMessageLevelError(400, fmt.Sprintf("invalid schema: %s", context.Message.Descriptor.Schema), err)
	}

	return nil, nil
}
