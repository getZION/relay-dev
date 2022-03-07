package api

import (
	"log"
	"net"

	. "github.com/getzion/relay/gen/proto/identityhub/v1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const (
	REQUEST_ID     = "3eb8ea70-7ea5-4069-a153-cfb0ea682df9"
	TARGET         = "TheTarget"
	OBJECT_ID      = "d82c0026-ed42-4b26-81f3-94805958a75c"
	DATE_CREATED   = "1645917431"
	DATE_PUBLISHED = "1645917431"
	SCHEMA         = "https://test.com"
	DATA_FORMAT    = "application/json"
)

var _ = Describe("IdentityHub", func() {

	var client HubRequestServiceClient
	var ctx context.Context
	var conn *grpc.ClientConn
	var err error

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

	Context("Request Level Tests", func() {

		It("receives an error if Request is missing", func() {
			request := &Request{}
			response, err := client.Process(ctx, request)
			Expect(err).To(BeNil())
			Expect(response).To(Not(BeNil()))
			Expect(response.Status).To(Not(BeNil()))
			Expect(response.Status.Code).To(Equal(int64(500)))
		})

		It("receives an error if RequestID is missing", func() {
			request := &Request{
				Target: TARGET,
			}
			response, err := client.Process(ctx, request)
			Expect(err).To(BeNil())
			Expect(response).To(Not(BeNil()))
			Expect(response.Status).To(Not(BeNil()))
			Expect(response.Status.Code).To(Equal(int64(500)))
		})

		It("receives an error if RequestID is not len 36 (uuid v4)", func() {
			request := &Request{
				Target:    TARGET,
				RequestId: "<invalid>",
			}
			response, err := client.Process(ctx, request)
			Expect(err).To(BeNil())
			Expect(response).To(Not(BeNil()))
			Expect(response.Status).To(Not(BeNil()))
			Expect(response.Status.Code).To(Equal(int64(500)))
		})

		It("receives an error if Target is missing", func() {
			request := &Request{
				RequestId: REQUEST_ID,
			}
			response, err := client.Process(ctx, request)
			Expect(err).To(BeNil())
			Expect(response).To(Not(BeNil()))
			Expect(response.Status).To(Not(BeNil()))
			Expect(response.Status.Code).To(Equal(int64(500)))
		})

		It("receives an error if Messages are missing", func() {
			request := &Request{
				RequestId: REQUEST_ID,
				Target:    TARGET,
			}
			response, err := client.Process(ctx, request)
			Expect(err).To(BeNil())
			Expect(response).To(Not(BeNil()))
			Expect(response.Status).To(Not(BeNil()))
			Expect(response.Status.Code).To(Equal(int64(500)))
		})

	})

	Context("Message Level Tests", func() {

		It("receives an error if a Message is missing Descriptor", func() {
			request := &Request{
				RequestId: REQUEST_ID,
				Target:    TARGET,
				Messages: []*Message{
					{},
				},
			}
			response, err := client.Process(ctx, request)
			Expect(err).To(BeNil())
			Expect(response).To(Not(BeNil()))
			Expect(response.Replies).To(Not(BeNil()))
			Expect(response.Replies).To(HaveLen(1))
			Expect(response.Replies[0].Status).To(Not(BeNil()))
			Expect(response.Replies[0].Status.Code).To(Equal(int64(400)))
		})

		It("receives an error if a Message Descriptor is missing Method", func() {
			request := &Request{
				RequestId: REQUEST_ID,
				Target:    TARGET,
				Messages: []*Message{
					{
						Descriptor_: &MessageDescriptor{},
					},
				},
			}
			response, err := client.Process(ctx, request)
			Expect(err).To(BeNil())
			Expect(response).To(Not(BeNil()))
			Expect(response.Replies).To(Not(BeNil()))
			Expect(response.Replies).To(HaveLen(1))
			Expect(response.Replies[0].Status).To(Not(BeNil()))
			Expect(response.Replies[0].Status.Code).To(Equal(int64(400)))
		})

		It("receives an error if a Message Descriptor method is not implemented", func() {
			request := &Request{
				RequestId: REQUEST_ID,
				Target:    TARGET,
				Messages: []*Message{
					{
						Descriptor_: &MessageDescriptor{
							Method: "CollectionsQueryTest",
						},
					},
				},
			}
			response, err := client.Process(ctx, request)
			Expect(err).To(BeNil())
			Expect(response).To(Not(BeNil()))
			Expect(response.Replies).To(Not(BeNil()))
			Expect(response.Replies).To(HaveLen(1))
			Expect(response.Replies[0].Status).To(Not(BeNil()))
			Expect(response.Replies[0].Status.Code).To(Equal(int64(501)))
		})

	})
})

type mockIdentityHubServer struct {
	IdentityHubService
}

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	RegisterHubRequestServiceServer(server, &mockIdentityHubServer{
		IdentityHubService: IdentityHubService{
			prefix:                   prefix,
			validHubInterfaceMethods: validHubInterfaceMethods,
		},
	})

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}
