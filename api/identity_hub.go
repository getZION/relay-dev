package api

import (
	"errors"
	"strconv"
	"time"

	. "github.com/getzion/relay/gen/proto/identityhub/v1"
	. "github.com/getzion/relay/utils"

	"github.com/go-playground/validator/v10"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func InitIdentityHubService() *IdentityHubService {

	validator := validator.New()
	identityHubService := &IdentityHubService{
		validator: validator,
	}

	return identityHubService
}

type IdentityHubService struct {
	UnimplementedHubRequestServiceServer

	validator *validator.Validate
}

func (hub *IdentityHubService) Process(ctx context.Context, r *Request) (*Response, error) {
	grpc.SendHeader(ctx, metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-unary"))
	grpc.SetTrailer(ctx, metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-unary"))

	//TODO(Umut): To make easier validation, can we use Validator? But first, we need to figure out, how we can add field tags in proto file or how we can make validation without field tags?

	if r == nil {
		return nil, errors.New("Request was empty")
	}

	if r.RequestId == "" {
		return nil, errors.New("Missing Request ID")
	}

	if len(r.RequestId) != 36 {
		return nil, errors.New("Request ID must be len 36 (uuid v4)")
	}

	if r.Target == "" {
		return nil, errors.New("Missing Target")
	}

	if len(r.Messages) == 0 {
		return nil, errors.New("Missing Messages")
	}

	for _, message := range r.Messages {
		if message.Descriptor_ == nil {
			return nil, errors.New("Missing descriptor")
		}

		if message.Descriptor_.Method != "CollectionsWrite" {
			return nil, errors.New("Descriptor method must be CollectionsWrite")
		}

		if len(message.Descriptor_.ObjectId) != 36 {
			return nil, errors.New("Descriptor ObjectID must be length 36 (uuid v4)")
		}

		dateCreated := message.Descriptor_.DateCreated
		if len(dateCreated) < 10 {
			return nil, errors.New("Invalid date")
		}

		dateSeconds, _ := strconv.Atoi(dateCreated)
		t := time.Unix(int64(dateSeconds), 0)
		Log.Info().Msg(t.String())
	}

	return &Response{
		RequestId: r.RequestId,
		Status: &Status{
			Code: 200,
		},
	}, nil
}
