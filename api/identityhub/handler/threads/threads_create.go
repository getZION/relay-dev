package threads

import (
	"fmt"
	"net/url"

	"github.com/getzion/relay/api/datastore"
	"github.com/getzion/relay/api/identityhub/errors"
	hub "github.com/getzion/relay/gen/proto/identityhub/v1"
	"github.com/google/uuid"
)

func ThreadsCreate(store *datastore.Store, m *hub.Message) ([]string, *errors.MessageLevelError) {

	var err error
	var objectId uuid.UUID
	var schema *url.URL

	if objectId, err = uuid.Parse(m.Descriptor_.ObjectId); err != nil {
		return nil, errors.NewMessageLevelError(400, errors.ImproperlyConstructedErrorMessage, err)
	} else if schema, err = url.ParseRequestURI(m.Descriptor_.Schema); err != nil {
		return nil, errors.NewMessageLevelError(400, errors.ImproperlyConstructedErrorMessage, err)
	}

	fmt.Printf("request -> objectId: %s, schema: %s", objectId.String(), schema.String())
	return nil, nil
}
