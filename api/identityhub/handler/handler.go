package handler

import (
	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/schema"
)

type RequestContext struct {
	SchemaManager *schema.SchemaManager
	Message       *api.Message
}
