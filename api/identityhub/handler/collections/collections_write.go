package collections

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/getzion/relay/api/identityhub/errors"
	"github.com/getzion/relay/api/identityhub/handler"
	"github.com/google/uuid"
)

func CollectionsWrite(context *handler.RequestContext) ([]string, *errors.MessageLevelError) {

	/*

		Processing Instructions
		When processing a CollectionsWrite message, Hub instances MUST perform the following additional steps:

		If the incoming message has a higher dateCreated value than all of the other messages for the logical entry known to the Hub Instance, the message MUST be designated as the latest state of the logical entry and fully replace all previous messages for the entry.
		If the incoming message has a lower dateCreated value than the message that represents the current state of the logical entry, the message MUST NOT be applied to the logical entry and its data MAY be discarded.
		If the incoming message has a dateCreated value equal to the message that represents the current state of the logical entry, the incoming message’s IPFS CID and the IPFS CID of the message that represents the current state must be lexicographically compared and handled as follows:
		If the incoming message has a higher lexicographic value than the message that represents the current state, perform the actions described in Step 1 of this instruction set.
		If the incoming message has a lower lexicographic value than the message that represents the current state, perform the actions described in Step 2 of this instruction set.

	*/

	var err error

	if _, err = uuid.Parse(context.Message.Descriptor_.ObjectId); err != nil {
		return nil, errors.NewMessageLevelError(400, fmt.Sprintf("invalid objectId: %s", context.Message.Descriptor_.ObjectId), err)
	} else if context.Message.Descriptor_.DateCreated == "" {
		err = fmt.Errorf("dateCreated cannot be null or empty")
		return nil, errors.NewMessageLevelError(400, err.Error(), err)
	} else if _, err := strconv.ParseInt(context.Message.Descriptor_.DateCreated, 10, 64); err != nil {
		return nil, errors.NewMessageLevelError(400, fmt.Sprintf("invalid dateCreated: %s", context.Message.Descriptor_.DateCreated), err)
	} else if context.Message.Descriptor_.Schema != "" {
		if _, err = url.ParseRequestURI(context.Message.Descriptor_.Schema); err != nil {
			return nil, errors.NewMessageLevelError(400, fmt.Sprintf("invalid schema: %s", context.Message.Descriptor_.Schema), err)
		}
	}

	if context.Message.Descriptor_.DatePublished != "" {
		_, err := strconv.ParseInt(context.Message.Descriptor_.DatePublished, 10, 64)
		if err != nil {
			return nil, errors.NewMessageLevelError(400, fmt.Sprintf("invalid datePublished: %s", context.Message.Descriptor_.DatePublished), err)
		}
	}

	if strings.Trim(context.Message.Data, " ") == "" {
		err = fmt.Errorf("data cannot be empty")
		return nil, errors.NewMessageLevelError(400, err.Error(), err)
	}

	schemaHandler, err := context.SchemaManager.GetSchemaHandler(context.Message.Descriptor_.Schema)
	if err != nil {
		return nil, errors.NewMessageLevelError(400, err.Error(), err)
	}

	data, err := schemaHandler.Execute([]byte(context.Message.Data), context.Message.Descriptor_.Method)
	if err != nil {
		return nil, errors.NewMessageLevelError(400, err.Error(), err)
	}

	if data != nil {
		var entries []string
		json, err := json.Marshal(&data)
		if err != nil {
			return nil, errors.NewMessageLevelError(500, err.Error(), err)
		}

		entries = append(entries, string(json))
		return entries, nil
	}

	return nil, nil
}
