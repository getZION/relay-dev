package identityhub

import (
	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/schema"
	"github.com/getzion/relay/api/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	THREADS_QUERY  = "ThreadsQuery"
	THREADS_CREATE = "ThreadsCreate"
	THREADS_REPLY  = "ThreadsReply"
	THREADS_CLOSE  = "ThreadsClose"
	THREADS_DELETE = "ThreadsDelete"
)

var _ = Describe("IdentityHub Threads", func() {
	var (
		t                GinkgoTestReporter
		gomockController *gomock.Controller
		app              *fiber.App
		st               *storage.MockStorage
	)

	BeforeEach(func() {
		gomockController = gomock.NewController(t)
		st = storage.NewMockStorage(gomockController)
		schemaManager := schema.NewSchemaManager(st)
		server := InitIdentityHubServer(schemaManager, st)
		app = server.app
	})

	AfterEach(func() {
		gomockController.Finish()
	})

	Context("Message Level", func() {

		Context("Query Tests", func() {

			It("receives a response if a Message Descriptor has valid", func() {
				request := &api.Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*api.Message{
						{
							Descriptor: &api.MessageDescriptor{
								Method: THREADS_QUERY,
							},
						},
					},
				}
				response, err := process(app, request)
				Expect(err).To(BeNil())
				Expect(response).To(Not(BeNil()))
				Expect(response.Replies).To(Not(BeNil()))
				Expect(response.Replies).To(HaveLen(1))
				Expect(response.Replies[0].Status).To(Not(BeNil()))
				Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusOK))
			})

		})

		Context("Create Tests", func() {

			It("receives an error if a Message Descriptor has missing objectID", func() {
				request := &api.Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*api.Message{
						{
							Descriptor: &api.MessageDescriptor{
								Method: THREADS_CREATE,
							},
						},
					},
				}
				response, err := process(app, request)
				Expect(err).To(BeNil())
				Expect(response).To(Not(BeNil()))
				Expect(response.Replies).To(Not(BeNil()))
				Expect(response.Replies).To(HaveLen(1))
				Expect(response.Replies[0].Status).To(Not(BeNil()))
				Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusBadRequest))
			})

			It("receives an error if a Message Descriptor has invalid objectID", func() {
				request := &api.Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*api.Message{
						{
							Descriptor: &api.MessageDescriptor{
								Method:   THREADS_CREATE,
								ObjectId: INVALID,
							},
						},
					},
				}
				response, err := process(app, request)
				Expect(err).To(BeNil())
				Expect(response).To(Not(BeNil()))
				Expect(response.Replies).To(Not(BeNil()))
				Expect(response.Replies).To(HaveLen(1))
				Expect(response.Replies[0].Status).To(Not(BeNil()))
				Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusBadRequest))
			})

			It("receives an error if a Message Descriptor has missing schema", func() {
				request := &api.Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*api.Message{
						{
							Descriptor: &api.MessageDescriptor{
								Method:   THREADS_CREATE,
								ObjectId: OBJECT_ID,
							},
						},
					},
				}
				response, err := process(app, request)
				Expect(err).To(BeNil())
				Expect(response).To(Not(BeNil()))
				Expect(response.Replies).To(Not(BeNil()))
				Expect(response.Replies).To(HaveLen(1))
				Expect(response.Replies[0].Status).To(Not(BeNil()))
				Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusBadRequest))
			})

			It("receives an error if a Message Descriptor has invalid schema", func() {
				request := &api.Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*api.Message{
						{
							Descriptor: &api.MessageDescriptor{
								Method:   THREADS_CREATE,
								ObjectId: OBJECT_ID,
								Schema:   INVALID,
							},
						},
					},
				}
				response, err := process(app, request)
				Expect(err).To(BeNil())
				Expect(response).To(Not(BeNil()))
				Expect(response.Replies).To(Not(BeNil()))
				Expect(response.Replies).To(HaveLen(1))
				Expect(response.Replies[0].Status).To(Not(BeNil()))
				Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusBadRequest))
			})

			It("receives a response if a Message Descriptor has valid", func() {
				request := &api.Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*api.Message{
						{
							Descriptor: &api.MessageDescriptor{
								Method:   THREADS_CREATE,
								ObjectId: OBJECT_ID,
								Schema:   SCHEMA,
							},
						},
					},
				}
				response, err := process(app, request)
				Expect(err).To(BeNil())
				Expect(response).To(Not(BeNil()))
				Expect(response.Replies).To(Not(BeNil()))
				Expect(response.Replies).To(HaveLen(1))
				Expect(response.Replies[0].Status).To(Not(BeNil()))
				Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusOK))
			})

		})

		Context("Reply Tests", func() {

			It("receives an error if a Message Descriptor has missing objectID", func() {
				request := &api.Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*api.Message{
						{
							Descriptor: &api.MessageDescriptor{
								Method: THREADS_REPLY,
							},
						},
					},
				}
				response, err := process(app, request)
				Expect(err).To(BeNil())
				Expect(response).To(Not(BeNil()))
				Expect(response.Replies).To(Not(BeNil()))
				Expect(response.Replies).To(HaveLen(1))
				Expect(response.Replies[0].Status).To(Not(BeNil()))
				Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusBadRequest))
			})

			It("receives an error if a Message Descriptor has invalid objectID", func() {
				request := &api.Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*api.Message{
						{
							Descriptor: &api.MessageDescriptor{
								Method:   THREADS_REPLY,
								ObjectId: INVALID,
							},
						},
					},
				}
				response, err := process(app, request)
				Expect(err).To(BeNil())
				Expect(response).To(Not(BeNil()))
				Expect(response.Replies).To(Not(BeNil()))
				Expect(response.Replies).To(HaveLen(1))
				Expect(response.Replies[0].Status).To(Not(BeNil()))
				Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusBadRequest))
			})

			It("receives an error if a Message Descriptor has missing schema", func() {
				request := &api.Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*api.Message{
						{
							Descriptor: &api.MessageDescriptor{
								Method:   THREADS_REPLY,
								ObjectId: OBJECT_ID,
							},
						},
					},
				}
				response, err := process(app, request)
				Expect(err).To(BeNil())
				Expect(response).To(Not(BeNil()))
				Expect(response.Replies).To(Not(BeNil()))
				Expect(response.Replies).To(HaveLen(1))
				Expect(response.Replies[0].Status).To(Not(BeNil()))
				Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusBadRequest))
			})

			It("receives an error if a Message Descriptor has invalid schema", func() {
				request := &api.Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*api.Message{
						{
							Descriptor: &api.MessageDescriptor{
								Method:   THREADS_REPLY,
								ObjectId: OBJECT_ID,
								Schema:   INVALID,
							},
						},
					},
				}
				response, err := process(app, request)
				Expect(err).To(BeNil())
				Expect(response).To(Not(BeNil()))
				Expect(response.Replies).To(Not(BeNil()))
				Expect(response.Replies).To(HaveLen(1))
				Expect(response.Replies[0].Status).To(Not(BeNil()))
				Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusBadRequest))
			})

			It("receives an error if a Message Descriptor has missing root", func() {
				request := &api.Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*api.Message{
						{
							Descriptor: &api.MessageDescriptor{
								Method:   THREADS_REPLY,
								ObjectId: OBJECT_ID,
								Schema:   SCHEMA,
							},
						},
					},
				}
				response, err := process(app, request)
				Expect(err).To(BeNil())
				Expect(response).To(Not(BeNil()))
				Expect(response.Replies).To(Not(BeNil()))
				Expect(response.Replies).To(HaveLen(1))
				Expect(response.Replies[0].Status).To(Not(BeNil()))
				Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusBadRequest))
			})

			It("receives an error if a Message Descriptor has invalid root", func() {
				request := &api.Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*api.Message{
						{
							Descriptor: &api.MessageDescriptor{
								Method:   THREADS_REPLY,
								ObjectId: OBJECT_ID,
								Schema:   SCHEMA,
								Root:     INVALID,
							},
						},
					},
				}
				response, err := process(app, request)
				Expect(err).To(BeNil())
				Expect(response).To(Not(BeNil()))
				Expect(response.Replies).To(Not(BeNil()))
				Expect(response.Replies).To(HaveLen(1))
				Expect(response.Replies[0].Status).To(Not(BeNil()))
				Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusBadRequest))
			})

			It("receives an error if a Message Descriptor has missing parent", func() {
				request := &api.Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*api.Message{
						{
							Descriptor: &api.MessageDescriptor{
								Method:   THREADS_REPLY,
								ObjectId: OBJECT_ID,
								Schema:   SCHEMA,
								Root:     ROOT,
							},
						},
					},
				}
				response, err := process(app, request)
				Expect(err).To(BeNil())
				Expect(response).To(Not(BeNil()))
				Expect(response.Replies).To(Not(BeNil()))
				Expect(response.Replies).To(HaveLen(1))
				Expect(response.Replies[0].Status).To(Not(BeNil()))
				Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusBadRequest))
			})

			It("receives an error if a Message Descriptor has invalid parent", func() {
				request := &api.Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*api.Message{
						{
							Descriptor: &api.MessageDescriptor{
								Method:   THREADS_REPLY,
								ObjectId: OBJECT_ID,
								Schema:   SCHEMA,
								Root:     ROOT,
								Parent:   INVALID,
							},
						},
					},
				}
				response, err := process(app, request)
				Expect(err).To(BeNil())
				Expect(response).To(Not(BeNil()))
				Expect(response.Replies).To(Not(BeNil()))
				Expect(response.Replies).To(HaveLen(1))
				Expect(response.Replies[0].Status).To(Not(BeNil()))
				Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusBadRequest))
			})

			It("receives a response if a Message Descriptor has valid", func() {
				request := &api.Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*api.Message{
						{
							Descriptor: &api.MessageDescriptor{
								Method:   THREADS_REPLY,
								ObjectId: OBJECT_ID,
								Schema:   SCHEMA,
								Root:     ROOT,
								Parent:   PARENT,
							},
						},
					},
				}
				response, err := process(app, request)
				Expect(err).To(BeNil())
				Expect(response).To(Not(BeNil()))
				Expect(response.Replies).To(Not(BeNil()))
				Expect(response.Replies).To(HaveLen(1))
				Expect(response.Replies[0].Status).To(Not(BeNil()))
				Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusOK))
			})

		})

		Context("Close Tests", func() {

			It("receives an error if a Message Descriptor has missing root", func() {
				request := &api.Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*api.Message{
						{
							Descriptor: &api.MessageDescriptor{
								Method: THREADS_CLOSE,
							},
						},
					},
				}
				response, err := process(app, request)
				Expect(err).To(BeNil())
				Expect(response).To(Not(BeNil()))
				Expect(response.Replies).To(Not(BeNil()))
				Expect(response.Replies).To(HaveLen(1))
				Expect(response.Replies[0].Status).To(Not(BeNil()))
				Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusBadRequest))
			})

			It("receives an error if a Message Descriptor has invalid root", func() {
				request := &api.Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*api.Message{
						{
							Descriptor: &api.MessageDescriptor{
								Method: THREADS_CLOSE,
								Root:   INVALID,
							},
						},
					},
				}
				response, err := process(app, request)
				Expect(err).To(BeNil())
				Expect(response).To(Not(BeNil()))
				Expect(response.Replies).To(Not(BeNil()))
				Expect(response.Replies).To(HaveLen(1))
				Expect(response.Replies[0].Status).To(Not(BeNil()))
				Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusBadRequest))
			})

			It("receives a response if a Message Descriptor has valid", func() {
				request := &api.Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*api.Message{
						{
							Descriptor: &api.MessageDescriptor{
								Method: THREADS_CLOSE,
								Root:   ROOT,
							},
						},
					},
				}
				response, err := process(app, request)
				Expect(err).To(BeNil())
				Expect(response).To(Not(BeNil()))
				Expect(response.Replies).To(Not(BeNil()))
				Expect(response.Replies).To(HaveLen(1))
				Expect(response.Replies[0].Status).To(Not(BeNil()))
				Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusOK))
			})

		})

		Context("Delete Tests", func() {

			It("receives an error if a Message Descriptor has missing root", func() {
				request := &api.Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*api.Message{
						{
							Descriptor: &api.MessageDescriptor{
								Method: THREADS_DELETE,
							},
						},
					},
				}
				response, err := process(app, request)
				Expect(err).To(BeNil())
				Expect(response).To(Not(BeNil()))
				Expect(response.Replies).To(Not(BeNil()))
				Expect(response.Replies).To(HaveLen(1))
				Expect(response.Replies[0].Status).To(Not(BeNil()))
				Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusBadRequest))
			})

			It("receives an error if a Message Descriptor has invalid root", func() {
				request := &api.Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*api.Message{
						{
							Descriptor: &api.MessageDescriptor{
								Method: THREADS_DELETE,
								Root:   INVALID,
							},
						},
					},
				}
				response, err := process(app, request)
				Expect(err).To(BeNil())
				Expect(response).To(Not(BeNil()))
				Expect(response.Replies).To(Not(BeNil()))
				Expect(response.Replies).To(HaveLen(1))
				Expect(response.Replies[0].Status).To(Not(BeNil()))
				Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusBadRequest))
			})

			It("receives a response if a Message Descriptor has valid", func() {
				request := &api.Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*api.Message{
						{
							Descriptor: &api.MessageDescriptor{
								Method: THREADS_DELETE,
								Root:   ROOT,
							},
						},
					},
				}
				response, err := process(app, request)
				Expect(err).To(BeNil())
				Expect(response).To(Not(BeNil()))
				Expect(response.Replies).To(Not(BeNil()))
				Expect(response.Replies).To(HaveLen(1))
				Expect(response.Replies[0].Status).To(Not(BeNil()))
				Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusOK))
			})

		})

	})

})
