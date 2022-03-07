package api

import (
	"context"

	. "github.com/getzion/relay/gen/proto/identityhub/v1"
)

func PermissionsRequest(ctx context.Context, m *Message) (string, *MessageLevelError) {
	return "", nil
}

func PermissionsQuery(ctx context.Context, m *Message) (string, *MessageLevelError) {
	return "", nil
}

func PermissionsGrant(ctx context.Context, m *Message) (string, *MessageLevelError) {
	return "", nil
}

func PermissionsRevoke(ctx context.Context, m *Message) (string, *MessageLevelError) {
	return "", nil
}
