package api

// package api_test

import (
	"log"
	"net"

	// api "github.com/getzion/relay/api"
	. "github.com/getzion/relay/gen/proto/identityhub/v1"
	"github.com/getzion/relay/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

var _ = Describe("Collections", func() {
	var client CollectionsServiceClient
	var ctx context.Context
	var conn *grpc.ClientConn
	var err error

	BeforeEach(func() {
		ctx = context.Background()
		conn, err = grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
		if err != nil {
			log.Fatal(err)
		}
		client = NewCollectionsServiceClient(conn)
	})

	AfterEach(func() {
		defer conn.Close()
	})

	It("receives a response from CollectionsWrite service", func() {
		request := &CollectionsWriteRequest{}
		response, err := client.CollectionsWrite(ctx, request)
		if err != nil {
			utils.Log.Error().Msg(err.Error())
		}
		Expect(response).To(Not(BeNil()))
	})
})

type mockCollectionsServer struct {
	CollectionsService
}

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	RegisterCollectionsServiceServer(server, &mockCollectionsServer{})
	// pb.RegisterDepositServiceServer(server, &mockDepositServer{})

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}
