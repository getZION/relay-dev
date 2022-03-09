package nodeinfo

import (
	context "context"

	"github.com/getzion/relay/api/lightning"
	zion "github.com/getzion/relay/gen/proto/zion/v1"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type NodeinfoService struct {
	zion.UnimplementedNodeInfoServiceServer
}

func (s *NodeinfoService) GetNodeInfo(ctx context.Context, q *zion.GetNodeInfoRequest) (*zion.GetNodeInfoResponse, error) {
	grpc.SendHeader(ctx, metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-unary"))
	grpc.SetTrailer(ctx, metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-unary"))

	balance, pubkey := lightning.GetNodeInfo()

	return &zion.GetNodeInfoResponse{
		Balance: balance,
		Pubkey:  pubkey,
	}, nil
}
