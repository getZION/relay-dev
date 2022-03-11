package permissions

import (
	"fmt"

	"github.com/getzion/relay/api/datastore"
	"github.com/getzion/relay/api/identityhub/errors"
	hub "github.com/getzion/relay/gen/proto/identityhub/v1"
	"github.com/google/uuid"
)

func PermissionsRevoke(store *datastore.Store, m *hub.Message) ([]string, *errors.MessageLevelError) {

	var err error
	var objectId uuid.UUID

	if objectId, err = uuid.Parse(m.Descriptor_.ObjectId); err != nil {
		return nil, errors.NewMessageLevelError(400, errors.ImproperlyConstructedErrorMessage, err)
	}

	fmt.Printf("request -> objectId: %s", objectId.String())
	return nil, nil
}
