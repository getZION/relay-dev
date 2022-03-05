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

//todo: need to split files

var _ = Describe("IdentityHub Tests", func() {
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

		Context("Collection Tests", func() {

			Context("CollectionsQuery Tests", func() {

				It("receives a response if a Message Descriptor has missing objectID", func() {
					request := &Request{
						RequestId: REQUEST_ID,
						Target:    TARGET,
						Messages: []*Message{
							{
								Data: "Data!",
								Descriptor_: &MessageDescriptor{
									Method: "CollectionsQuery",
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

				It("receives an error if a Message Descriptor has invalid objectID", func() {
					request := &Request{
						RequestId: REQUEST_ID,
						Target:    TARGET,
						Messages: []*Message{
							{
								Descriptor_: &MessageDescriptor{
									Method:   "CollectionsQuery",
									ObjectId: "<invalid>",
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

				It("receives a response if a Message Descriptor has valid objectID", func() {
					request := &Request{
						RequestId: REQUEST_ID,
						Target:    TARGET,
						Messages: []*Message{
							{
								Descriptor_: &MessageDescriptor{
									Method:   "CollectionsQuery",
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

				It("receives an error if a Message Descriptor has invalid schema", func() {
					request := &Request{
						RequestId: REQUEST_ID,
						Target:    TARGET,
						Messages: []*Message{
							{
								Descriptor_: &MessageDescriptor{
									Method:   "CollectionsQuery",
									ObjectId: OBJECT_ID,
									Schema:   "<invalid>",
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

				It("receives a response if a Message Descriptor has valid schema", func() {
					request := &Request{
						RequestId: REQUEST_ID,
						Target:    TARGET,
						Messages: []*Message{
							{
								Descriptor_: &MessageDescriptor{
									Method:   "CollectionsQuery",
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

				It("receives an error if a Message Descriptor has invalid dataFormat", func() {
					request := &Request{
						RequestId: REQUEST_ID,
						Target:    TARGET,
						Messages: []*Message{
							{
								Descriptor_: &MessageDescriptor{
									Method:     "CollectionsQuery",
									ObjectId:   OBJECT_ID,
									Schema:     SCHEMA,
									DataFormat: "<invalid>",
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

				It("receives a response if a Message Descriptor has valid dataFormat", func() {
					request := &Request{
						RequestId: REQUEST_ID,
						Target:    TARGET,
						Messages: []*Message{
							{
								Descriptor_: &MessageDescriptor{
									Method:     "CollectionsQuery",
									ObjectId:   OBJECT_ID,
									Schema:     SCHEMA,
									DataFormat: DATA_FORMAT,
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

				It("receives an error if a Message Descriptor has invalid dateSort", func() {
					request := &Request{
						RequestId: REQUEST_ID,
						Target:    TARGET,
						Messages: []*Message{
							{
								Descriptor_: &MessageDescriptor{
									Method:     "CollectionsQuery",
									ObjectId:   OBJECT_ID,
									Schema:     SCHEMA,
									DataFormat: DATA_FORMAT,
									DateSort:   "<invalid>",
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

				It("receives a response if a Message Descriptor has valid dateSort", func() {

					//todo: test with other valid dateSort parameters
					request := &Request{
						RequestId: REQUEST_ID,
						Target:    TARGET,
						Messages: []*Message{
							{
								Descriptor_: &MessageDescriptor{
									Method:     "CollectionsQuery",
									ObjectId:   OBJECT_ID,
									Schema:     SCHEMA,
									DataFormat: DATA_FORMAT,
									DateSort:   "createdAscending",
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

			Context("CollectionsWrite Tests", func() {

				It("receives an error if a Message Descriptor has missing objectID", func() {
					request := &Request{
						RequestId: REQUEST_ID,
						Target:    TARGET,
						Messages: []*Message{
							{
								Data: "Data!",
								Descriptor_: &MessageDescriptor{
									Method: "CollectionsWrite",
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
									Method:   "CollectionsWrite",
									ObjectId: "<invalid>",
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

				It("receives an error if a Message Descriptor has missing dateCreated", func() {
					request := &Request{
						RequestId: REQUEST_ID,
						Target:    TARGET,
						Messages: []*Message{
							{
								Descriptor_: &MessageDescriptor{
									Method:   "CollectionsWrite",
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

				It("receives an error if a Message Descriptor has invalid dateCreated", func() {
					request := &Request{
						RequestId: REQUEST_ID,
						Target:    TARGET,
						Messages: []*Message{
							{
								Descriptor_: &MessageDescriptor{
									Method:      "CollectionsWrite",
									ObjectId:    OBJECT_ID,
									DateCreated: "<invalid>",
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

				It("receives a response if a Message Descriptor has valid dateCreated", func() {
					request := &Request{
						RequestId: REQUEST_ID,
						Target:    TARGET,
						Messages: []*Message{
							{
								Descriptor_: &MessageDescriptor{
									Method:      "CollectionsWrite",
									ObjectId:    OBJECT_ID,
									DateCreated: DATE_CREATED,
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

				It("receives an error if a Message Descriptor has invalid schema", func() {
					request := &Request{
						RequestId: REQUEST_ID,
						Target:    TARGET,
						Messages: []*Message{
							{
								Descriptor_: &MessageDescriptor{
									Method:      "CollectionsWrite",
									ObjectId:    OBJECT_ID,
									DateCreated: DATE_CREATED,
									Schema:      "<invalid>",
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

				It("receives an error if a Message Descriptor has valid schema", func() {
					request := &Request{
						RequestId: REQUEST_ID,
						Target:    TARGET,
						Messages: []*Message{
							{
								Descriptor_: &MessageDescriptor{
									Method:      "CollectionsWrite",
									ObjectId:    OBJECT_ID,
									DateCreated: DATE_CREATED,
									Schema:      SCHEMA,
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

				It("receives an error if a Message Descriptor has invalid datePublished", func() {
					request := &Request{
						RequestId: REQUEST_ID,
						Target:    TARGET,
						Messages: []*Message{
							{
								Descriptor_: &MessageDescriptor{
									Method:        "CollectionsWrite",
									ObjectId:      OBJECT_ID,
									DateCreated:   DATE_CREATED,
									Schema:        SCHEMA,
									DatePublished: "<invalid>",
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

				It("receives a response if a Message Descriptor has valid datePublished", func() {
					request := &Request{
						RequestId: REQUEST_ID,
						Target:    TARGET,
						Messages: []*Message{
							{
								Descriptor_: &MessageDescriptor{
									Method:        "CollectionsWrite",
									ObjectId:      OBJECT_ID,
									DateCreated:   DATE_CREATED,
									Schema:        SCHEMA,
									DatePublished: DATE_PUBLISHED,
								},
							},
						},
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
				})

			})

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
