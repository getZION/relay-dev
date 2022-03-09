package identityhub

import (
	"log"

	. "github.com/getzion/relay/gen/proto/identityhub/v1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	PERMISSIONS_REQUEST = "PermissionsRequest"
	PERMISSIONS_QUERY   = "PermissionsQuery"
	PERMISSIONS_GRANT   = "PermissionsGrant"
	PERMISSIONS_REVOKE  = "PermissionsRevoke"
)

var _ = Describe("IdentityHub Permissions", func() {
	var client HubRequestServiceClient
	var ctx context.Context
	var conn *grpc.ClientConn
	var err error

	BeforeEach(func() {
		ctx = context.Background()
		conn, err = grpc.DialContext(ctx, "", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithContextDialer(dialer()))
		if err != nil {
			log.Fatal(err)
		}
		client = NewHubRequestServiceClient(conn)
	})

	AfterEach(func() {
		defer conn.Close()
	})

	Context("Message Level", func() {

		Context("Request Tests", func() {

			It("receives an error if a Message Descriptor has missing schema", func() {
				request := &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Descriptor_: &MessageDescriptor{
								Method: PERMISSIONS_REQUEST,
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
			})

			It("receives an error if a Message Descriptor has invalid schema", func() {
				request := &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Descriptor_: &MessageDescriptor{
								Method: PERMISSIONS_REQUEST,
								Schema: INVALID,
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
			})

			It("receives a response if a Message Descriptor has valid", func() {
				request := &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Descriptor_: &MessageDescriptor{
								Method: PERMISSIONS_REQUEST,
								Schema: SCHEMA,
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
				Expect(response.Replies[0].Status.Code).To(Equal(int64(200)))
			})

		})

		Context("Query Tests", func() {

			It("receives an error if a Message Descriptor has missing schema", func() {
				request := &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Descriptor_: &MessageDescriptor{
								Method: PERMISSIONS_QUERY,
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
			})

			It("receives an error if a Message Descriptor has invalid schema", func() {
				request := &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Descriptor_: &MessageDescriptor{
								Method: PERMISSIONS_QUERY,
								Schema: INVALID,
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
			})

			It("receives a response if a Message Descriptor has valid", func() {
				request := &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Descriptor_: &MessageDescriptor{
								Method: PERMISSIONS_QUERY,
								Schema: SCHEMA,
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
				Expect(response.Replies[0].Status.Code).To(Equal(int64(200)))
			})

		})

		Context("Grant Tests", func() {

			It("receives an error if a Message Descriptor has missing schema", func() {
				request := &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Descriptor_: &MessageDescriptor{
								Method: PERMISSIONS_GRANT,
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
			})

			It("receives an error if a Message Descriptor has invalid schema", func() {
				request := &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Descriptor_: &MessageDescriptor{
								Method: PERMISSIONS_GRANT,
								Schema: INVALID,
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
			})

			It("receives a response if a Message Descriptor has valid", func() {
				request := &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Descriptor_: &MessageDescriptor{
								Method: PERMISSIONS_GRANT,
								Schema: SCHEMA,
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
				Expect(response.Replies[0].Status.Code).To(Equal(int64(200)))
			})

		})

		Context("Revoke Tests", func() {

			It("receives an error if a Message Descriptor has missing ObjectId", func() {
				request := &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Descriptor_: &MessageDescriptor{
								Method: PERMISSIONS_REVOKE,
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
			})

			It("receives an error if a Message Descriptor has invalid ObjectId", func() {
				request := &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Descriptor_: &MessageDescriptor{
								Method:   PERMISSIONS_REVOKE,
								ObjectId: INVALID,
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
			})

			It("receives a response if a Message Descriptor has valid", func() {
				request := &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Descriptor_: &MessageDescriptor{
								Method:   PERMISSIONS_REVOKE,
								ObjectId: OBJECT_ID,
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
				Expect(response.Replies[0].Status.Code).To(Equal(int64(200)))
			})

		})

	})

})
