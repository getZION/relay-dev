package api

import (
	context "context"

	. "github.com/getzion/relay/gen/proto/zion/v1"
	"github.com/getzion/relay/lightning"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type NodeinfoService struct {
	UnimplementedNodeInfoServiceServer
}

func (s *NodeinfoService) GetNodeInfo(ctx context.Context, q *GetNodeInfoRequest) (*GetNodeInfoResponse, error) {
	grpc.SendHeader(ctx, metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-unary"))
	grpc.SetTrailer(ctx, metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-unary"))

	balance, pubkey := lightning.GetNodeInfo()

	return &GetNodeInfoResponse{
		Balance: balance,
		Pubkey:  pubkey,
	}, nil
}
