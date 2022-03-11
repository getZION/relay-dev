package permissions

import (
	"fmt"
	"net/url"

	"github.com/getzion/relay/api/datastore"
	"github.com/getzion/relay/api/identityhub/errors"
	hub "github.com/getzion/relay/gen/proto/identityhub/v1"
)

func PermissionsRequest(store *datastore.Store, m *hub.Message) (string, *errors.MessageLevelError) {

	var err error
	var schema *url.URL

	if schema, err = url.ParseRequestURI(m.Descriptor_.Schema); err != nil {
		return "", errors.NewMessageLevelError(400, errors.ImproperlyConstructedErrorMessage, err)
	}

	fmt.Printf("request -> schema: %s", schema.String())
	return "", nil
}
