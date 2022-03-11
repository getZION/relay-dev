package collections

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/gabriel-vasile/mimetype"
	"github.com/getzion/relay/api/datastore"
	"github.com/getzion/relay/api/identityhub/errors"
	hub "github.com/getzion/relay/gen/proto/identityhub/v1"
	"github.com/google/uuid"
)

func CollectionsQuery(store *datastore.Store, m *hub.Message) (string, *errors.MessageLevelError) {

	var err error
	var objectId uuid.UUID
	var schema *url.URL
	var dataFormat *mimetype.MIME

	if objectId, err = uuid.Parse(m.Descriptor_.ObjectId); err != nil {
		return "", errors.NewMessageLevelError(400, errors.ImproperlyConstructedErrorMessage, err)
	}

	if m.Descriptor_.Schema != "" {
		if schema, err = url.ParseRequestURI(m.Descriptor_.Schema); err != nil {
			return "", errors.NewMessageLevelError(400, errors.ImproperlyConstructedErrorMessage, err)
		}
	}

	if m.Descriptor_.DataFormat != "" {
		if dataFormat = mimetype.Lookup(m.Descriptor_.DataFormat); dataFormat == nil {
			return "", errors.NewMessageLevelError(400, errors.ImproperlyConstructedErrorMessage, fmt.Errorf("invalid dataFormat: %s", m.Descriptor_.DataFormat))
		}
	}

	if m.Descriptor_.DateSort != "" && (m.Descriptor_.DateSort != "createdAscending" && m.Descriptor_.DateSort != "createdDescending" &&
		m.Descriptor_.DateSort != "publishedAscending" && m.Descriptor_.DateSort != "publishedDescending") {
		return "", errors.NewMessageLevelError(400, errors.ImproperlyConstructedErrorMessage, fmt.Errorf("invalid dateSort: %s", m.Descriptor_.DateSort))
	}

	//todo: check data & dataFormat only for application/json or do we need provide other formats?

	//todo: process the request
	fmt.Printf("request -> objectId: %s", objectId.String())
	if schema != nil {

	}

	communities := store.CommunityService.GetAll()
	result, err := json.Marshal(communities)
	if err != nil {
		return "", errors.NewMessageLevelError(500, errors.RequestLevelProcessingErrorMessage, err)
	}

	//todo: return entry as a array
	return string(result), nil
}
