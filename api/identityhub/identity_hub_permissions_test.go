package identityhub

import (
	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/schema"
	"github.com/getzion/relay/api/storage"
	"github.com/getzion/relay/api/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	PERMISSIONS_REQUEST = "PermissionsRequest"
	PERMISSIONS_QUERY   = "PermissionsQuery"
	PERMISSIONS_GRANT   = "PermissionsGrant"
	PERMISSIONS_REVOKE  = "PermissionsRevoke"
)

var _ = Describe("IdentityHub Permissions", func() {
	var (
		t                GinkgoTestReporter
		gomockController *gomock.Controller
		app              *fiber.App
		st               *storage.MockStorage
	)

	BeforeEach(func() {
		validator.InitValidator()
		gomockController = gomock.NewController(t)
		st = storage.NewMockStorage(gomockController)
		schemaManager := schema.NewSchemaManager(st)

		server := InitIdentityHubServer(schemaManager)
		app = server.app
	})

	AfterEach(func() {
		gomockController.Finish()
	})

	Context("Message Level", func() {

		Context("Request Tests", func() {

			It("receives an error if a Message Descriptor has missing schema", func() {
				request := &api.Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*api.Message{
						{
							Descriptor: &api.MessageDescriptor{
								Method: PERMISSIONS_REQUEST,
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
								Method: PERMISSIONS_REQUEST,
								Schema: INVALID,
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
								Method: PERMISSIONS_REQUEST,
								Schema: SCHEMA,
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

		Context("Query Tests", func() {

			It("receives an error if a Message Descriptor has missing schema", func() {
				request := &api.Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*api.Message{
						{
							Descriptor: &api.MessageDescriptor{
								Method: PERMISSIONS_QUERY,
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
								Method: PERMISSIONS_QUERY,
								Schema: INVALID,
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
								Method: PERMISSIONS_QUERY,
								Schema: SCHEMA,
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

		Context("Grant Tests", func() {

			It("receives an error if a Message Descriptor has missing schema", func() {
				request := &api.Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*api.Message{
						{
							Descriptor: &api.MessageDescriptor{
								Method: PERMISSIONS_GRANT,
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
								Method: PERMISSIONS_GRANT,
								Schema: INVALID,
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
								Method: PERMISSIONS_GRANT,
								Schema: SCHEMA,
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

		Context("Revoke Tests", func() {

			It("receives an error if a Message Descriptor has missing ObjectId", func() {
				request := &api.Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*api.Message{
						{
							Descriptor: &api.MessageDescriptor{
								Method: PERMISSIONS_REVOKE,
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

			It("receives an error if a Message Descriptor has invalid ObjectId", func() {
				request := &api.Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*api.Message{
						{
							Descriptor: &api.MessageDescriptor{
								Method:   PERMISSIONS_REVOKE,
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

			It("receives a response if a Message Descriptor has valid", func() {
				request := &api.Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*api.Message{
						{
							Descriptor: &api.MessageDescriptor{
								Method:   PERMISSIONS_REVOKE,
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
				Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusOK))
			})

		})

	})

})
