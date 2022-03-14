package handler

import (
	"github.com/getzion/relay/api/datastore"
	hub "github.com/getzion/relay/gen/proto/identityhub/v1"
)

type RequestContext struct {
	Store   *datastore.Store
	Message *hub.Message
}
