package threads

import (
	"fmt"
	"net/url"

	"github.com/getzion/relay/api/identityhub/errors"
	"github.com/getzion/relay/api/identityhub/handler"
	"github.com/google/uuid"
)

func ThreadsCreate(context *handler.RequestContext) ([]string, *errors.MessageLevelError) {

	var err error
	var objectId uuid.UUID
	var schema *url.URL

	if objectId, err = uuid.Parse(context.Message.Descriptor_.ObjectId); err != nil {
		return nil, errors.NewMessageLevelError(400, errors.ImproperlyConstructedErrorMessage, err)
	} else if schema, err = url.ParseRequestURI(context.Message.Descriptor_.Schema); err != nil {
		return nil, errors.NewMessageLevelError(400, errors.ImproperlyConstructedErrorMessage, err)
	}

	fmt.Printf("request -> objectId: %s, schema: %s", objectId.String(), schema.String())
	return nil, nil
}
