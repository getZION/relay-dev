package threads

import (
	"github.com/getzion/relay/api/datastore"
	"github.com/getzion/relay/api/identityhub/errors"
	hub "github.com/getzion/relay/gen/proto/identityhub/v1"
)

func ThreadsQuery(store *datastore.Store, m *hub.Message) ([]string, *errors.MessageLevelError) {
	return nil, nil
}
