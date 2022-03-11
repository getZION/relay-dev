package collections

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/getzion/relay/api/datastore"
	"github.com/getzion/relay/api/identityhub/errors"
	hub "github.com/getzion/relay/gen/proto/identityhub/v1"
	"github.com/google/uuid"
)

func CollectionsCommit(store *datastore.Store, m *hub.Message) (string, *errors.MessageLevelError) {

	var err error
	var objectId uuid.UUID
	var schema *url.URL
	var dateCreated int64
	var datePublished int64

	if objectId, err = uuid.Parse(m.Descriptor_.ObjectId); err != nil {
		return "", errors.NewMessageLevelError(400, errors.ImproperlyConstructedErrorMessage, err)
	}

	if m.Descriptor_.Schema != "" {
		if schema, err = url.ParseRequestURI(m.Descriptor_.Schema); err != nil {
			return "", errors.NewMessageLevelError(400, errors.ImproperlyConstructedErrorMessage, err)
		}
	}

	if m.Descriptor_.DateCreated == "" {
		return "", errors.NewMessageLevelError(400, errors.ImproperlyConstructedErrorMessage, fmt.Errorf("dateCreated cannot be null"))
	} else if dateCreated, err = strconv.ParseInt(m.Descriptor_.DateCreated, 10, 64); err != nil {
		return "", errors.NewMessageLevelError(400, errors.ImproperlyConstructedErrorMessage, err)
	} else if m.Descriptor_.DatePublished != "" {
		datePublished, err = strconv.ParseInt(m.Descriptor_.DatePublished, 10, 64)
		if err != nil {
			return "", errors.NewMessageLevelError(400, errors.ImproperlyConstructedErrorMessage, err)
		}
	}

	//todo: process the request
	fmt.Printf("request -> objectId: %s, dateCreated: %d", objectId.String(), dateCreated)
	if schema != nil || datePublished == 0 {

	}

	return "", nil
}
