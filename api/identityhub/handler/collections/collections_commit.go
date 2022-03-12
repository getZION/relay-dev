package collections

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/getzion/relay/api/identityhub/errors"
	"github.com/getzion/relay/api/identityhub/handler"
	"github.com/google/uuid"
)

func CollectionsCommit(context *handler.RequestContext) ([]string, *errors.MessageLevelError) {

	var err error
	var objectId uuid.UUID
	var schema *url.URL
	var dateCreated int64
	var datePublished int64

	if objectId, err = uuid.Parse(context.Message.Descriptor_.ObjectId); err != nil {
		return nil, errors.NewMessageLevelError(400, errors.ImproperlyConstructedErrorMessage, err)
	}

	if context.Message.Descriptor_.Schema != "" {
		if schema, err = url.ParseRequestURI(context.Message.Descriptor_.Schema); err != nil {
			return nil, errors.NewMessageLevelError(400, errors.ImproperlyConstructedErrorMessage, err)
		}
	}

	if context.Message.Descriptor_.DateCreated == "" {
		return nil, errors.NewMessageLevelError(400, errors.ImproperlyConstructedErrorMessage, fmt.Errorf("dateCreated cannot be null"))
	} else if dateCreated, err = strconv.ParseInt(context.Message.Descriptor_.DateCreated, 10, 64); err != nil {
		return nil, errors.NewMessageLevelError(400, errors.ImproperlyConstructedErrorMessage, err)
	} else if context.Message.Descriptor_.DatePublished != "" {
		datePublished, err = strconv.ParseInt(context.Message.Descriptor_.DatePublished, 10, 64)
		if err != nil {
			return nil, errors.NewMessageLevelError(400, errors.ImproperlyConstructedErrorMessage, err)
		}
	}

	//todo: process the request
	fmt.Printf("request -> objectId: %s, dateCreated: %d", objectId.String(), dateCreated)
	if schema != nil || datePublished == 0 {

	}

	return nil, nil
}
