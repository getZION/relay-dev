package api

import (
	. "github.com/getzion/relay/gen/proto/identityhub/v1"
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

	Log.Info().Msg("CollectionsQuery API placeholder")

	return &CollectionsQueryResponse{}, nil
}

func (s *CollectionsService) CollectionsWrite(ctx context.Context, q *CollectionsWriteRequest) (*CollectionsWriteResponse, error) {
	grpc.SendHeader(ctx, metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-unary"))
	grpc.SetTrailer(ctx, metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-unary"))

	Log.Info().Msg("CollectionsWrite API placeholder")

	return &CollectionsWriteResponse{}, nil
}
