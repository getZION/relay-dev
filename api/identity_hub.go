package api

import (
	"strconv"
	"time"

	. "github.com/getzion/relay/gen/proto/identityhub/v1"
	. "github.com/getzion/relay/utils"
	"github.com/google/uuid"

	//"github.com/ipfs/go-cid"
	//"github.com/multiformats/go-multihash"

	"github.com/go-playground/validator/v10"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type IdentityHubService struct {
	UnimplementedHubRequestServiceServer

	validator                *validator.Validate
	validHubInterfaceMethods map[string]string
}

const (
	requestLevelProcessingErrorMessage string = "The request could not be processed correctly"
	targetDIDNotFoundErrorMessage      string = "Target DID not found within the Identity Hub instance"

	improperlyConstructedErrorMessage   string = "The message was malformed or improperly constructed"
	interfaceNotImplementedErrorMessage string = "The interface method is not implemented"
	authorizationFailedErrorMessage     string = "The message failed authorization requirements"
)

func InitIdentityHubService() *IdentityHubService {

	validator := validator.New()

	identityHubService := &IdentityHubService{
		validator: validator,
		validHubInterfaceMethods: map[string]string{
			"CollectionsQuery":   "",
			"CollectionsWrite":   "",
			"CollectionsCommit":  "",
			"CollectionsDelete":  "",
			"ThreadsQuery":       "",
			"ThreadsCreate":      "",
			"ThreadsReply":       "",
			"ThreadsClose":       "",
			"ThreadsDelete":      "",
			"PermissionsRequest": "",
			"PermissionsQuery":   "",
			"PermissionsGrant":   "",
			"PermissionsRevoke":  "",
		},
	}

	return identityHubService
}

func (hub *IdentityHubService) Process(ctx context.Context, r *Request) (*Response, error) {
	grpc.SendHeader(ctx, metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-unary"))
	grpc.SetTrailer(ctx, metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-unary"))

	if r == nil || r.RequestId == "" || r.Target == "" || len(r.Messages) == 0 {
		return hub.defaultRequestLevelStatusCoding(500, r.RequestId, requestLevelProcessingErrorMessage), nil
	} else if _, err := uuid.Parse(r.RequestId); err != nil {
		return hub.defaultRequestLevelStatusCoding(500, r.RequestId, requestLevelProcessingErrorMessage), nil
	}

	//todo: If the DID targeted by a request object is not found within the Hub instance, the implementation **MUST** produce a request-level status with the code 404, and **SHOULD** use Target DID not found within the Identity Hub instance as the status text

	for _, message := range r.Messages {

		if message.Descriptor_ == nil || message.Descriptor_.Method == "" ||
			len(message.Descriptor_.DateCreated) < 10 ||
			(message.Data != "" && message.Descriptor_.DataFormat == "") {
			return hub.defaultMessageLevelStatusCoding(400, r.RequestId, improperlyConstructedErrorMessage), nil
		} else if _, err := uuid.Parse(message.Descriptor_.ObjectId); err != nil {
			return hub.defaultMessageLevelStatusCoding(400, r.RequestId, improperlyConstructedErrorMessage), nil
		} else if _, ok := hub.validHubInterfaceMethods[message.Descriptor_.Method]; !ok {
			return hub.defaultMessageLevelStatusCoding(501, r.RequestId, interfaceNotImplementedErrorMessage), nil
		}

		dateSeconds, _ := strconv.Atoi(message.Descriptor_.DateCreated)
		t := time.Unix(int64(dateSeconds), 0)
		Log.Info().Msg(t.String())
	}

	//todo: must have messageId which must be the stringified v1 CID of the associated message from its Request object
	/*
		pref := cid.Prefix{
			Version: 1,
			Codec: cid.Raw,
			MhType: multihash.SHA2_256,
			MhLength: -1, // default length
		}
		The object **MUST** have a messageId property,
		and its value **MUST** be the stringified Version 1 CID of the associated message in the Request Object from which it was received.
	*/

	return &Response{
		RequestId: r.RequestId,
		Status: &Status{
			Code: 200,
		},
		Replies: &Reply{
			Status: &Status{
				Code: 200,
			},
		},
	}, nil
}

func (hub *IdentityHubService) defaultRequestLevelStatusCoding(code int64, requestId, message string) *Response {
	return &Response{
		RequestId: requestId,
		Status: &Status{
			Code:    code,
			Message: message,
		},
	}
}

func (hub *IdentityHubService) defaultMessageLevelStatusCoding(code int64, requestId, message string) *Response {
	return &Response{
		RequestId: requestId,
		Status: &Status{
			Code: 200,
		},
		Replies: &Reply{
			MessageId: "", //todo: its value **MUST** be the stringified Version 1 CID
			Status: &Status{
				Code:    code,
				Message: message,
			},
		},
	}
}
