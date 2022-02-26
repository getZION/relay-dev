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
	var validDemoMessages = []*Message{
		{
			Data: "Data!",
			Descriptor_: &MessageDescriptor{
				Method: "CollectionsWrite",
			},
		},
	}

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

	Describe("CollectionsWrite", func() {
		It("receives a response", func() {
			request := &CollectionsWriteRequest{
				Request: &Request{
					RequestId: "Hello",
					Target:    "TheTarget",
					Messages:  validDemoMessages,
				},
			}
			response, err := client.CollectionsWrite(ctx, request)
			if err != nil {
				utils.Log.Error().Msg(err.Error())
			}
			Expect(response).To(Not(BeNil()))
		})

		It("receives an error if Request is missing", func() {
			request := &CollectionsWriteRequest{}
			_, err := client.CollectionsWrite(ctx, request)
			Expect(err).To(Not(BeNil()))
		})

		It("receives an error if RequestID is missing", func() {
			request := &CollectionsWriteRequest{
				Request: &Request{
					Target:   "TheTarget",
					Messages: validDemoMessages,
				},
			}
			_, err := client.CollectionsWrite(ctx, request)
			Expect(err).To(Not(BeNil()))
		})

		It("receives an error if Target is missing", func() {
			request := &CollectionsWriteRequest{
				Request: &Request{
					RequestId: "09j23f09j23f0j",
					Messages:  validDemoMessages,
				},
			}
			_, err := client.CollectionsWrite(ctx, request)
			Expect(err).To(Not(BeNil()))
		})

		It("receives an error if Messages are missing", func() {
			request := &CollectionsWriteRequest{
				Request: &Request{
					RequestId: "09j23f09j23f0j",
					Target:    "atarget",
				},
			}
			_, err := client.CollectionsWrite(ctx, request)
			Expect(err).To(Not(BeNil()))
		})

		It("receives an error if a Message is missing Descriptor", func() {
			request := &CollectionsWriteRequest{
				Request: &Request{
					RequestId: "09j23f09j23f0j",
					Target:    "atarget",
					Messages: []*Message{
						{
							Data: "Data!",
						},
					},
				},
			}
			_, err := client.CollectionsWrite(ctx, request)
			Expect(err).To(Not(BeNil()))
		})

		It("receives an error if a Message Descriptor is missing method", func() {
			request := &CollectionsWriteRequest{
				Request: &Request{
					RequestId: "09j23f09j23f0j",
					Target:    "atarget",
					Messages: []*Message{
						{
							Data:        "Data!",
							Descriptor_: &MessageDescriptor{},
						},
					},
				},
			}
			_, err := client.CollectionsWrite(ctx, request)
			Expect(err).To(Not(BeNil()))
		})
	})

	Describe("CollectionsQuery", func() {
		It("receives a response", func() {
			request := &CollectionsQueryRequest{}
			response, err := client.CollectionsQuery(ctx, request)
			if err != nil {
				utils.Log.Error().Msg(err.Error())
			}
			Expect(response).To(Not(BeNil()))
		})

		// Request RequestID must be a uuid
		// Every Message object must contain a descriptor property which includes:
		// -- a method property which must have the string value "CollectionsWrite"
		// -- must contain an objectID which must be UUID v4
		// -- may contain schema, if present it must be a URI string
		// -- must contain dateCreated and must be a unix epoch timestamp
	})

	Describe("CollectionsCommit", func() {
		It("receives a response", func() {
			request := &CollectionsCommitRequest{}
			response, err := client.CollectionsCommit(ctx, request)
			if err != nil {
				utils.Log.Error().Msg(err.Error())
			}
			Expect(response).To(Not(BeNil()))
		})
	})

	Describe("CollectionsDelete", func() {
		It("receives a response", func() {
			request := &CollectionsDeleteRequest{}
			response, err := client.CollectionsDelete(ctx, request)
			if err != nil {
				utils.Log.Error().Msg(err.Error())
			}
			Expect(response).To(Not(BeNil()))
		})
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
