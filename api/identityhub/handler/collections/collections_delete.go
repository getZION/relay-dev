package collections

import (
	"fmt"

	"github.com/getzion/relay/api/identityhub/errors"
	"github.com/getzion/relay/api/identityhub/handler"
	"github.com/google/uuid"
)

func CollectionsDelete(context *handler.RequestContext) ([]string, *errors.MessageLevelError) {

	var err error
	var objectId uuid.UUID

	if objectId, err = uuid.Parse(context.Message.Descriptor_.ObjectId); err != nil {
		return nil, errors.NewMessageLevelError(400, errors.ImproperlyConstructedErrorMessage, err)
	}

	//todo: process the request
	fmt.Printf("request -> objectId: %s", objectId.String())

	return nil, nil
}
