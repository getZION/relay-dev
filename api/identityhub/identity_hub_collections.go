package identityhub

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/getzion/relay/api/datastore"
	hub "github.com/getzion/relay/gen/proto/identityhub/v1"
	v1 "github.com/getzion/relay/gen/proto/zion/v1"
	"github.com/google/uuid"
)

func CollectionsQuery(store *datastore.Store, m *hub.Message) (string, *MessageLevelError) {

	var err error
	var objectId uuid.UUID
	var schema *url.URL
	var dataFormat *mimetype.MIME

	if objectId, err = uuid.Parse(m.Descriptor_.ObjectId); err != nil {
		return "", NewMessageLevelError(400, improperlyConstructedErrorMessage, err)
	}

	if m.Descriptor_.Schema != "" {
		if schema, err = url.ParseRequestURI(m.Descriptor_.Schema); err != nil {
			return "", NewMessageLevelError(400, improperlyConstructedErrorMessage, err)
		}
	}

	if m.Descriptor_.DataFormat != "" {
		if dataFormat = mimetype.Lookup(m.Descriptor_.DataFormat); dataFormat == nil {
			return "", NewMessageLevelError(400, improperlyConstructedErrorMessage, fmt.Errorf("invalid mime type: %s", m.Descriptor_.DataFormat))
		}
	}

	if m.Descriptor_.DateSort != "" && (m.Descriptor_.DateSort != "createdAscending" && m.Descriptor_.DateSort != "createdDescending" &&
		m.Descriptor_.DateSort != "publishedAscending" && m.Descriptor_.DateSort != "publishedDescending") {
		return "", NewMessageLevelError(400, improperlyConstructedErrorMessage, fmt.Errorf("invalid mime type: %s", m.Descriptor_.DataFormat))
	}

	//todo: check data & dataFormat only for application/json or do we need provide other formats?

	//todo: process the request
	fmt.Printf("request -> objectId: %s", objectId.String())
	if schema != nil {

	}

	communities := store.CommunityService.GetAll()
	result, err := json.Marshal(communities)
	if err != nil {
		return "", NewMessageLevelError(500, requestLevelProcessingErrorMessage, err)
	}

	return string(result), nil
}

func CollectionsWrite(store *datastore.Store, m *hub.Message) (string, *MessageLevelError) {

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
	var objectId uuid.UUID
	var schema *url.URL
	var dateCreated int64
	var datePublished int64

	if objectId, err = uuid.Parse(m.Descriptor_.ObjectId); err != nil {
		return "", NewMessageLevelError(400, improperlyConstructedErrorMessage, err)
	} else if m.Descriptor_.DateCreated == "" {
		return "", NewMessageLevelError(400, improperlyConstructedErrorMessage, fmt.Errorf("dateCreated cannot be null or empty"))
	} else if dateCreated, err = strconv.ParseInt(m.Descriptor_.DateCreated, 10, 64); err != nil {
		return "", NewMessageLevelError(400, improperlyConstructedErrorMessage, err)
	} else if m.Descriptor_.Schema != "" {
		if schema, err = url.ParseRequestURI(m.Descriptor_.Schema); err != nil {
			return "", NewMessageLevelError(400, improperlyConstructedErrorMessage, err)
		}
	}

	if m.Descriptor_.DatePublished != "" {
		datePublished, err = strconv.ParseInt(m.Descriptor_.DatePublished, 10, 64)
		if err != nil {
			return "", NewMessageLevelError(400, improperlyConstructedErrorMessage, err)
		}
	}

	//todo: process the request
	fmt.Printf("request -> objectId: %s, dateCreated: %d", objectId.String(), dateCreated)
	if schema != nil || datePublished == 0 {

	}

	if strings.Trim(m.Data, " ") == "" {
		return "", NewMessageLevelError(400, improperlyConstructedErrorMessage, fmt.Errorf("data cannot be null or empty"))
	}

	var community v1.CommunityORM
	json.Unmarshal([]byte(m.Data), &community)
	store.CommunityService.Insert(&community)

	return "", nil
}

func CollectionsCommit(store *datastore.Store, m *hub.Message) (string, *MessageLevelError) {

	var err error
	var objectId uuid.UUID
	var schema *url.URL
	var dateCreated int64
	var datePublished int64

	if objectId, err = uuid.Parse(m.Descriptor_.ObjectId); err != nil {
		return "", NewMessageLevelError(400, improperlyConstructedErrorMessage, err)
	}

	if m.Descriptor_.Schema != "" {
		if schema, err = url.ParseRequestURI(m.Descriptor_.Schema); err != nil {
			return "", NewMessageLevelError(400, improperlyConstructedErrorMessage, err)
		}
	}

	if m.Descriptor_.DateCreated == "" {
		return "", NewMessageLevelError(400, improperlyConstructedErrorMessage, fmt.Errorf("dateCreated cannot be null"))
	} else if dateCreated, err = strconv.ParseInt(m.Descriptor_.DateCreated, 10, 64); err != nil {
		return "", NewMessageLevelError(400, improperlyConstructedErrorMessage, err)
	} else if m.Descriptor_.DatePublished != "" {
		datePublished, err = strconv.ParseInt(m.Descriptor_.DatePublished, 10, 64)
		if err != nil {
			return "", NewMessageLevelError(400, improperlyConstructedErrorMessage, err)
		}
	}

	//todo: process the request
	fmt.Printf("request -> objectId: %s, dateCreated: %d", objectId.String(), dateCreated)
	if schema != nil || datePublished == 0 {

	}

	return "", nil
}

func CollectionsDelete(store *datastore.Store, m *hub.Message) (string, *MessageLevelError) {

	var err error
	var objectId uuid.UUID

	if objectId, err = uuid.Parse(m.Descriptor_.ObjectId); err != nil {
		return "", NewMessageLevelError(400, improperlyConstructedErrorMessage, err)
	}

	//todo: process the request
	fmt.Printf("request -> objectId: %s", objectId.String())

	return "", nil
}
