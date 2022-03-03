package api

import (
	"encoding/json"
	"strconv"
	"time"

	. "github.com/getzion/relay/gen/proto/identityhub/v1"
	. "github.com/getzion/relay/utils"
	"github.com/google/uuid"

	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"

	"github.com/go-playground/validator/v10"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type IdentityHubService struct {
	UnimplementedHubRequestServiceServer

	validator                *validator.Validate
	prefix                   cid.Prefix
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
		prefix: cid.Prefix{
			Version:  1,
			Codec:    cid.Raw,
			MhType:   multihash.SHA2_256,
			MhLength: -1,
		},
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

	exampleBytes, _ := json.Marshal(r)
	messageId, err := hub.prefix.Sum(exampleBytes)
	response := &Response{
		RequestId: r.RequestId,
		Status: &Status{
			Code: 200,
		},
	}

	if err != nil {
		return hub.defaultRequestLevelStatusCoding(response, 500, requestLevelProcessingErrorMessage), nil
	} else if r.RequestId == "" || r.Target == "" || len(r.Messages) == 0 {
		return hub.defaultRequestLevelStatusCoding(response, 500, requestLevelProcessingErrorMessage), nil
	} else if _, err := uuid.Parse(r.RequestId); err != nil {
		return hub.defaultRequestLevelStatusCoding(response, 500, requestLevelProcessingErrorMessage), nil
	}

	//todo: If the DID targeted by a request object is not found within the Hub instance, the implementation **MUST** produce a request-level status with the code 404, and **SHOULD** use Target DID not found within the Identity Hub instance as the status text

	response.Replies = &Reply{
		MessageId: messageId.String(),
		Status: &Status{
			Code: 200,
		},
	}

	for _, message := range r.Messages {

		if message.Descriptor_ == nil || message.Descriptor_.Method == "" ||
			len(message.Descriptor_.DateCreated) < 10 ||
			(message.Data != "" && message.Descriptor_.DataFormat == "") {
			return hub.defaultMessageLevelStatusCoding(response, 400, improperlyConstructedErrorMessage), nil
		} else if _, err := uuid.Parse(message.Descriptor_.ObjectId); err != nil {
			return hub.defaultMessageLevelStatusCoding(response, 400, improperlyConstructedErrorMessage), nil
		} else if _, ok := hub.validHubInterfaceMethods[message.Descriptor_.Method]; !ok {
			return hub.defaultMessageLevelStatusCoding(response, 501, interfaceNotImplementedErrorMessage), nil
		}

		dateSeconds, _ := strconv.Atoi(message.Descriptor_.DateCreated)
		t := time.Unix(int64(dateSeconds), 0)
		Log.Info().Msg(t.String())
	}

	return response, nil
}

func (hub *IdentityHubService) defaultRequestLevelStatusCoding(r *Response, code int64, message string) *Response {
	r.Status.Code = code
	r.Status.Message = message
	return r
}

func (hub *IdentityHubService) defaultMessageLevelStatusCoding(r *Response, code int64, message string) *Response {
	r.Replies.Status.Code = code
	r.Replies.Status.Message = message
	return r
}
