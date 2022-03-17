package handler

import (
	"github.com/getzion/relay/api/schema"
	hub "github.com/getzion/relay/gen/proto/identityhub/v1"
)

type RequestContext struct {
	SchemaManager *schema.SchemaManager
	Message       *hub.Message
}
