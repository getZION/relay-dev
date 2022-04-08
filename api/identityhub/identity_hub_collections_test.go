package identityhub

import (
	"errors"

	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/constants"
	"github.com/getzion/relay/api/schema"
	"github.com/getzion/relay/api/storage"
	"github.com/getzion/relay/api/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("IdentityHub Collections", func() {
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
		server := InitIdentityHubServer(schemaManager, st)
		app = server.app
	})

	AfterEach(func() {
		gomockController.Finish()
	})

	Context("Message Level", func() {

		Context("Query", func() {

			var request *api.Request

			BeforeEach(func() {
				request = &api.Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*api.Message{
						{
							Descriptor: &api.MessageDescriptor{
								Method: constants.COLLECTIONS_QUERY,
							},
						},
					},
				}
			})

			Context("Validation Tests", func() {

				It("receives an error if Message Descriptor has missing objectID", func() {
					response, err := process(app, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusBadRequest))
				})

				It("receives an error if Message Descriptor has invalid objectID", func() {
					request.Messages[0].Descriptor.ObjectId = INVALID

					response, err := process(app, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusBadRequest))
				})

				It("receives an error if Message Descriptor has invalid schema", func() {
					request.Messages[0].Descriptor.ObjectId = OBJECT_ID
					request.Messages[0].Descriptor.Schema = INVALID

					response, err := process(app, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusBadRequest))
				})

				It("receives an error if Message Descriptor has invalid dataFormat", func() {
					request.Messages[0].Descriptor.ObjectId = OBJECT_ID
					request.Messages[0].Descriptor.Schema = SCHEMA
					request.Messages[0].Descriptor.DataFormat = INVALID

					response, err := process(app, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusBadRequest))
				})

				It("receives an error if Message Descriptor has invalid dateSort", func() {
					request.Messages[0].Descriptor.ObjectId = OBJECT_ID
					request.Messages[0].Descriptor.Schema = SCHEMA
					request.Messages[0].Descriptor.DataFormat = DATA_FORMAT
					request.Messages[0].Descriptor.DateSort = INVALID

					response, err := process(app, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusBadRequest))
				})

			})

			//todo: add more expectation about response & entries
			Context("Communities Tests", func() {

				BeforeEach(func() {
					request.Messages[0].Descriptor.Schema = constants.SCHEMA_COMMUNITY
					st.EXPECT().GetCommunities().Times(1).Return([]api.Community{{Id: 1, Zid: "Zid"}}, nil)
				})

				It("receives a response if Message Descriptor has valid objectID", func() {
					request.Messages[0].Descriptor.ObjectId = OBJECT_ID

					response, err := process(app, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusOK))
				})

				It("receives a response if Message Descriptor has valid dataFormat", func() {
					request.Messages[0].Descriptor.ObjectId = OBJECT_ID
					request.Messages[0].Descriptor.DataFormat = DATA_FORMAT

					response, err := process(app, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusOK))
				})

				It("receives a response if Message Descriptor has valid dateSort (createdAscending)", func() {
					request.Messages[0].Descriptor.ObjectId = OBJECT_ID
					request.Messages[0].Descriptor.DataFormat = DATA_FORMAT
					request.Messages[0].Descriptor.DateSort = "createdAscending"

					response, err := process(app, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusOK))
				})

				It("receives a response if Message Descriptor has valid dateSort (createdDescending)", func() {
					request.Messages[0].Descriptor.ObjectId = OBJECT_ID
					request.Messages[0].Descriptor.DataFormat = DATA_FORMAT
					request.Messages[0].Descriptor.DateSort = "createdDescending"

					response, err := process(app, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusOK))
				})

				It("receives a response if Message Descriptor has valid dateSort (publishedAscending)", func() {
					request.Messages[0].Descriptor.ObjectId = OBJECT_ID
					request.Messages[0].Descriptor.DataFormat = DATA_FORMAT
					request.Messages[0].Descriptor.DateSort = "publishedAscending"

					response, err := process(app, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusOK))
				})

				It("receives a response if Message Descriptor has valid dateSort (publishedDescending)", func() {

					request.Messages[0].Descriptor.ObjectId = OBJECT_ID
					request.Messages[0].Descriptor.DataFormat = DATA_FORMAT
					request.Messages[0].Descriptor.DateSort = "publishedDescending"

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

		Context("Write", func() {

			var request *api.Request

			BeforeEach(func() {
				request = &api.Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*api.Message{
						{
							Descriptor: &api.MessageDescriptor{
								Method: constants.COLLECTIONS_WRITE,
							},
						},
					},
				}
			})

			Context("Validation Tests", func() {

				It("receives an error if a Message Descriptor has missing objectID", func() {
					response, err := process(app, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusBadRequest))
				})

				It("receives an error if a Message Descriptor has invalid objectID", func() {
					request.Messages[0].Descriptor.ObjectId = INVALID
					response, err := process(app, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusBadRequest))
				})

				It("receives an error if a Message Descriptor has missing dateCreated", func() {
					request.Messages[0].Descriptor.ObjectId = OBJECT_ID
					response, err := process(app, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusBadRequest))
				})

				It("receives an error if a Message Descriptor has invalid dateCreated", func() {
					request.Messages[0].Descriptor.ObjectId = OBJECT_ID
					request.Messages[0].Descriptor.DateCreated = INVALID
					response, err := process(app, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusBadRequest))
				})

				It("receives an error if a Message Descriptor has invalid schema", func() {
					request.Messages[0].Descriptor.ObjectId = OBJECT_ID
					request.Messages[0].Descriptor.DateCreated = DATE_CREATED
					request.Messages[0].Descriptor.Schema = INVALID
					response, err := process(app, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusBadRequest))
				})

				It("receives an error if a Message Descriptor has invalid datePublished", func() {
					request.Messages[0].Descriptor.ObjectId = OBJECT_ID
					request.Messages[0].Descriptor.DateCreated = DATE_CREATED
					request.Messages[0].Descriptor.Schema = constants.SCHEMA_COMMUNITY
					request.Messages[0].Descriptor.DatePublished = INVALID
					response, err := process(app, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusBadRequest))
				})

			})

			Context("Communities Tests", func() {

				BeforeEach(func() {
					request.Messages[0].Data = `{ "Name": "test", "Description": "test", "OwnerUsername": "test_username", "OwnerDid": "test_did", "EscrowAmount": 10, "OwnerAlias": "test", "OwnerPubkey": "test", "PricePerMessage": 10, "PriceToJoin": 10 }`
					request.Messages[0].Descriptor.Schema = constants.SCHEMA_COMMUNITY
					request.Messages[0].Descriptor.ObjectId = OBJECT_ID
					request.Messages[0].Descriptor.DateCreated = DATE_CREATED
				})

				It("receives a response if Message Descriptor has valid", func() {
					st.EXPECT().InsertCommunity(gomock.Any()).Times(1).Return(nil)

					response, err := process(app, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.RequestId).To(Equal(request.RequestId))
					Expect(response.Status).To(Not(BeNil()))
					Expect(response.Status.Code).To(Equal(fiber.StatusOK))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusOK))
				})

				It("receives an error if Community already exist", func() {
					st.EXPECT().InsertCommunity(gomock.Any()).Times(1).Return(errors.New("the specified community already exist: test"))

					response, err := process(app, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.RequestId).To(Equal(request.RequestId))
					Expect(response.Status).To(Not(BeNil()))
					Expect(response.Status.Code).To(Equal(fiber.StatusOK))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusBadRequest))
					Expect(response.Replies[0].Status.Message, "the specified community already exist: test")
				})

			})

			Context("Users Tests", func() {

				BeforeEach(func() {
					request.Messages[0].Data = `{ "Did": "did", "Name": "test_name", "Username": "test_username", "Email": "test@test.org" }`
					request.Messages[0].Descriptor.Schema = constants.SCHEMA_PERSON
					request.Messages[0].Descriptor.ObjectId = OBJECT_ID
					request.Messages[0].Descriptor.DateCreated = DATE_CREATED
				})

				It("receive a response if Message Descriptor has valid", func() {
					st.EXPECT().InsertUser(gomock.Any()).Times(1).Return(nil)
					response, err := process(app, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.RequestId).To(Equal(request.RequestId))
					Expect(response.Status).To(Not(BeNil()))
					Expect(response.Status.Code).To(Equal(fiber.StatusOK))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusOK))
				})

				It("receive an error if User already exist", func() {
					st.EXPECT().InsertUser(gomock.Any()).Times(1).Return(errors.New("the specified username already exist: test_username"))
					response, err := process(app, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.RequestId).To(Equal(request.RequestId))
					Expect(response.Status).To(Not(BeNil()))
					Expect(response.Status.Code).To(Equal(fiber.StatusOK))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusBadRequest))
					Expect(response.Replies[0].Status.Message, "the specified username already exist: test_username")
				})

			})

			Context("Conversation Tests", func() {

				BeforeEach(func() {
					request.Messages[0].Data = `{ "CommunityZid": "test_zid", "UserDid": "test_did", "Text": "test"  }`
					request.Messages[0].Descriptor.Schema = constants.SCHEMA_CONVERSATION
					request.Messages[0].Descriptor.ObjectId = OBJECT_ID
					request.Messages[0].Descriptor.DateCreated = DATE_CREATED
				})

				It("receive a response If Message Descriptor has valid", func() {
					st.EXPECT().InsertConversation(gomock.Any()).Times(1).Return(nil)
					response, err := process(app, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.RequestId).To(Equal(request.RequestId))
					Expect(response.Status).To(Not(BeNil()))
					Expect(response.Status.Code).To(Equal(fiber.StatusOK))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusOK))
				})

				//todo: test without Text and Link
				//todo: test with Text, without Link
				//todo: test without Text, without Link

			})

			Context("Join Community Tests", func() {

				BeforeEach(func() {
					request.Messages[0].Data = `{ "communityZid": "zid", "userDid": "did" }`
					request.Messages[0].Descriptor.Schema = constants.SCHEMA_JOIN_COMMUNITY
					request.Messages[0].Descriptor.ObjectId = OBJECT_ID
					request.Messages[0].Descriptor.DateCreated = DATE_CREATED
				})

				It("receive a response if Message Descriptor has valid", func() {
					st.EXPECT().GetCommunityByZid("zid").Times(1).Return(&api.Community{Zid: "zid"}, nil)
					st.EXPECT().GetUserByDid("did").Times(1).Return(&api.User{Did: "did"}, nil)
					st.EXPECT().AddUserToCommunity("zid", "did").Times(1).Return(nil)

					response, err := process(app, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.RequestId).To(Equal(request.RequestId))
					Expect(response.Status).To(Not(BeNil()))
					Expect(response.Status.Code).To(Equal(fiber.StatusOK))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusOK))
				})

			})

			Context("Leave Community Tests", func() {

				BeforeEach(func() {
					request.Messages[0].Data = `{ "communityZid": "zid", "userDid": "did" }`
					request.Messages[0].Descriptor.Schema = constants.SCHEMA_LEAVE_COMMUNITY
					request.Messages[0].Descriptor.ObjectId = OBJECT_ID
					request.Messages[0].Descriptor.DateCreated = DATE_CREATED
				})

				It("receive a response if Message Descriptor has valid", func() {
					st.EXPECT().GetCommunityByZid("zid").Times(1).Return(&api.Community{Zid: "zid"}, nil)
					st.EXPECT().GetUserByDid("did").Times(1).Return(&api.User{Did: "did"}, nil)
					st.EXPECT().RemoveUserToCommunity("zid", "did", "").Times(1).Return(nil)

					response, err := process(app, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.RequestId).To(Equal(request.RequestId))
					Expect(response.Status).To(Not(BeNil()))
					Expect(response.Status.Code).To(Equal(fiber.StatusOK))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusOK))
				})

			})

			Context("Kick User Community Tests", func() {

				BeforeEach(func() {
					request.Messages[0].Data = `{ "communityZid": "zid", "userDid": "did" }`
					request.Messages[0].Descriptor.Schema = constants.SCHEMA_KICK_USER
					request.Messages[0].Descriptor.ObjectId = OBJECT_ID
					request.Messages[0].Descriptor.DateCreated = DATE_CREATED
				})

				It("receive a response if Message Descriptor has valid", func() {
					st.EXPECT().GetCommunityByZid("zid").Times(1).Return(&api.Community{Zid: "zid"}, nil)
					st.EXPECT().GetUserByDid("did").Times(1).Return(&api.User{Did: "did"}, nil)
					st.EXPECT().RemoveUserToCommunity("zid", "did", "Kicked by Owner").Times(1).Return(nil)

					response, err := process(app, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.RequestId).To(Equal(request.RequestId))
					Expect(response.Status).To(Not(BeNil()))
					Expect(response.Status.Code).To(Equal(fiber.StatusOK))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusOK))
				})

			})

		})

		Context("Commit", func() {

			var request *api.Request

			BeforeEach(func() {
				request = &api.Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*api.Message{
						{
							Descriptor: &api.MessageDescriptor{
								Method: constants.COLLECTIONS_COMMIT,
							},
						},
					},
				}
			})

			Context("Validation Tests", func() {

				It("receives an error if Message Descriptor has missing objectID", func() {
					response, err := process(app, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusBadRequest))
				})

				It("receives an error if Message Descriptor has invalid objectID", func() {
					request.Messages[0].Descriptor.ObjectId = INVALID

					response, err := process(app, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusBadRequest))
				})

				It("receives an error if Message Descriptor has missing dateCreated", func() {
					request.Messages[0].Descriptor.ObjectId = OBJECT_ID

					response, err := process(app, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusBadRequest))
				})

				It("receives an error if Message Descriptor has invalid dateCreated", func() {
					request.Messages[0].Descriptor.ObjectId = OBJECT_ID
					request.Messages[0].Descriptor.DateCreated = INVALID

					response, err := process(app, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusBadRequest))
				})

				It("receives an error if Message Descriptor has invalid datePublished", func() {
					request.Messages[0].Descriptor.ObjectId = OBJECT_ID
					request.Messages[0].Descriptor.DateCreated = DATE_CREATED
					request.Messages[0].Descriptor.DatePublished = INVALID

					response, err := process(app, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusBadRequest))
				})

			})

			Context("Communities Tests", func() {

				BeforeEach(func() {
					request.Messages[0].Data = `{ "Name": "test", "Description": "test", "OwnerUsername": "test_username", "OwnerDid": "test_did", "EscrowAmount": 10, "OwnerAlias": "test", "OwnerPubkey": "test", "PricePerMessage": 10, "PriceToJoin": 10 }`
					request.Messages[0].Descriptor.Schema = constants.SCHEMA_COMMUNITY
				})

				It("receives a response if Message Descriptor has valid dateCreated", func() {
					request.Messages[0].Descriptor.ObjectId = OBJECT_ID
					request.Messages[0].Descriptor.DateCreated = DATE_CREATED

					response, err := process(app, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusOK))
				})

				It("receives a response if Message Descriptor has valid datePublished", func() {
					request.Messages[0].Descriptor.ObjectId = OBJECT_ID
					request.Messages[0].Descriptor.DateCreated = DATE_CREATED
					request.Messages[0].Descriptor.DatePublished = DATE_PUBLISHED

					response, err := process(app, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.RequestId).To(Equal(request.RequestId))
					Expect(response.Status).To(Not(BeNil()))
					Expect(response.Status.Code).To(Equal(fiber.StatusOK))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusOK))
				})

			})

		})

		Context("Delete", func() {

			var request *api.Request

			BeforeEach(func() {
				request = &api.Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*api.Message{
						{
							Descriptor: &api.MessageDescriptor{
								Method: constants.COLLECTIONS_DELETE,
							},
						},
					},
				}
			})

			Context("Validation Tests", func() {

				It("receives an error if Message Descriptor has missing objectID", func() {
					response, err := process(app, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusBadRequest))
				})

				It("receives an error if Message Descriptor has invalid objectID", func() {
					request.Messages[0].Descriptor.ObjectId = INVALID
					response, err := process(app, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(fiber.StatusBadRequest))
				})

			})

			Context("Communities Tests", func() {

				BeforeEach(func() {
					request.Messages[0].Descriptor.Schema = constants.SCHEMA_COMMUNITY
				})

				It("receives a response if Message Descriptor has valid objectID", func() {
					request.Messages[0].Descriptor.ObjectId = OBJECT_ID

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
})
