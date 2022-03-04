package api

import (
	"context"

	. "github.com/getzion/relay/gen/proto/identityhub/v1"
)

func ThreadsQuery(ctx context.Context, r *Request) (string, error) {
	return "", nil
}

func ThreadsCreate(ctx context.Context, r *Request) (string, error) {

	/*

		========================= VALIDATION RULES =========================

		The message object MUST descriptor property MUST be a JSON object composed as follows:

		The object MUST contain a method property, and its value MUST be the string ThreadsCreate.
		The object MUST contain an objectId property, and its value MUST be an [RFC4122] UUID Version 4 string for the Thread being created.
		The object MUST contain a schema property, and its value MUST be a URI string that indicates the schema of the associated data.

	*/

	return "", nil
}

func ThreadsReply(ctx context.Context, r *Request) (string, error) {

	/*

		========================= VALIDATION RULES =========================

		The message object MUST descriptor property MUST be a JSON object composed as follows:

		The object MUST contain a method property, and its value MUST be the string ThreadsReply.
		The object MUST contain an objectId property, and its value MUST be an [RFC4122] UUID Version 4 string representing the reply object.
		The object MUST contain a schema property, and its value MUST be a URI string that indicates the schema of the associated data.
		The object MUST contain a root property, and its value MUST be an [RFC4122] UUID Version 4 string of the initiating Thread message.
		The object MUST contain a parent property, and its value MUST be an [RFC4122] UUID Version 4 string of the message in the Thread being replied to.

	*/

	return "", nil
}

func ThreadsClose(ctx context.Context, r *Request) (string, error) {

	/*

		========================= VALIDATION RULES =========================

		The message object MUST descriptor property MUST be a JSON object composed as follows:

		The object MUST contain a method property, and its value MUST be the string ThreadsClose.
		The object MUST contain a root property, and its value MUST be an [RFC4122] UUID Version 4 string of the initiating Thread message.

	*/

	return "", nil
}

func ThreadsDelete(ctx context.Context, r *Request) (string, error) {

	/*

		========================= VALIDATION RULES =========================

		The message object MUST descriptor property MUST be a JSON object composed as follows:

		The object MUST contain a method property, and its value MUST be the string ThreadsDelete.
		The object MUST contain a root property, and its value MUST be an [RFC4122] UUID Version 4 string of the initiating Thread message.

	*/

	return "", nil
}
