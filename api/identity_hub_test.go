package api

import (
	"log"
	"net"

	. "github.com/getzion/relay/gen/proto/identityhub/v1"
	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

var _ = Describe("IdentityHub Request/Message Level Tests", func() {
	var client HubRequestServiceClient
	var ctx context.Context
	var conn *grpc.ClientConn
	var err error
	var validDemoMessages = []*Message{
		{
			Data: `{ "Data": "Test" }`,
			Descriptor_: &MessageDescriptor{
				Method:      "CollectionsWrite",
				ObjectId:    "d82c0026-ed42-4b26-81f3-94805958a75c",
				DateCreated: "1645917431",
				DataFormat:  "application/json",
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
			Expect(err).To(BeNil())
			Expect(response).To(Not(BeNil()))
			Expect(response.RequestId).To(Equal(request.RequestId))
			Expect(response.Status).To(Not(BeNil()))
			Expect(response.Status.Code).To(Equal(int64(200)))
			Expect(response.Replies).To(Not(BeNil()))
			Expect(response.Replies).To(HaveLen(1))
			Expect(response.Replies[0].Status).To(Not(BeNil()))
			Expect(response.Replies[0].Status.Code).To(Equal(int64(200)))
			Expect(response.Replies[0].MessageId).To(Equal("bafkreib53jzxohpjwiu6zetqokcx6ds4wuqjmsiuuvrfvs3jdzbuegckfq"))
		})

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
				Target:   "TheTarget",
				Messages: validDemoMessages,
			}
			response, err := client.Process(ctx, request)
			Expect(err).To(BeNil())
			Expect(response).To(Not(BeNil()))
			Expect(response.Status).To(Not(BeNil()))
			Expect(response.Status.Code).To(Equal(int64(500)))
		})

		It("receives an error if RequestID is not len 36 (uuid v4)", func() {
			request := &Request{
				Target:    "TheTarget",
				Messages:  validDemoMessages,
				RequestId: "shortrequestid",
			}
			response, err := client.Process(ctx, request)
			Expect(err).To(BeNil())
			Expect(response).To(Not(BeNil()))
			Expect(response.Status).To(Not(BeNil()))
			Expect(response.Status.Code).To(Equal(int64(500)))
		})

		It("receives an error if Target is missing", func() {
			request := &Request{
				RequestId: "3eb8ea70-7ea5-4069-a153-cfb0ea682df9",
				Messages:  validDemoMessages,
			}
			response, err := client.Process(ctx, request)
			Expect(err).To(BeNil())
			Expect(response).To(Not(BeNil()))
			Expect(response.Status).To(Not(BeNil()))
			Expect(response.Status.Code).To(Equal(int64(500)))
		})

		It("receives an error if Messages are missing", func() {
			request := &Request{
				RequestId: "3eb8ea70-7ea5-4069-a153-cfb0ea682df9",
				Target:    "atarget",
			}
			response, err := client.Process(ctx, request)
			Expect(err).To(BeNil())
			Expect(response).To(Not(BeNil()))
			Expect(response.Status).To(Not(BeNil()))
			Expect(response.Status.Code).To(Equal(int64(500)))
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
			response, err := client.Process(ctx, request)
			Expect(err).To(BeNil())
			Expect(response).To(Not(BeNil()))
			Expect(response.Replies).To(Not(BeNil()))
			Expect(response.Replies).To(HaveLen(1))
			Expect(response.Replies[0].Status).To(Not(BeNil()))
			Expect(response.Replies[0].Status.Code).To(Equal(int64(400)))
			Expect(response.Replies[0].MessageId).To(Equal("bafkreigapt7ktrjqxijenwe4oan3ggu237567xwknqr7gohs47tnhcbomm"))
		})

		It("receives an error if a Message Descriptor is missing method", func() {
			request := &Request{
				RequestId: "3eb8ea70-7ea5-4069-a153-cfb0ea682df9",
				Target:    "atarget",
				Messages: []*Message{
					{
						Data:        "Data!",
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
			Expect(response.Replies[0].MessageId).To(Equal("bafkreifhqaw33p7f7ew6lxv4gn4yb2pbrzel6jbyuzbq4ufwa33j4majha"))
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
			response, err := client.Process(ctx, request)
			Expect(err).To(BeNil())
			Expect(response).To(Not(BeNil()))
			Expect(response.Replies).To(Not(BeNil()))
			Expect(response.Replies).To(HaveLen(1))
			Expect(response.Replies[0].Status).To(Not(BeNil()))
			Expect(response.Replies[0].Status.Code).To(Equal(int64(400)))
			Expect(response.Replies[0].MessageId).To(Equal("bafkreih5dg7lfywqhztwwim4ttlxnnmu4c6u633l6epgnxbsrozzscgjye"))
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
			response, err := client.Process(ctx, request)
			Expect(err).To(BeNil())
			Expect(response).To(Not(BeNil()))
			Expect(response.Replies).To(Not(BeNil()))
			Expect(response.Replies).To(HaveLen(1))
			Expect(response.Replies[0].Status).To(Not(BeNil()))
			Expect(response.Replies[0].Status.Code).To(Equal(int64(400)))
			Expect(response.Replies[0].MessageId).To(Equal("bafkreift7ckv3vsw4nr5avjpognmj3jhrfgf7gwzchsv5tc3ctqal3cwna"))
		})

		It("receives an error if a Message has data but dataFormat is missing", func() {
			request := &Request{
				RequestId: "3eb8ea70-7ea5-4069-a153-cfb0ea682df9",
				Target:    "atarget",
				Messages: []*Message{
					{
						Data: "Data!",
						Descriptor_: &MessageDescriptor{
							Method:      "CollectionsWrite",
							ObjectId:    "3eb8ea70-7ea5-4069-a153-cfb0ea682df9",
							DateCreated: "1645917369",
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
			Expect(response.Replies[0].Status.Code).To(Equal(int64(400)))
			Expect(response.Replies[0].MessageId).To(Equal("bafkreigvriutejhgmq2pfg2457ggoixt7bippn5imrd5eal3zcbl46fmqa"))
		})

		It("receives an error if a Message Descriptor method is not implemented", func() {
			request := &Request{
				RequestId: "3eb8ea70-7ea5-4069-a153-cfb0ea682df9",
				Target:    "atarget",
				Messages: []*Message{
					{
						Data: `{ "Data": "Test" }`,
						Descriptor_: &MessageDescriptor{
							Method:      "CollectionsWriteTest",
							ObjectId:    "3eb8ea70-7ea5-4069-a153-cfb0ea682df9",
							DateCreated: "1645917369",
							DataFormat:  "application/json",
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
			Expect(response.Replies[0].MessageId).To(Equal("bafkreigixgythzb46m6xnck3zhvbjyye2n7hbkwlgy6kkll4dslsuy3ghy"))
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
			prefix: cid.Prefix{
				Version:  1,
				Codec:    cid.Raw,
				MhType:   multihash.SHA2_256,
				MhLength: -1,
			},
			validHubInterfaceMethods: map[string]interfaceMethodHandler{
				"CollectionsQuery":   func(ctx context.Context, r *Request) (string, error) { return "", nil },
				"CollectionsWrite":   func(ctx context.Context, r *Request) (string, error) { return "", nil },
				"CollectionsCommit":  func(ctx context.Context, r *Request) (string, error) { return "", nil },
				"CollectionsDelete":  func(ctx context.Context, r *Request) (string, error) { return "", nil },
				"ThreadsQuery":       func(ctx context.Context, r *Request) (string, error) { return "", nil },
				"ThreadsCreate":      func(ctx context.Context, r *Request) (string, error) { return "", nil },
				"ThreadsReply":       func(ctx context.Context, r *Request) (string, error) { return "", nil },
				"ThreadsClose":       func(ctx context.Context, r *Request) (string, error) { return "", nil },
				"ThreadsDelete":      func(ctx context.Context, r *Request) (string, error) { return "", nil },
				"PermissionsRequest": func(ctx context.Context, r *Request) (string, error) { return "", nil },
				"PermissionsQuery":   func(ctx context.Context, r *Request) (string, error) { return "", nil },
				"PermissionsGrant":   func(ctx context.Context, r *Request) (string, error) { return "", nil },
				"PermissionsRevoke":  func(ctx context.Context, r *Request) (string, error) { return "", nil },
			},
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
