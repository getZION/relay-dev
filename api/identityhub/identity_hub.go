package identityhub

import (
	"context"
	"encoding/json"

	"github.com/getzion/relay/api/datastore"
	"github.com/getzion/relay/api/identityhub/errors"
	"github.com/getzion/relay/api/identityhub/handler"
	"github.com/getzion/relay/api/identityhub/handler/collections"
	"github.com/getzion/relay/api/identityhub/handler/permissions"
	"github.com/getzion/relay/api/identityhub/handler/threads"
	hub "github.com/getzion/relay/gen/proto/identityhub/v1"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
)

type (
	//todo: wrap request with a handler?
	interfaceMethodHandler func(handler *handler.RequestContext) ([]string, *errors.MessageLevelError)

	IdentityHubService struct {
		hub.UnimplementedHubRequestServiceServer

		prefix                   cid.Prefix
		validHubInterfaceMethods map[string]interfaceMethodHandler

		store *datastore.Store
	}
)

var (
	//note: https://github.com/ipfs/go-cid maybe we can use this package
	prefix = cid.Prefix{
		Version:  1,
		Codec:    cid.Raw,
		MhType:   multihash.SHA2_256,
		MhLength: -1,
	}

	//todo: change request handler implemantation
	validHubInterfaceMethods = map[string]interfaceMethodHandler{
		"CollectionsQuery":   collections.CollectionsQuery,
		"CollectionsWrite":   collections.CollectionsWrite,
		"CollectionsCommit":  collections.CollectionsCommit,
		"CollectionsDelete":  collections.CollectionsDelete,
		"ThreadsQuery":       threads.ThreadsQuery,
		"ThreadsCreate":      threads.ThreadsCreate,
		"ThreadsReply":       threads.ThreadsReply,
		"ThreadsClose":       threads.ThreadsClose,
		"ThreadsDelete":      threads.ThreadsDelete,
		"PermissionsRequest": permissions.PermissionsRequest,
		"PermissionsQuery":   permissions.PermissionsQuery,
		"PermissionsGrant":   permissions.PermissionsGrant,
		"PermissionsRevoke":  permissions.PermissionsRevoke,
	}
)

func InitIdentityHubService(store *datastore.Store) *IdentityHubService {

	identityHubService := &IdentityHubService{
		prefix:                   prefix,
		validHubInterfaceMethods: validHubInterfaceMethods,
		store:                    store,
	}

	return identityHubService
}

func (identityHub *IdentityHubService) Process(ctx context.Context, r *hub.Request) (*hub.Response, error) {
	//grpc.SendHeader(ctx, metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-unary"))
	//grpc.SetTrailer(ctx, metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-unary"))

	response := &hub.Response{
		RequestId: r.RequestId,
		Status: &hub.Status{
			Code: 200,
		},
	}

	if r.RequestId == "" || r.Target == "" || len(r.Messages) == 0 {
		response.Status.Code = 500
		response.Status.Message = errors.RequestLevelProcessingErrorMessage
		return response, nil
	} else if _, err := uuid.Parse(r.RequestId); err != nil {
		response.Status.Code = 500
		response.Status.Message = errors.RequestLevelProcessingErrorMessage
		return response, nil
	}

	//todo: If the DID targeted by a request object is not found within the Hub instance, the implementation **MUST** produce a request-level status with the code 404, and **SHOULD** use Target DID not found within the Identity Hub instance as the status text
	response.Replies = []*hub.Reply{}

	var method interfaceMethodHandler
	var ok bool
	for _, message := range r.Messages {

		reply := &hub.Reply{
			Status: &hub.Status{},
		}

		messageByte, err := json.Marshal(message)
		if err != nil {
			reply.Status.Code = 500
			reply.Status.Message = errors.ImproperlyConstructedErrorMessage
			response.Replies = append(response.Replies, reply)
			//todo: how we handle internal server error?
			continue
		}

		messageId, err := identityHub.prefix.Sum(messageByte)
		if err != nil {
			reply.Status.Code = 500
			reply.Status.Message = errors.ImproperlyConstructedErrorMessage
			response.Replies = append(response.Replies, reply)
			//todo: how we handle internal server error?
			continue
		}

		reply.MessageId = messageId.String()

		if message.Descriptor_ == nil || message.Descriptor_.Method == "" {
			reply.Status.Code = 400
			reply.Status.Message = errors.ImproperlyConstructedErrorMessage
			response.Replies = append(response.Replies, reply)
			continue
		} else if method, ok = identityHub.validHubInterfaceMethods[message.Descriptor_.Method]; !ok {
			reply.Status.Code = 501
			reply.Status.Message = errors.InterfaceNotImplementedErrorMessage
			response.Replies = append(response.Replies, reply)
			continue
		}

		context := handler.RequestContext{
			Store:   identityHub.store,
			Message: message,
		}

		entry, mErr := method(&context)
		if mErr != nil {
			reply.Status.Code = mErr.Code
			reply.Status.Message = mErr.Message
			response.Replies = append(response.Replies, reply)
			logrus.Error(mErr.Error)
			continue
		}

		if entry != nil {
			reply.Entries = entry
		}
		reply.Status.Code = 200
		reply.Status.Message = errors.MessageSuccessfullyMessage
		response.Replies = append(response.Replies, reply)
	}

	return response, nil
}
