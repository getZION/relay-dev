package api

import (
	"log"

	. "github.com/getzion/relay/gen/proto/identityhub/v1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	THREADS_QUERY  = "ThreadsQuery"
	THREADS_CREATE = "ThreadsCreate"
	THREADS_REPLY  = "ThreadsReply"
	THREADS_CLOSE  = "ThreadsClose"
	THREADS_DELETE = "ThreadsDelete"
)

var _ = Describe("IdentityHub Threads", func() {
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

	Context("Message Level", func() {

		Context("Query Tests", func() {

			It("receives a response if a Message Descriptor has valid", func() {
				request := &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Descriptor_: &MessageDescriptor{
								Method: THREADS_QUERY,
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

		Context("Create Tests", func() {

			It("receives an error if a Message Descriptor has missing objectID", func() {
				request := &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Descriptor_: &MessageDescriptor{
								Method: THREADS_CREATE,
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

			It("receives an error if a Message Descriptor has invalid objectID", func() {
				request := &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Descriptor_: &MessageDescriptor{
								Method:   THREADS_CREATE,
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

			It("receives an error if a Message Descriptor has missing schema", func() {
				request := &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Descriptor_: &MessageDescriptor{
								Method:   THREADS_CREATE,
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
				Expect(response.Replies[0].Status.Code).To(Equal(int64(400)))
			})

			It("receives an error if a Message Descriptor has invalid schema", func() {
				request := &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Descriptor_: &MessageDescriptor{
								Method:   THREADS_CREATE,
								ObjectId: OBJECT_ID,
								Schema:   INVALID,
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
								Method:   THREADS_CREATE,
								ObjectId: OBJECT_ID,
								Schema:   SCHEMA,
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

		Context("Reply Tests", func() {

			It("receives an error if a Message Descriptor has missing objectID", func() {
				request := &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Descriptor_: &MessageDescriptor{
								Method: THREADS_REPLY,
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

			It("receives an error if a Message Descriptor has invalid objectID", func() {
				request := &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Descriptor_: &MessageDescriptor{
								Method:   THREADS_REPLY,
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

			It("receives an error if a Message Descriptor has missing schema", func() {
				request := &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Descriptor_: &MessageDescriptor{
								Method:   THREADS_REPLY,
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
				Expect(response.Replies[0].Status.Code).To(Equal(int64(400)))
			})

			It("receives an error if a Message Descriptor has invalid schema", func() {
				request := &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Descriptor_: &MessageDescriptor{
								Method:   THREADS_REPLY,
								ObjectId: OBJECT_ID,
								Schema:   INVALID,
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

			It("receives an error if a Message Descriptor has missing root", func() {
				request := &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Descriptor_: &MessageDescriptor{
								Method:   THREADS_REPLY,
								ObjectId: OBJECT_ID,
								Schema:   SCHEMA,
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

			It("receives an error if a Message Descriptor has invalid root", func() {
				request := &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Descriptor_: &MessageDescriptor{
								Method:   THREADS_REPLY,
								ObjectId: OBJECT_ID,
								Schema:   SCHEMA,
								Root:     INVALID,
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

			It("receives an error if a Message Descriptor has missing parent", func() {
				request := &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Descriptor_: &MessageDescriptor{
								Method:   THREADS_REPLY,
								ObjectId: OBJECT_ID,
								Schema:   SCHEMA,
								Root:     ROOT,
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

			It("receives an error if a Message Descriptor has invalid parent", func() {
				request := &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Descriptor_: &MessageDescriptor{
								Method:   THREADS_REPLY,
								ObjectId: OBJECT_ID,
								Schema:   SCHEMA,
								Root:     ROOT,
								Parent:   INVALID,
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
								Method:   THREADS_REPLY,
								ObjectId: OBJECT_ID,
								Schema:   SCHEMA,
								Root:     ROOT,
								Parent:   PARENT,
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

		Context("Close Tests", func() {

			It("receives an error if a Message Descriptor has missing root", func() {
				request := &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Descriptor_: &MessageDescriptor{
								Method: THREADS_CLOSE,
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

			It("receives an error if a Message Descriptor has invalid root", func() {
				request := &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Descriptor_: &MessageDescriptor{
								Method: THREADS_CLOSE,
								Root:   INVALID,
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
								Method: THREADS_CLOSE,
								Root:   ROOT,
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

		Context("Delete Tests", func() {

			It("receives an error if a Message Descriptor has missing root", func() {
				request := &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Descriptor_: &MessageDescriptor{
								Method: THREADS_DELETE,
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

			It("receives an error if a Message Descriptor has invalid root", func() {
				request := &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Descriptor_: &MessageDescriptor{
								Method: THREADS_DELETE,
								Root:   INVALID,
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
								Method: THREADS_DELETE,
								Root:   ROOT,
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