package threads

import (
	"fmt"

	"github.com/getzion/relay/api/identityhub/errors"
	"github.com/getzion/relay/api/identityhub/handler"
	"github.com/google/uuid"
)

func ThreadsDelete(context *handler.RequestContext) ([]string, *errors.MessageLevelError) {

	var err error
	var root uuid.UUID

	if root, err = uuid.Parse(context.Message.Descriptor_.Root); err != nil {
		return nil, errors.NewMessageLevelError(400, errors.ImproperlyConstructedErrorMessage, err)
	}

	fmt.Printf("request -> root: %s", root.String())
	return nil, nil
}
