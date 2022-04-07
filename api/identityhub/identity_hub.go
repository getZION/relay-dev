package identityhub

import (
	"encoding/json"

	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/constants"
	"github.com/getzion/relay/api/identityhub/errors"
	"github.com/getzion/relay/api/identityhub/handler"
	"github.com/getzion/relay/api/identityhub/handler/collections"
	"github.com/getzion/relay/api/identityhub/handler/permissions"
	"github.com/getzion/relay/api/identityhub/handler/threads"
	"github.com/getzion/relay/api/schema"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
)

type (
	//todo: wrap request with a handler?
	interfaceMethodHandler func(handler *handler.RequestContext) ([]string, *errors.MessageLevelError)

	IdentityHubServer struct {
		prefix                   cid.Prefix
		validHubInterfaceMethods map[string]interfaceMethodHandler
		schemaManager            *schema.SchemaManager
		app                      *fiber.App
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

func InitIdentityHubServer(schemaManager *schema.SchemaManager) *IdentityHubServer {
	identityHubServer := &IdentityHubServer{
		prefix:                   prefix,
		validHubInterfaceMethods: validHubInterfaceMethods,
		schemaManager:            schemaManager,
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, e error) error {
			c.Response().Header.Set("Content-Type", "application/json")

			response := api.Response{
				RequestId: "",
				Status: &api.Status{
					Code:    500,
					Message: "Something goes wrong, please try again.",
				},
			}

			if err, ok := e.(*api.ErrorHandler); ok {
				response.RequestId = err.RequestId
				response.Status.Message = err.Message
				response.Status.Code = err.StatusCode
				if err.Err != nil {
					logrus.Error(err.Err)
				}
			}

			return c.Status(response.Status.Code).JSON(response)
		},
	})
	app.Post("/process", identityHubServer.Process)
	identityHubServer.app = app

	return identityHubServer
}
func (identityHub *IdentityHubServer) Listen(addr string) error {
	return identityHub.app.Listen(addr)
}

func (identityHub *IdentityHubServer) Process(ctx *fiber.Ctx) error {

	request := api.Request{}
	if err := ctx.BodyParser(&request); err != nil {
		ctx.SendStatus(fiber.StatusBadRequest)
		return &api.ErrorHandler{StatusCode: fiber.StatusBadRequest, Message: "request body must be json encoded data", Err: err}
	} else if request.RequestId == "" || request.Target == "" || len(request.Messages) == 0 {
		logrus.Infof("request failed for requestId: %s, target: %s, messages: %d", request.RequestId, request.Target, len(request.Messages))
		return &api.ErrorHandler{StatusCode: fiber.StatusBadRequest, Message: errors.RequestLevelProcessingErrorMessage, RequestId: request.RequestId}
	} else if _, err := uuid.Parse(request.RequestId); err != nil {
		logrus.Infof("requestId must be uuid: %s", request.RequestId)
		return &api.ErrorHandler{StatusCode: fiber.StatusBadRequest, Message: errors.RequestLevelProcessingErrorMessage, RequestId: request.RequestId, Err: err}
	}

	//todo: If the DID targeted by a request object is not found within the Hub instance, the implementation **MUST** produce a request-level status with the code 404, and **SHOULD** use Target DID not found within the Identity Hub instance as the status text
	response := &api.Response{
		RequestId: request.RequestId,
		Status: &api.Status{
			Code: 200,
		},
		Replies: []*api.Reply{},
	}

	var method interfaceMethodHandler
	var ok bool
	for _, message := range request.Messages {

		reply := &api.Reply{
			Status: &api.Status{},
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

		if message.Descriptor == nil || message.Descriptor.Method == "" {
			reply.Status.Code = 400
			reply.Status.Message = errors.ImproperlyConstructedErrorMessage
			response.Replies = append(response.Replies, reply)
			logrus.Info("request message descriptor or method cannot be null or empty")
			continue
		} else if method, ok = identityHub.validHubInterfaceMethods[message.Descriptor.Method]; !ok {
			reply.Status.Code = 501
			reply.Status.Message = errors.InterfaceNotImplementedErrorMessage
			response.Replies = append(response.Replies, reply)
			logrus.Infof("interface method is not implemented: %s", message.Descriptor.Method)
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

	ctx.Response().Header.Set("Content-Type", "application/json")
	return ctx.Status(fiber.StatusOK).JSON(response)
}
