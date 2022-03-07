package api

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/gabriel-vasile/mimetype"
	. "github.com/getzion/relay/gen/proto/identityhub/v1"
	"github.com/google/uuid"
)

func CollectionsQuery(ctx context.Context, m *Message) (string, *MessageLevelError) {

	/*

		========================= VALIDATION RULES =========================

		The message object MUST contain a descriptor property, and its value MUST be a JSON object composed as follows:

		+ The object MUST contain a method property, and its value MUST be the string CollectionsQuery.
		+ The object MAY contain an objectId property, and if present its value MUST be an [RFC4122] UUID Version 4 string intended to identify a logical object the Hub Instance contains.
		+ The object MAY contain a schema property, and if present its value MUST be a URI string that indicates the schema of the associated data.
		+ The object MAY contain a dataFormat property, and its value MUST be a string that indicates the format of the data in accordance with its MIME type designation. The most common format is JSON, which is indicated by setting the value of the dataFormat property to application/json.
		+ The object MAY contain a dateSort field, and if present its value MUST be one of the following strings:

			createdAscending: return results in order from the earliest dateCreated value to the latest.
			createdDescending: return results in order from the latest dateCreated value to the earliest.
			publishedAscending: return results in order from the earliest datePublished value to the latest.
			publishedDescending: return results in order from the latest datePublished value to the earliest.

	*/

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

	return "", nil
}

func CollectionsWrite(ctx context.Context, m *Message) (string, *MessageLevelError) {

	/*

		========================= VALIDATION RULES =========================

		The message object MUST contain a descriptor property, and its value MUST be a JSON object composed as follows:

		+ The object MUST contain a method property, and its value MUST be the string CollectionsWrite.
		+ The object MUST contain an objectId property, and its value MUST be an [RFC4122] UUID Version 4 string.
		+ The object MUST contain a dateCreated property, and its value MUST be an Unix epoch timestamp that MUST be set and interpreted as the time the logical entry was created by the DID owner or another permitted party.
		+ The object MAY contain a schema property, and if present its value MUST be a URI string that indicates the schema of the associated data.
		+ The object MAY contain a datePublished property, and its value MUST be an Unix epoch timestamp that MUST be set and interpreted as the time the logical entry was published by the DID owner or another permitted party.

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

	return "", nil
}

func CollectionsCommit(ctx context.Context, m *Message) (string, *MessageLevelError) {

	/*

		========================= VALIDATION RULES =========================

		The message object MUST descriptor property MUST be a JSON object composed as follows:

		+ The object MUST contain a method property, and its value MUST be the string CollectionsCommit.
		+ The object MUST contain an objectId property, and its value MUST be an [RFC4122] UUID Version 4 string.
		+ The object MUST contain a dateCreated property, and its value MUST be an Unix epoch timestamp that MUST be set and interpreted as the time the logical entry was created by the DID owner or another permitted party.
		+ The object MAY contain a schema property, and if present its value MUST be a URI string that indicates the schema of the associated data.
		+ The object MAY contain a datePublished property, and its value MUST be an Unix epoch timestamp that MUST be set and interpreted as the time the logical entry was published by the DID owner or another permitted party.

	*/

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

func CollectionsDelete(ctx context.Context, m *Message) (string, *MessageLevelError) {

	/*
		========================= VALIDATION RULES =========================

		The message object MUST descriptor property MUST be a JSON object composed as follows:

		+ The object MUST contain a method property, and its value MUST be the string CollectionsDelete.
		+ The object MUST contain an objectId property, and its value MUST be an [RFC4122] UUID Version 4 string of the object to be deleted.

	*/

	var err error
	var objectId uuid.UUID

	if objectId, err = uuid.Parse(m.Descriptor_.ObjectId); err != nil {
		return "", NewMessageLevelError(400, improperlyConstructedErrorMessage, err)
	}

	//todo: process the request
	fmt.Printf("request -> objectId: %s", objectId.String())

	return "", nil
}
