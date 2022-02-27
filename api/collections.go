package api

import (
	"errors"
	"strconv"
	"time"

	. "github.com/getzion/relay/gen/proto/identityhub/v1"
	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"

	. "github.com/getzion/relay/utils"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type CollectionsService struct {
	UnimplementedCollectionsServiceServer
}

func (s *CollectionsService) CollectionsQuery(ctx context.Context, q *CollectionsQueryRequest) (*CollectionsQueryResponse, error) {
	grpc.SendHeader(ctx, metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-unary"))
	grpc.SetTrailer(ctx, metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-unary"))

	return &CollectionsQueryResponse{}, nil
}

func (s *CollectionsService) CollectionsWrite(ctx context.Context, q *CollectionsWriteRequest) (*CollectionsWriteResponse, error) {
	grpc.SendHeader(ctx, metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-unary"))
	grpc.SetTrailer(ctx, metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-unary"))

	if q.Request == nil {
		return nil, errors.New("Request was empty")
	}

	if q.Request.RequestId == "" {
		return nil, errors.New("Missing Request ID")
	}

	if len(q.Request.RequestId) != 36 {
		return nil, errors.New("Request ID must be len 36 (uuid v4)")
	}

	if q.Request.Target == "" {
		return nil, errors.New("Missing Target")
	}

	if len(q.Request.Messages) == 0 {
		return nil, errors.New("Missing Messages")
	}

	for _, message := range q.Request.Messages {
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

		// Create a cid manually by specifying the 'prefix' parameters
		pref := cid.Prefix{
			Version:  1,
			Codec:    cid.Raw,
			MhType:   multihash.SHA2_256,
			MhLength: -1, // default length
		}
		c, err := pref.Sum([]byte(message.Data))
		if err != nil {
			Log.Error().Msg(err.Error())
		} else {
			Log.Info().Str("cid", c.String()).Msg("Generated CID")
		}
	}

	return &CollectionsWriteResponse{}, nil
}

func (s *CollectionsService) CollectionsCommit(ctx context.Context, q *CollectionsCommitRequest) (*CollectionsCommitResponse, error) {
	grpc.SendHeader(ctx, metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-unary"))
	grpc.SetTrailer(ctx, metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-unary"))

	return &CollectionsCommitResponse{}, nil
}

func (s *CollectionsService) CollectionsDelete(ctx context.Context, q *CollectionsDeleteRequest) (*CollectionsDeleteResponse, error) {
	grpc.SendHeader(ctx, metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-unary"))
	grpc.SetTrailer(ctx, metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-unary"))

	return &CollectionsDeleteResponse{}, nil
}
