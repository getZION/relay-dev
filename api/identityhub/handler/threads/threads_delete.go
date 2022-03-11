package threads

import (
	"fmt"

	"github.com/getzion/relay/api/datastore"
	"github.com/getzion/relay/api/identityhub/errors"
	hub "github.com/getzion/relay/gen/proto/identityhub/v1"
	"github.com/google/uuid"
)

func ThreadsDelete(store *datastore.Store, m *hub.Message) (string, *errors.MessageLevelError) {

	var err error
	var root uuid.UUID

	if root, err = uuid.Parse(m.Descriptor_.Root); err != nil {
		return "", errors.NewMessageLevelError(400, errors.ImproperlyConstructedErrorMessage, err)
	}

	fmt.Printf("request -> root: %s", root.String())
	return "", nil
}
