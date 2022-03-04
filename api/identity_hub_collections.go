package api

import (
	"context"

	. "github.com/getzion/relay/gen/proto/identityhub/v1"
)

func CollectionsQuery(ctx context.Context, r *Request) (string, error) {

	/*

		========================= VALIDATION RULES =========================

		The message object MUST contain a descriptor property, and its value MUST be a JSON object composed as follows:

		The object MUST contain a method property, and its value MUST be the string CollectionsQuery.
		The object MAY contain an objectId property, and if present its value MUST be an [RFC4122] UUID Version 4 string intended to identify a logical object the Hub Instance contains.
		The object MAY contain a schema property, and if present its value MUST be a URI string that indicates the schema of the associated data.
		The object MAY contain a dataFormat property, and its value MUST be a string that indicates the format of the data in accordance with its MIME type designation. The most common format is JSON, which is indicated by setting the value of the dataFormat property to application/json.
		The object MAY contain a dateSort field, and if present its value MUST be one of the following strings:

		createdAscending: return results in order from the earliest dateCreated value to the latest.
		createdDescending: return results in order from the latest dateCreated value to the earliest.
		publishedAscending: return results in order from the earliest datePublished value to the latest.
		publishedDescending: return results in order from the latest datePublished value to the earliest.

	*/

	return "", nil
}

func CollectionsWrite(ctx context.Context, r *Request) (string, error) {

	/*

		========================= VALIDATION RULES =========================

		The message object MUST contain a descriptor property, and its value MUST be a JSON object composed as follows:

		The object MUST contain a method property, and its value MUST be the string CollectionsWrite.
		The object MUST contain an objectId property, and its value MUST be an [RFC4122] UUID Version 4 string.
		The object MAY contain a schema property, and if present its value MUST be a URI string that indicates the schema of the associated data.
		The object MUST contain a dateCreated property, and its value MUST be an Unix epoch timestamp that MUST be set and interpreted as the time the logical entry was created by the DID owner or another permitted party.
		The object MAY contain a datePublished property, and its value MUST be an Unix epoch timestamp that MUST be set and interpreted as the time the logical entry was published by the DID owner or another permitted party.

		Processing Instructions
		When processing a CollectionsWrite message, Hub instances MUST perform the following additional steps:

		If the incoming message has a higher dateCreated value than all of the other messages for the logical entry known to the Hub Instance, the message MUST be designated as the latest state of the logical entry and fully replace all previous messages for the entry.
		If the incoming message has a lower dateCreated value than the message that represents the current state of the logical entry, the message MUST NOT be applied to the logical entry and its data MAY be discarded.
		If the incoming message has a dateCreated value equal to the message that represents the current state of the logical entry, the incoming message’s IPFS CID and the IPFS CID of the message that represents the current state must be lexicographically compared and handled as follows:
		If the incoming message has a higher lexicographic value than the message that represents the current state, perform the actions described in Step 1 of this instruction set.
		If the incoming message has a lower lexicographic value than the message that represents the current state, perform the actions described in Step 2 of this instruction set.

	*/

	return "", nil
}

func CollectionsCommit(ctx context.Context, r *Request) (string, error) {

	/*

		========================= VALIDATION RULES =========================

		The message object MUST descriptor property MUST be a JSON object composed as follows:

		The object MUST contain a method property, and its value MUST be the string CollectionsCommit.
		The object MUST contain an objectId property, and its value MUST be an [RFC4122] UUID Version 4 string.
		The object MAY contain a schema property, and if present its value MUST be a URI string that indicates the schema of the associated data.
		The object MUST contain a dateCreated property, and its value MUST be an Unix epoch timestamp that MUST be set and interpreted as the time the logical entry was created by the DID owner or another permitted party.
		The object MAY contain a datePublished property, and its value MUST be an Unix epoch timestamp that MUST be set and interpreted as the time the logical entry was published by the DID owner or another permitted party.

	*/

	return "", nil
}

func CollectionsDelete(ctx context.Context, r *Request) (string, error) {

	/*
		========================= VALIDATION RULES =========================

		The message object MUST descriptor property MUST be a JSON object composed as follows:

		The object MUST contain a method property, and its value MUST be the string CollectionsDelete.
		The object MUST contain an objectId property, and its value MUST be an [RFC4122] UUID Version 4 string of the object to be deleted.

	*/

	return "", nil
}
