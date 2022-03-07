package api

import (
	"context"
	"fmt"
	"net/url"

	. "github.com/getzion/relay/gen/proto/identityhub/v1"
	"github.com/google/uuid"
)

func ThreadsQuery(ctx context.Context, m *Message) (string, *MessageLevelError) {
	return "", nil
}

func ThreadsCreate(ctx context.Context, m *Message) (string, *MessageLevelError) {

	var err error
	var objectId uuid.UUID
	var schema *url.URL

	if objectId, err = uuid.Parse(m.Descriptor_.ObjectId); err != nil {
		return "", NewMessageLevelError(400, improperlyConstructedErrorMessage, err)
	} else if schema, err = url.ParseRequestURI(m.Descriptor_.Schema); err != nil {
		return "", NewMessageLevelError(400, improperlyConstructedErrorMessage, err)
	}

	fmt.Printf("request -> objectId: %s, schema: %s", objectId.String(), schema.String())

	return "", nil
}

func ThreadsReply(ctx context.Context, m *Message) (string, *MessageLevelError) {

	var err error
	var objectId uuid.UUID
	var root uuid.UUID
	var parent uuid.UUID
	var schema *url.URL

	if objectId, err = uuid.Parse(m.Descriptor_.ObjectId); err != nil {
		return "", NewMessageLevelError(400, improperlyConstructedErrorMessage, err)
	} else if schema, err = url.ParseRequestURI(m.Descriptor_.Schema); err != nil {
		return "", NewMessageLevelError(400, improperlyConstructedErrorMessage, err)
	} else if root, err = uuid.Parse(m.Descriptor_.Root); err != nil {
		return "", NewMessageLevelError(400, improperlyConstructedErrorMessage, err)
	} else if parent, err = uuid.Parse(m.Descriptor_.Parent); err != nil {
		return "", NewMessageLevelError(400, improperlyConstructedErrorMessage, err)
	}

	fmt.Printf("request -> objectId: %s, schema: %s, root: %s, parent: %s", objectId.String(), schema.String(), root.String(), parent.String())

	return "", nil
}

func ThreadsClose(ctx context.Context, m *Message) (string, *MessageLevelError) {

	var err error
	var root uuid.UUID

	if root, err = uuid.Parse(m.Descriptor_.Root); err != nil {
		return "", NewMessageLevelError(400, improperlyConstructedErrorMessage, err)
	}

	fmt.Printf("request -> root: %s", root.String())

	return "", nil
}

func ThreadsDelete(ctx context.Context, m *Message) (string, *MessageLevelError) {

	var err error
	var root uuid.UUID

	if root, err = uuid.Parse(m.Descriptor_.Root); err != nil {
		return "", NewMessageLevelError(400, improperlyConstructedErrorMessage, err)
	}

	fmt.Printf("request -> root: %s", root.String())

	return "", nil
}
