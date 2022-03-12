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
	var objectId uuid.UUID
	var root uuid.UUID
	var parent uuid.UUID
	var schema *url.URL

	if objectId, err = uuid.Parse(context.Message.Descriptor_.ObjectId); err != nil {
		return nil, errors.NewMessageLevelError(400, errors.ImproperlyConstructedErrorMessage, err)
	} else if schema, err = url.ParseRequestURI(context.Message.Descriptor_.Schema); err != nil {
		return nil, errors.NewMessageLevelError(400, errors.ImproperlyConstructedErrorMessage, err)
	} else if root, err = uuid.Parse(context.Message.Descriptor_.Root); err != nil {
		return nil, errors.NewMessageLevelError(400, errors.ImproperlyConstructedErrorMessage, err)
	} else if parent, err = uuid.Parse(context.Message.Descriptor_.Parent); err != nil {
		return nil, errors.NewMessageLevelError(400, errors.ImproperlyConstructedErrorMessage, err)
	}

	fmt.Printf("request -> objectId: %s, schema: %s, root: %s, parent: %s", objectId.String(), schema.String(), root.String(), parent.String())
	return nil, nil
}
