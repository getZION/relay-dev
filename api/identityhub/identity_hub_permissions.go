package identityhub

import (
	"fmt"
	"net/url"

	"github.com/getzion/relay/api/datastore"
	hub "github.com/getzion/relay/gen/proto/identityhub/v1"
	"github.com/google/uuid"
)

func PermissionsRequest(store *datastore.Store, m *hub.Message) (string, *MessageLevelError) {

	var err error
	var schema *url.URL

	if schema, err = url.ParseRequestURI(m.Descriptor_.Schema); err != nil {
		return "", NewMessageLevelError(400, improperlyConstructedErrorMessage, err)
	}

	fmt.Printf("request -> schema: %s", schema.String())

	return "", nil
}

func PermissionsQuery(store *datastore.Store, m *hub.Message) (string, *MessageLevelError) {

	var err error
	var schema *url.URL

	if schema, err = url.ParseRequestURI(m.Descriptor_.Schema); err != nil {
		return "", NewMessageLevelError(400, improperlyConstructedErrorMessage, err)
	}

	fmt.Printf("request -> schema: %s", schema.String())

	return "", nil
}

func PermissionsGrant(store *datastore.Store, m *hub.Message) (string, *MessageLevelError) {

	var err error
	var schema *url.URL

	if schema, err = url.ParseRequestURI(m.Descriptor_.Schema); err != nil {
		return "", NewMessageLevelError(400, improperlyConstructedErrorMessage, err)
	}

	fmt.Printf("request -> schema: %s", schema.String())

	return "", nil
}

func PermissionsRevoke(store *datastore.Store, m *hub.Message) (string, *MessageLevelError) {

	var err error
	var objectId uuid.UUID

	if objectId, err = uuid.Parse(m.Descriptor_.ObjectId); err != nil {
		return "", NewMessageLevelError(400, improperlyConstructedErrorMessage, err)
	}

	fmt.Printf("request -> objectId: %s", objectId.String())

	return "", nil
}
