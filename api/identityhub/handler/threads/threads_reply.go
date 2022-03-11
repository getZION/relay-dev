package threads

import (
	"fmt"
	"net/url"

	"github.com/getzion/relay/api/datastore"
	"github.com/getzion/relay/api/identityhub/errors"
	hub "github.com/getzion/relay/gen/proto/identityhub/v1"
	"github.com/google/uuid"
)

func ThreadsReply(store *datastore.Store, m *hub.Message) (string, *errors.MessageLevelError) {

	var err error
	var objectId uuid.UUID
	var root uuid.UUID
	var parent uuid.UUID
	var schema *url.URL

	if objectId, err = uuid.Parse(m.Descriptor_.ObjectId); err != nil {
		return "", errors.NewMessageLevelError(400, errors.ImproperlyConstructedErrorMessage, err)
	} else if schema, err = url.ParseRequestURI(m.Descriptor_.Schema); err != nil {
		return "", errors.NewMessageLevelError(400, errors.ImproperlyConstructedErrorMessage, err)
	} else if root, err = uuid.Parse(m.Descriptor_.Root); err != nil {
		return "", errors.NewMessageLevelError(400, errors.ImproperlyConstructedErrorMessage, err)
	} else if parent, err = uuid.Parse(m.Descriptor_.Parent); err != nil {
		return "", errors.NewMessageLevelError(400, errors.ImproperlyConstructedErrorMessage, err)
	}

	fmt.Printf("request -> objectId: %s, schema: %s, root: %s, parent: %s", objectId.String(), schema.String(), root.String(), parent.String())

	return "", nil
}
