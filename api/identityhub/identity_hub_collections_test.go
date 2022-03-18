package identityhub

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/getzion/relay/api/constants"
	"github.com/getzion/relay/api/datastore"
	"github.com/getzion/relay/api/schema"
	. "github.com/getzion/relay/gen/proto/identityhub/v1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"golang.org/x/net/context"
)

var _ = Describe("IdentityHub Collections", func() {
	var client *IdentityHubService
	var ctx context.Context
	var mock sqlmock.Sqlmock
	var store *datastore.Store

	BeforeEach(func() {
		store, mock = datastore.NewTestStore()
		schemaManager := schema.NewSchemaManager(store)
		client = InitIdentityHubService(schemaManager)
	})

	Context("Message Level", func() {

		Context("Query", func() {

			var request *Request

			BeforeEach(func() {
				request = &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Descriptor_: &MessageDescriptor{
								Method: constants.COLLECTIONS_QUERY,
							},
						},
					},
				}
			})

			Context("Validation Tests", func() {

				It("receives an error if a Message Descriptor has missing objectID", func() {
					response, err := client.Process(ctx, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(int64(400)))
				})

				It("receives an error if a Message Descriptor has invalid objectID", func() {
					request.Messages[0].Descriptor_.ObjectId = INVALID

					response, err := client.Process(ctx, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(int64(400)))
				})

				It("receives an error if a Message Descriptor has invalid schema", func() {
					request.Messages[0].Descriptor_.ObjectId = OBJECT_ID
					request.Messages[0].Descriptor_.Schema = INVALID

					response, err := client.Process(ctx, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(int64(400)))
				})

				It("receives an error if a Message Descriptor has invalid dataFormat", func() {
					request.Messages[0].Descriptor_.ObjectId = OBJECT_ID
					request.Messages[0].Descriptor_.Schema = SCHEMA
					request.Messages[0].Descriptor_.DataFormat = INVALID

					response, err := client.Process(ctx, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(int64(400)))
				})

				It("receives an error if a Message Descriptor has invalid dateSort", func() {
					request.Messages[0].Descriptor_.ObjectId = OBJECT_ID
					request.Messages[0].Descriptor_.Schema = SCHEMA
					request.Messages[0].Descriptor_.DataFormat = DATA_FORMAT
					request.Messages[0].Descriptor_.DateSort = INVALID

					response, err := client.Process(ctx, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(int64(400)))
				})

			})

			//todo: add more expectation about response & entries
			Context("Communities Tests", func() {

				BeforeEach(func() {
					request.Messages[0].Descriptor_.Schema = constants.SCHEMA_COMMUNITY
				})

				It("receives a response if a Message Descriptor has valid objectID", func() {

					mock.ExpectQuery("SELECT[a-zA-Z *]*").
						WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description", "escrowAmount", "owner_alias", "owner_pubkey", "price_per_message", "price_to_join"}).
							AddRow(1, "test", "desc", 0, "alias", "pubkey", 10, 10))

					request.Messages[0].Descriptor_.ObjectId = OBJECT_ID

					response, err := client.Process(ctx, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(int64(200)))
				})

				It("receives a response if a Message Descriptor has valid dataFormat", func() {

					mock.ExpectQuery("SELECT[a-zA-Z *]*").
						WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description", "escrowAmount", "owner_alias", "owner_pubkey", "price_per_message", "price_to_join"}).
							AddRow(1, "test", "desc", 0, "alias", "pubkey", 10, 10).
							AddRow(1, "test2", "desc2", 0, "alias2", "pubkey2", 20, 20))

					request.Messages[0].Descriptor_.ObjectId = OBJECT_ID
					request.Messages[0].Descriptor_.DataFormat = DATA_FORMAT

					response, err := client.Process(ctx, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(int64(200)))
				})

				It("receives a response if a Message Descriptor has valid dateSort (createdAscending)", func() {

					mock.ExpectQuery("SELECT[a-zA-Z *]*").
						WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description", "escrowAmount", "owner_alias", "owner_pubkey", "price_per_message", "price_to_join"}).
							AddRow(1, "test", "desc", 0, "alias", "pubkey", 10, 10))

					request.Messages[0].Descriptor_.ObjectId = OBJECT_ID
					request.Messages[0].Descriptor_.DataFormat = DATA_FORMAT
					request.Messages[0].Descriptor_.DateSort = "createdAscending"

					response, err := client.Process(ctx, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(int64(200)))
				})

				It("receives a response if a Message Descriptor has valid dateSort (createdDescending)", func() {

					mock.ExpectQuery("SELECT[a-zA-Z *]*").
						WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description", "escrowAmount", "owner_alias", "owner_pubkey", "price_per_message", "price_to_join"}).
							AddRow(1, "test", "desc", 0, "alias", "pubkey", 10, 10))

					request.Messages[0].Descriptor_.ObjectId = OBJECT_ID
					request.Messages[0].Descriptor_.DataFormat = DATA_FORMAT
					request.Messages[0].Descriptor_.DateSort = "createdDescending"

					response, err := client.Process(ctx, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(int64(200)))
				})

				It("receives a response if a Message Descriptor has valid dateSort (publishedAscending)", func() {

					mock.ExpectQuery("SELECT[a-zA-Z *]*").
						WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description", "escrowAmount", "owner_alias", "owner_pubkey", "price_per_message", "price_to_join"}).
							AddRow(1, "test", "desc", 0, "alias", "pubkey", 10, 10))

					request.Messages[0].Descriptor_.ObjectId = OBJECT_ID
					request.Messages[0].Descriptor_.DataFormat = DATA_FORMAT
					request.Messages[0].Descriptor_.DateSort = "publishedAscending"

					response, err := client.Process(ctx, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(int64(200)))
				})

				It("receives a response if a Message Descriptor has valid dateSort (publishedDescending)", func() {

					mock.ExpectQuery("SELECT[a-zA-Z *]*").
						WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description", "escrowAmount", "owner_alias", "owner_pubkey", "price_per_message", "price_to_join"}).
							AddRow(1, "test", "desc", 0, "alias", "pubkey", 10, 10))

					request.Messages[0].Descriptor_.ObjectId = OBJECT_ID
					request.Messages[0].Descriptor_.DataFormat = DATA_FORMAT
					request.Messages[0].Descriptor_.DateSort = "publishedDescending"

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

		Context("Write", func() {

			var request *Request

			BeforeEach(func() {
				request = &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Descriptor_: &MessageDescriptor{
								Method: constants.COLLECTIONS_WRITE,
							},
						},
					},
				}
			})

			Context("Validation Tests", func() {

				It("receives an error if a Message Descriptor has missing objectID", func() {
					response, err := client.Process(ctx, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(int64(400)))
				})

				It("receives an error if a Message Descriptor has invalid objectID", func() {
					request.Messages[0].Descriptor_.ObjectId = INVALID
					response, err := client.Process(ctx, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(int64(400)))
				})

				It("receives an error if a Message Descriptor has missing dateCreated", func() {
					request.Messages[0].Descriptor_.ObjectId = OBJECT_ID
					response, err := client.Process(ctx, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(int64(400)))
				})

				It("receives an error if a Message Descriptor has invalid dateCreated", func() {
					request.Messages[0].Descriptor_.ObjectId = OBJECT_ID
					request.Messages[0].Descriptor_.DateCreated = INVALID
					response, err := client.Process(ctx, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(int64(400)))
				})

				It("receives an error if a Message Descriptor has invalid schema", func() {
					request.Messages[0].Descriptor_.ObjectId = OBJECT_ID
					request.Messages[0].Descriptor_.DateCreated = DATE_CREATED
					request.Messages[0].Descriptor_.Schema = INVALID
					response, err := client.Process(ctx, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(int64(400)))
				})

				It("receives an error if a Message Descriptor has invalid datePublished", func() {
					request.Messages[0].Descriptor_.ObjectId = OBJECT_ID
					request.Messages[0].Descriptor_.DateCreated = DATE_CREATED
					request.Messages[0].Descriptor_.Schema = constants.SCHEMA_COMMUNITY
					request.Messages[0].Descriptor_.DatePublished = INVALID
					response, err := client.Process(ctx, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(int64(400)))
				})

			})

			Context("Communities Tests", func() {

				BeforeEach(func() {
					request.Messages[0].Data = `{ "Name": "test", "Description": "test", "OwnerUsername": "test_username", "OwnerDid": "test_did", "EscrowAmount": 10, "OwnerAlias": "test", "OwnerPubkey": "test", "PricePerMessage": 10, "PriceToJoin": 10 }`
					request.Messages[0].Descriptor_.Schema = constants.SCHEMA_COMMUNITY
				})

				It("receives a response if a Message Descriptor has valid dateCreated", func() {

					mock.ExpectBegin()
					mock.ExpectExec("INSERT INTO `communities`[a-zA-Z *]*").WillReturnResult(sqlmock.NewResult(1, 1))
					mock.ExpectCommit()

					request.Messages[0].Descriptor_.ObjectId = OBJECT_ID
					request.Messages[0].Descriptor_.DateCreated = DATE_CREATED

					response, err := client.Process(ctx, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(int64(200)))
					Expect(mock.ExpectationsWereMet()).To(BeNil())
				})

				It("receives a response if a Message Descriptor has valid datePublished", func() {

					mock.ExpectBegin()
					mock.ExpectExec("INSERT INTO `communities`[a-zA-Z *]*").WillReturnResult(sqlmock.NewResult(1, 1))
					mock.ExpectCommit()

					request.Messages[0].Descriptor_.ObjectId = OBJECT_ID
					request.Messages[0].Descriptor_.DateCreated = DATE_CREATED
					request.Messages[0].Descriptor_.DatePublished = DATE_PUBLISHED

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
					Expect(mock.ExpectationsWereMet()).To(BeNil())
				})

			})

		})

		Context("Commit", func() {

			var request *Request

			BeforeEach(func() {
				request = &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Descriptor_: &MessageDescriptor{
								Method: constants.COLLECTIONS_COMMIT,
							},
						},
					},
				}
			})

			Context("Validation Tests", func() {

				It("receives an error if a Message Descriptor has missing objectID", func() {
					response, err := client.Process(ctx, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(int64(400)))
				})

				It("receives an error if a Message Descriptor has invalid objectID", func() {
					request.Messages[0].Descriptor_.ObjectId = INVALID

					response, err := client.Process(ctx, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(int64(400)))
				})

				It("receives an error if a Message Descriptor has missing dateCreated", func() {
					request.Messages[0].Descriptor_.ObjectId = OBJECT_ID

					response, err := client.Process(ctx, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(int64(400)))
				})

				It("receives an error if a Message Descriptor has invalid dateCreated", func() {
					request.Messages[0].Descriptor_.ObjectId = OBJECT_ID
					request.Messages[0].Descriptor_.DateCreated = INVALID

					response, err := client.Process(ctx, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(int64(400)))
				})

				It("receives an error if a Message Descriptor has invalid datePublished", func() {
					request.Messages[0].Descriptor_.ObjectId = OBJECT_ID
					request.Messages[0].Descriptor_.DateCreated = DATE_CREATED
					request.Messages[0].Descriptor_.DatePublished = INVALID

					response, err := client.Process(ctx, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(int64(400)))
				})

			})

			Context("Communities Tests", func() {

				BeforeEach(func() {
					request.Messages[0].Data = `{ "Name": "test", "Description": "test", "OwnerUsername": "test_username", "OwnerDid": "test_did", "EscrowAmount": 10, "OwnerAlias": "test", "OwnerPubkey": "test", "PricePerMessage": 10, "PriceToJoin": 10 }`
					request.Messages[0].Descriptor_.Schema = constants.SCHEMA_COMMUNITY
				})

				It("receives a response if a Message Descriptor has valid dateCreated", func() {
					request.Messages[0].Descriptor_.ObjectId = OBJECT_ID
					request.Messages[0].Descriptor_.DateCreated = DATE_CREATED

					response, err := client.Process(ctx, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(int64(200)))
				})

				It("receives a response if a Message Descriptor has valid datePublished", func() {
					request.Messages[0].Descriptor_.ObjectId = OBJECT_ID
					request.Messages[0].Descriptor_.DateCreated = DATE_CREATED
					request.Messages[0].Descriptor_.DatePublished = DATE_PUBLISHED

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

		Context("Delete", func() {

			var request *Request

			BeforeEach(func() {
				request = &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Descriptor_: &MessageDescriptor{
								Method: constants.COLLECTIONS_DELETE,
							},
						},
					},
				}
			})

			Context("Validation Tests", func() {

				It("receives an error if a Message Descriptor has missing objectID", func() {
					response, err := client.Process(ctx, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(int64(400)))
				})

				It("receives an error if a Message Descriptor has invalid objectID", func() {
					request.Messages[0].Descriptor_.ObjectId = INVALID
					response, err := client.Process(ctx, request)
					Expect(err).To(BeNil())
					Expect(response).To(Not(BeNil()))
					Expect(response.Replies).To(Not(BeNil()))
					Expect(response.Replies).To(HaveLen(1))
					Expect(response.Replies[0].Status).To(Not(BeNil()))
					Expect(response.Replies[0].Status.Code).To(Equal(int64(400)))
				})

			})

			Context("Communities Tests", func() {

				BeforeEach(func() {
					request.Messages[0].Descriptor_.Schema = constants.SCHEMA_COMMUNITY
				})

				It("receives a response if a Message Descriptor has valid objectID", func() {
					request.Messages[0].Descriptor_.ObjectId = OBJECT_ID

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
})
