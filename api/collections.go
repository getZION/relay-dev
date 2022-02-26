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

	if q.Request.Target == "" {
		return nil, errors.New("Missing Target")
	}

	if len(q.Request.Messages) == 0 {
		return nil, errors.New("Missing Messages")
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
