package identityhub

import (
	"context"
	"encoding/json"

	"github.com/getzion/relay/api/datastore"
	hub "github.com/getzion/relay/gen/proto/identityhub/v1"
	"github.com/google/uuid"

	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"

	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type (
	interfaceMethodHandler func(ctx context.Context, m *hub.Message) (string, *MessageLevelError)

	IdentityHubService struct {
		hub.UnimplementedHubRequestServiceServer

		validator                *validator.Validate
		prefix                   cid.Prefix
		validHubInterfaceMethods map[string]interfaceMethodHandler

		store *datastore.Store
	}

	MessageLevelError struct {
		Message string
		Code    int64
		Error   error
	}
)

const (
	requestLevelProcessingErrorMessage string = "The request could not be processed correctly"
	targetDIDNotFoundErrorMessage      string = "Target DID not found within the Identity Hub instance"

	improperlyConstructedErrorMessage   string = "The message was malformed or improperly constructed"
	interfaceNotImplementedErrorMessage string = "The interface method is not implemented"
	authorizationFailedErrorMessage     string = "The message failed authorization requirements"

	messageSuccessfullyMessage string = "The message was successfully processed"
)

var (
	//note: https://github.com/ipfs/go-cid maybe we can use this package
	prefix = cid.Prefix{
		Version:  1,
		Codec:    cid.Raw,
		MhType:   multihash.SHA2_256,
		MhLength: -1,
	}

	validHubInterfaceMethods = map[string]interfaceMethodHandler{
		"CollectionsQuery":   CollectionsQuery,
		"CollectionsWrite":   CollectionsWrite,
		"CollectionsCommit":  CollectionsCommit,
		"CollectionsDelete":  CollectionsDelete,
		"ThreadsQuery":       ThreadsQuery,
		"ThreadsCreate":      ThreadsCreate,
		"ThreadsReply":       ThreadsReply,
		"ThreadsClose":       ThreadsClose,
		"ThreadsDelete":      ThreadsDelete,
		"PermissionsRequest": PermissionsRequest,
		"PermissionsQuery":   PermissionsQuery,
		"PermissionsGrant":   PermissionsGrant,
		"PermissionsRevoke":  PermissionsRevoke,
	}
)

func InitIdentityHubService(store *datastore.Store) *IdentityHubService {

	validator := validator.New()

	identityHubService := &IdentityHubService{
		validator:                validator,
		prefix:                   prefix,
		validHubInterfaceMethods: validHubInterfaceMethods,
		store:                    store,
	}

	return identityHubService
}

func (identityHub *IdentityHubService) Process(ctx context.Context, r *hub.Request) (*hub.Response, error) {
	grpc.SendHeader(ctx, metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-unary"))
	grpc.SetTrailer(ctx, metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-unary"))

	response := &hub.Response{
		RequestId: r.RequestId,
		Status: &hub.Status{
			Code: 200,
		},
	}

	if r.RequestId == "" || r.Target == "" || len(r.Messages) == 0 {
		response.Status.Code = 500
		response.Status.Message = requestLevelProcessingErrorMessage
		return response, nil
	} else if _, err := uuid.Parse(r.RequestId); err != nil {
		response.Status.Code = 500
		response.Status.Message = requestLevelProcessingErrorMessage
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
			reply.Status.Message = improperlyConstructedErrorMessage
			response.Replies = append(response.Replies, reply)
			//todo: how we handle internal server error?
			continue
		}

		messageId, err := identityHub.prefix.Sum(messageByte)
		if err != nil {
			reply.Status.Code = 500
			reply.Status.Message = improperlyConstructedErrorMessage
			response.Replies = append(response.Replies, reply)
			//todo: how we handle internal server error?
			continue
		}

		reply.MessageId = messageId.String()

		if message.Descriptor_ == nil || message.Descriptor_.Method == "" {
			reply.Status.Code = 400
			reply.Status.Message = improperlyConstructedErrorMessage
			response.Replies = append(response.Replies, reply)
			continue
		} else if method, ok = identityHub.validHubInterfaceMethods[message.Descriptor_.Method]; !ok {
			reply.Status.Code = 501
			reply.Status.Message = interfaceNotImplementedErrorMessage
			response.Replies = append(response.Replies, reply)
			continue
		}

		entry, mErr := method(ctx, message)
		if mErr != nil {
			reply.Status.Code = mErr.Code
			reply.Status.Message = mErr.Message
			response.Replies = append(response.Replies, reply)
			continue
		}

		if entry != "" {
			reply.Entries = []string{entry}
		}
		reply.Status.Code = 200
		reply.Status.Message = messageSuccessfullyMessage
		response.Replies = append(response.Replies, reply)
	}

	return response, nil
}

func NewMessageLevelError(code int64, message string, err error) *MessageLevelError {
	return &MessageLevelError{
		Code:    code,
		Message: message,
		Error:   err,
	}
}
