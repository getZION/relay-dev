package identityhub

import (
	"context"
	"encoding/hex"
	"encoding/json"

	"github.com/decred/dcrd/dcrec/secp256k1/v4/schnorr"
	"github.com/getzion/relay/api/constants"
	"github.com/getzion/relay/api/identityhub/errors"
	"github.com/getzion/relay/api/identityhub/handler"
	"github.com/getzion/relay/api/identityhub/handler/collections"
	"github.com/getzion/relay/api/identityhub/handler/permissions"
	"github.com/getzion/relay/api/identityhub/handler/threads"
	"github.com/getzion/relay/api/schema"
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

		schemaManager *schema.SchemaManager
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
		constants.COLLECTIONS_QUERY:   collections.CollectionsQuery,
		constants.COLLECTIONS_WRITE:   collections.CollectionsWrite,
		constants.COLLECTIONS_COMMIT:  collections.CollectionsCommit,
		constants.COLLECTIONS_DELETE:  collections.CollectionsDelete,
		constants.THREADS_QUERY:       threads.ThreadsQuery,
		constants.THREADS_CREATE:      threads.ThreadsCreate,
		constants.THREADS_REPLY:       threads.ThreadsReply,
		constants.THREADS_CLOSE:       threads.ThreadsClose,
		constants.THREADS_DELETE:      threads.ThreadsDelete,
		constants.PERMISSIONS_REQUEST: permissions.PermissionsRequest,
		constants.PERMISSIONS_QUERY:   permissions.PermissionsQuery,
		constants.PERMISSIONS_GRANT:   permissions.PermissionsGrant,
		constants.PERMISSIONS_REVOKE:  permissions.PermissionsRevoke,
	}
)

func InitIdentityHubService(schemaManager *schema.SchemaManager) *IdentityHubService {

	identityHubService := &IdentityHubService{
		prefix:                   prefix,
		validHubInterfaceMethods: validHubInterfaceMethods,
		schemaManager:            schemaManager,
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
		logrus.Infof("request failed for requestId: %s, target: %s, messages: %d", r.RequestId, r.Target, len(r.Messages))
		return response, nil
	} else if _, err := uuid.Parse(r.RequestId); err != nil {
		response.Status.Code = 500
		response.Status.Message = errors.RequestLevelProcessingErrorMessage
		logrus.Infof("requestId must be uuid: %s", r.RequestId)
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

		//todo: auth processes
		if message.Descriptor_.Encryption != "JWE" {
			pubKeyBytes, err := hex.DecodeString("02a673638cb9587cb68ea08dbef685c6f2d" + "2a751a8b3c6f2a7e9a4999e6e4bfaf5")
			if err != nil {
				reply.Status.Code = 500
				reply.Status.Message = errors.ImproperlyConstructedErrorMessage
				response.Replies = append(response.Replies, reply)
				logrus.Errorf("publicKeyBytes decoding failed: %v", err)
				continue
			}

			pubKey, err := schnorr.ParsePubKey(pubKeyBytes)
			if err != nil {
				reply.Status.Code = 500
				reply.Status.Message = errors.ImproperlyConstructedErrorMessage
				response.Replies = append(response.Replies, reply)
				logrus.Errorf("publicKey parse failed: %v", err)
				continue
			}

			logrus.Info(pubKey)

			sigBytes, err := hex.DecodeString("970603d8ccd2475b1ff66cfb3ce7e622c59383" +
				"48304c5a7bc2e6015fb98e3b457d4e912fcca6ca87c04390aa5e6e0e613bbbba7ffd" + "6f15bc59f95bbd92ba50f0")
			if err != nil {
				reply.Status.Code = 500
				reply.Status.Message = errors.ImproperlyConstructedErrorMessage
				response.Replies = append(response.Replies, reply)
				logrus.Errorf("signature decoding failed: %v", err)
				continue
			}

			signature, err := schnorr.ParseSignature(sigBytes)
			if err != nil {
				reply.Status.Code = 500
				reply.Status.Message = errors.ImproperlyConstructedErrorMessage
				response.Replies = append(response.Replies, reply)
				logrus.Errorf("signature parse failed: %v", err)
				continue
			}

			logrus.Info(signature)
		}

		messageByte, err := json.Marshal(message)
		if err != nil {
			reply.Status.Code = 500
			reply.Status.Message = errors.ImproperlyConstructedErrorMessage
			response.Replies = append(response.Replies, reply)
			logrus.Errorf("message serialization failed: %v", err)
			continue
		}

		messageId, err := identityHub.prefix.Sum(messageByte)
		if err != nil {
			reply.Status.Code = 500
			reply.Status.Message = errors.ImproperlyConstructedErrorMessage
			response.Replies = append(response.Replies, reply)
			logrus.Errorf("generating messageId failed: %v", err)
			continue
		}

		reply.MessageId = messageId.String()

		if message.Descriptor_ == nil || message.Descriptor_.Method == "" {
			reply.Status.Code = 400
			reply.Status.Message = errors.ImproperlyConstructedErrorMessage
			response.Replies = append(response.Replies, reply)
			logrus.Info("request message descriptor or method cannot be null or empty")
			continue
		} else if method, ok = identityHub.validHubInterfaceMethods[message.Descriptor_.Method]; !ok {
			reply.Status.Code = 501
			reply.Status.Message = errors.InterfaceNotImplementedErrorMessage
			response.Replies = append(response.Replies, reply)
			logrus.Infof("interface method is not implemented: %s", message.Descriptor_.Method)
			continue
		}

		context := handler.RequestContext{
			SchemaManager: identityHub.schemaManager,
			Message:       message,
		}

		entry, mErr := method(&context)
		if mErr != nil {
			reply.Status.Code = mErr.Code
			reply.Status.Message = mErr.Message
			response.Replies = append(response.Replies, reply)
			logrus.Infof("identity hub method execution failed: %v", mErr)
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
