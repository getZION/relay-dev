package collections

import (
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"

	"github.com/gabriel-vasile/mimetype"
	"github.com/getzion/relay/api/identityhub/errors"
	"github.com/getzion/relay/api/identityhub/handler"
	"github.com/google/uuid"
)

func CollectionsQuery(context *handler.RequestContext) ([]string, *errors.MessageLevelError) {

	var err error

	if _, err = uuid.Parse(context.Message.Descriptor_.ObjectId); err != nil {
		return nil, errors.NewMessageLevelError(400, fmt.Sprintf("invalid objectId: %s", context.Message.Descriptor_.ObjectId), err)
	}

	if context.Message.Descriptor_.Schema != "" {
		if _, err = url.ParseRequestURI(context.Message.Descriptor_.Schema); err != nil {
			return nil, errors.NewMessageLevelError(400, fmt.Sprintf("invalid schema: %s", context.Message.Descriptor_.Schema), err)
		}
	}

	if context.Message.Descriptor_.DataFormat != "" {
		if dataFormat := mimetype.Lookup(context.Message.Descriptor_.DataFormat); dataFormat == nil {
			err = fmt.Errorf("invalid dataFormat: %s", context.Message.Descriptor_.DataFormat)
			return nil, errors.NewMessageLevelError(400, err.Error(), err)
		}
	}

	if context.Message.Descriptor_.DateSort != "" && (context.Message.Descriptor_.DateSort != "createdAscending" && context.Message.Descriptor_.DateSort != "createdDescending" &&
		context.Message.Descriptor_.DateSort != "publishedAscending" && context.Message.Descriptor_.DateSort != "publishedDescending") {
		err = fmt.Errorf("invalid dateSort: %s", context.Message.Descriptor_.DateSort)
		return nil, errors.NewMessageLevelError(400, err.Error(), err)
	}

	//todo: check data & dataFormat only for application/json or do we need provide other formats?

	service, err := context.Store.GetServiceBySchema(context.Message.Descriptor_.Schema)
	if err != nil {
		return nil, errors.NewMessageLevelError(400, err.Error(), err)
	}

	data, err := service.GetAll()
	if err != nil {
		return nil, errors.NewMessageLevelError(500, err.Error(), err)
	}

	var entries []string
	v := reflect.ValueOf(data)
	for i := 0; i < v.Len(); i++ {
		val := v.Index(i).Interface()
		result, err := json.Marshal(&val)
		if err != nil {
			return nil, errors.NewMessageLevelError(500, err.Error(), err)
		}
		entries = append(entries, string(result))
	}

	return entries, nil
}
