package api

import (
	"log"
	"net"

	. "github.com/getzion/relay/gen/proto/identityhub/v1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

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

		It("receives a response", func() {
			request := &Request{
				RequestId: "3eb8ea70-7ea5-4069-a153-cfb0ea682df9",
				Target:    "TheTarget",
				Messages: []*Message{
					{
						Data: `{ "Data": "Test" }`,
						Descriptor_: &MessageDescriptor{
							Method:      "CollectionsWrite",
							ObjectId:    "d82c0026-ed42-4b26-81f3-94805958a75c",
							DateCreated: "1645917431",
							DataFormat:  "application/json",
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
				Target: "TheTarget",
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
				RequestId: "3eb8ea70-7ea5-4069-a153-cfb0ea682df9",
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

	})

	Context("Message Level Tests", func() {

		It("receives an error if a Message is missing Descriptor", func() {
			request := &Request{
				RequestId: "3eb8ea70-7ea5-4069-a153-cfb0ea682df9",
				Target:    "atarget",
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
			Expect(response.Replies[0].MessageId).To(Equal("bafkreicecnx2gvntm6fbcrvnc336qze6st5u7qq7457igegamd3bzkx7ri"))
		})

		It("receives an error if a Message Descriptor is missing Method", func() {
			request := &Request{
				RequestId: "3eb8ea70-7ea5-4069-a153-cfb0ea682df9",
				Target:    "atarget",
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
			Expect(response.Replies[0].MessageId).To(Equal("bafkreibrw36mqjy3ndlmfuolbyr45dgsnoig7kqhtrhtd6nrjmraodeqga"))
		})

		It("receives an error if a Message Descriptor method is not implemented", func() {
			request := &Request{
				RequestId: "3eb8ea70-7ea5-4069-a153-cfb0ea682df9",
				Target:    "atarget",
				Messages: []*Message{
					{
						Descriptor_: &MessageDescriptor{
							Method: "CollectionsWriteTest",
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
			Expect(response.Replies[0].MessageId).To(Equal("bafkreiau3c63kbr35o3wzbmxueh3tpqmp34qqn74t5zfl7naiss5mof4sm"))
		})

		Context("Collection Tests", func() {

			Context("CollectionsQuery Tests", func() {

				It("receives an response if a Message Descriptor has missing objectID", func() {
					request := &Request{
						RequestId: "3eb8ea70-7ea5-4069-a153-cfb0ea682df9",
						Target:    "atarget",
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
					Expect(response.Replies[0].MessageId).To(Equal("bafkreidjwd7nigpmsvz5gfxvfhi2d6o4es37m56gqmwj7qu3ebs4w6krz4"))
				})

				It("receives an error if a Message Descriptor has invalid objectID", func() {
					request := &Request{
						RequestId: "3eb8ea70-7ea5-4069-a153-cfb0ea682df9",
						Target:    "atarget",
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
					Expect(response.Replies[0].MessageId).To(Equal("bafkreic326ktl6w3hwidjznjpmqepakqbncqj46skfp6ur5vbqm6akzbwa"))
				})

				It("receives a response if a Message Descriptor has valid objectID", func() {
					request := &Request{
						RequestId: "3eb8ea70-7ea5-4069-a153-cfb0ea682df9",
						Target:    "atarget",
						Messages: []*Message{
							{
								Descriptor_: &MessageDescriptor{
									Method:   "CollectionsQuery",
									ObjectId: "d82c0026-ed42-4b26-81f3-94805958a75c",
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
					Expect(response.Replies[0].MessageId).To(Equal("bafkreifo63i6zbffdn7gw63zu5n5mwmzjqd2nuk5s3jebyuxfp7ynbffgy"))
				})

				It("receives an error if a Message Descriptor has invalid schema", func() {
					request := &Request{
						RequestId: "3eb8ea70-7ea5-4069-a153-cfb0ea682df9",
						Target:    "atarget",
						Messages: []*Message{
							{
								Descriptor_: &MessageDescriptor{
									Method: "CollectionsQuery",
									Schema: "<invalid>",
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
					Expect(response.Replies[0].MessageId).To(Equal("bafkreig7mrtisjqg5vaw55uxhylopnhkxbpz3xfgzf4l24z35wqbo5on24"))
				})

				It("receives a response if a Message Descriptor has valid schema", func() {
					request := &Request{
						RequestId: "3eb8ea70-7ea5-4069-a153-cfb0ea682df9",
						Target:    "atarget",
						Messages: []*Message{
							{
								Descriptor_: &MessageDescriptor{
									Method: "CollectionsQuery",
									Schema: "https://test.com",
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
					Expect(response.Replies[0].MessageId).To(Equal("bafkreif2hijtph6lxoe4ziujaicssrqxzgnrpauuio44vcncnonygpbdpe"))
				})

				It("receives an error if a Message Descriptor has invalid dataFormat", func() {
					request := &Request{
						RequestId: "3eb8ea70-7ea5-4069-a153-cfb0ea682df9",
						Target:    "atarget",
						Messages: []*Message{
							{
								Descriptor_: &MessageDescriptor{
									Method:     "CollectionsQuery",
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
					Expect(response.Replies[0].MessageId).To(Equal("bafkreihk6asib3lt2vec4zwl2gbljqawfbvpjiiisr3kgeult2drbjdedm"))
				})

				It("receives a response if a Message Descriptor has valid dataFormat", func() {
					request := &Request{
						RequestId: "3eb8ea70-7ea5-4069-a153-cfb0ea682df9",
						Target:    "atarget",
						Messages: []*Message{
							{
								Descriptor_: &MessageDescriptor{
									Method:     "CollectionsQuery",
									DataFormat: "application/json",
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
					Expect(response.Replies[0].MessageId).To(Equal("bafkreifrvxr7mn7lskgsu3yvda4alw2ci5onqieo44gnhgpzj627ugnpti"))
				})

				It("receives an error if a Message Descriptor has invalid dateSort", func() {
					request := &Request{
						RequestId: "3eb8ea70-7ea5-4069-a153-cfb0ea682df9",
						Target:    "atarget",
						Messages: []*Message{
							{
								Descriptor_: &MessageDescriptor{
									Method:   "CollectionsQuery",
									DateSort: "<invalid>",
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
					Expect(response.Replies[0].MessageId).To(Equal("bafkreicb5ymqyxksvqtgsj4xgr77ux2ldgjrw4govmxlpink4k4vmf2hki"))
				})

				It("receives an error if a Message Descriptor has valid dateSort", func() {

					request := &Request{
						RequestId: "3eb8ea70-7ea5-4069-a153-cfb0ea682df9",
						Target:    "atarget",
						Messages: []*Message{
							{
								Descriptor_: &MessageDescriptor{
									Method:   "CollectionsQuery",
									DateSort: "createdAscending",
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
					Expect(response.Replies[0].MessageId).To(Equal("bafkreihbret5lecce2ft5jnal3kfv3kze3mgoje3ejpno4v4debge3ucn4"))
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
