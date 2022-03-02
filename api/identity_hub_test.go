package api

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

var _ = Describe("IdentityHub", func() {
	var client HubRequestServiceClient
	var ctx context.Context
	var conn *grpc.ClientConn
	var err error
	var validDemoMessages = []*Message{
		{
			Data: "Data!",
			Descriptor_: &MessageDescriptor{
				Method:      "CollectionsWrite",
				ObjectId:    "d82c0026-ed42-4b26-81f3-94805958a75c",
				DateCreated: "1645917431",
			},
		},
	}

	BeforeEach(func() {
		ctx = context.Background()
		conn, err = grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
		if err != nil {
			log.Fatal(err)
		}
		client = NewHubRequestServiceClient(conn)
	})

	AfterEach(func() {
		defer conn.Close()
	})

	Describe("request", func() {
		It("receives a response", func() {
			request := &Request{
				RequestId: "3eb8ea70-7ea5-4069-a153-cfb0ea682df9",
				Target:    "TheTarget",
				Messages:  validDemoMessages,
			}
			response, err := client.Process(ctx, request)
			if err != nil {
				utils.Log.Error().Msg(err.Error())
			}
			Expect(response).To(Not(BeNil()))
			Expect(response.RequestId).To(Equal(request.RequestId))
			Expect(response.Status.Code).To(Equal(int64(200)))
		})

		It("receives an error if Request is missing", func() {
			request := &Request{}
			_, err := client.Process(ctx, request)
			Expect(err).To(Not(BeNil()))
		})

		It("receives an error if RequestID is missing", func() {
			request := &Request{
				Target:   "TheTarget",
				Messages: validDemoMessages,
			}
			_, err := client.Process(ctx, request)
			Expect(err).To(Not(BeNil()))
		})

		It("receives an error if RequestID is not len 36 (uuid v4)", func() {
			request := &Request{
				Target:    "TheTarget",
				Messages:  validDemoMessages,
				RequestId: "shortrequestid",
			}
			_, err := client.Process(ctx, request)
			Expect(err).To(Not(BeNil()))
		})

		It("receives an error if Target is missing", func() {
			request := &Request{
				RequestId: "3eb8ea70-7ea5-4069-a153-cfb0ea682df9",
				Messages:  validDemoMessages,
			}
			_, err := client.Process(ctx, request)
			Expect(err).To(Not(BeNil()))
		})

		It("receives an error if Messages are missing", func() {
			request := &Request{
				RequestId: "3eb8ea70-7ea5-4069-a153-cfb0ea682df9",
				Target:    "atarget",
			}
			_, err := client.Process(ctx, request)
			Expect(err).To(Not(BeNil()))
		})

		It("receives an error if a Message is missing Descriptor", func() {
			request := &Request{
				RequestId: "3eb8ea70-7ea5-4069-a153-cfb0ea682df9",
				Target:    "atarget",
				Messages: []*Message{
					{
						Data: "Data!",
					},
				},
			}
			_, err := client.Process(ctx, request)
			Expect(err).To(Not(BeNil()))
		})

		It("receives an error if a Message Descriptor is missing method", func() {
			request := &Request{
				RequestId: "09j23f09j23f0j",
				Target:    "atarget",
				Messages: []*Message{
					{
						Data:        "Data!",
						Descriptor_: &MessageDescriptor{},
					},
				},
			}
			_, err := client.Process(ctx, request)
			Expect(err).To(Not(BeNil()))
		})

		It("receives an error if a Message Descriptor is missing objectID", func() {
			request := &Request{
				RequestId: "3eb8ea70-7ea5-4069-a153-cfb0ea682df9",
				Target:    "atarget",
				Messages: []*Message{
					{
						Data: "Data!",
						Descriptor_: &MessageDescriptor{
							Method:      "CollectionsWrite",
							ObjectId:    "12342135",
							DateCreated: "1645917369",
						},
					},
				},
			}
			_, err := client.Process(ctx, request)
			Expect(err).To(Not(BeNil()))
		})

		It("receives an error if a Message Descriptor is missing dateCreated", func() {
			request := &Request{
				RequestId: "3eb8ea70-7ea5-4069-a153-cfb0ea682df9",
				Target:    "atarget",
				Messages: []*Message{
					{
						Data: "Data!",
						Descriptor_: &MessageDescriptor{
							Method:   "CollectionsWrite",
							ObjectId: "b9b672ba-68a7-46b1-b24d-104a860aafdf",
						},
					},
				},
			}
			_, err := client.Process(ctx, request)
			Expect(err).To(Not(BeNil()))
		})
	})
})

type mockIdentityHubServer struct {
	IdentityHubService
}

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	RegisterHubRequestServiceServer(server, &mockIdentityHubServer{})

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}
