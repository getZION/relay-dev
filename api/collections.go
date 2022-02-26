package api

import (
	"errors"

	. "github.com/getzion/relay/gen/proto/identityhub/v1"
	// . "github.com/getzion/relay/utils"
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

	// Log.Info().Msg("CollectionsQuery API placeholder")

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
		// Log.Info().Msg(message.Data)
		if message.Descriptor_ == nil {
			return nil, errors.New("Missing descriptor")
		}
		// Log.Info().Msg(string(len(message.Descriptor_.Method)))
		// Log.Info().Msg(string(len(message.Descriptor_.Method)))
		if message.Descriptor_.Method != "CollectionsWrite" {
			return nil, errors.New("Descriptor method must be CollectionsWrite")
		}

		if len(message.Descriptor_.ObjectId) != 36 {
			return nil, errors.New("Descriptor ObjectID must be length 36 (uuid v4)")
		}
	}

	return &CollectionsWriteResponse{}, nil
}

func (s *CollectionsService) CollectionsCommit(ctx context.Context, q *CollectionsCommitRequest) (*CollectionsCommitResponse, error) {
	grpc.SendHeader(ctx, metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-unary"))
	grpc.SetTrailer(ctx, metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-unary"))

	// Log.Info().Msg("CollectionsCommit API placeholder")

	return &CollectionsCommitResponse{}, nil
}

func (s *CollectionsService) CollectionsDelete(ctx context.Context, q *CollectionsDeleteRequest) (*CollectionsDeleteResponse, error) {
	grpc.SendHeader(ctx, metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-unary"))
	grpc.SetTrailer(ctx, metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-unary"))

	// Log.Info().Msg("CollectionsDelete API placeholder")

	return &CollectionsDeleteResponse{}, nil
}
