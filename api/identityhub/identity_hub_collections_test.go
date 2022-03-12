package identityhub

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/datastore"
	. "github.com/getzion/relay/gen/proto/identityhub/v1"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

const (
	COLLECTIONS_QUERY  = "CollectionsQuery"
	COLLECTIONS_WRITE  = "CollectionsWrite"
	COLLECTIONS_COMMIT = "CollectionsCommit"
	COLLECTIONS_DELETE = "CollectionsDelete"
)

var _ = Describe("IdentityHub Collections", func() {
	var client *IdentityHubService
	var ctx context.Context
	var mock sqlmock.Sqlmock

	BeforeEach(func() {
		var err error
		var db *sql.DB

		db, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
		if err != nil {
			logrus.Panic(err)
		}

		gormDb, err := gorm.Open("mysql", db)
		if err != nil {
			logrus.Panic(err)
		}

		connection := &api.Connection{
			DB: gormDb,
		}

		store, err := datastore.NewStore(connection)
		if err != nil {
			logrus.Panic(err)
		}

		client = InitIdentityHubService(store)
	})

	AfterEach(func() {
		//defer conn.Close()
	})

	Context("Message Level", func() {

		Context("Query", func() {

			Context("Validation Tests", func() {

				It("receives an error if a Message Descriptor has missing objectID", func() {
					request := &Request{
						RequestId: REQUEST_ID,
						Target:    TARGET,
						Messages: []*Message{
							{
								Descriptor_: &MessageDescriptor{
									Method: COLLECTIONS_QUERY,
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
									Method:   COLLECTIONS_QUERY,
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

				It("receives an error if a Message Descriptor has invalid schema", func() {
					request := &Request{
						RequestId: REQUEST_ID,
						Target:    TARGET,
						Messages: []*Message{
							{
								Descriptor_: &MessageDescriptor{
									Method:   COLLECTIONS_QUERY,
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

				It("receives an error if a Message Descriptor has invalid dataFormat", func() {
					request := &Request{
						RequestId: REQUEST_ID,
						Target:    TARGET,
						Messages: []*Message{
							{
								Descriptor_: &MessageDescriptor{
									Method:     COLLECTIONS_QUERY,
									ObjectId:   OBJECT_ID,
									Schema:     SCHEMA,
									DataFormat: INVALID,
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

				It("receives an error if a Message Descriptor has invalid dateSort", func() {
					request := &Request{
						RequestId: REQUEST_ID,
						Target:    TARGET,
						Messages: []*Message{
							{
								Descriptor_: &MessageDescriptor{
									Method:     COLLECTIONS_QUERY,
									ObjectId:   OBJECT_ID,
									Schema:     SCHEMA,
									DataFormat: DATA_FORMAT,
									DateSort:   INVALID,
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

			})

			//todo: add more expectation about response & entries
			Context("Communities Tests", func() {

				It("receives a response if a Message Descriptor has valid objectID", func() {

					mock.ExpectQuery("SELECT[a-zA-Z *]*").
						WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description", "escrowAmount", "owner_alias", "owner_pubkey", "price_per_message", "price_to_join"}).
							AddRow(1, "test", "desc", 0, "alias", "pubkey", 10, 10))

					request := &Request{
						RequestId: REQUEST_ID,
						Target:    TARGET,
						Messages: []*Message{
							{
								Descriptor_: &MessageDescriptor{
									Method:   COLLECTIONS_QUERY,
									Schema:   SCHEMA_ORGANIZATION,
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

				It("receives a response if a Message Descriptor has valid schema", func() {

					mock.ExpectQuery("SELECT[a-zA-Z *]*").
						WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description", "escrowAmount", "owner_alias", "owner_pubkey", "price_per_message", "price_to_join"}).
							AddRow(1, "test", "desc", 0, "alias", "pubkey", 10, 10))

					request := &Request{
						RequestId: REQUEST_ID,
						Target:    TARGET,
						Messages: []*Message{
							{
								Descriptor_: &MessageDescriptor{
									Method:   COLLECTIONS_QUERY,
									Schema:   SCHEMA_ORGANIZATION,
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

				It("receives a response if a Message Descriptor has valid dataFormat", func() {

					mock.ExpectQuery("SELECT[a-zA-Z *]*").
						WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description", "escrowAmount", "owner_alias", "owner_pubkey", "price_per_message", "price_to_join"}).
							AddRow(1, "test", "desc", 0, "alias", "pubkey", 10, 10).
							AddRow(1, "test2", "desc2", 0, "alias2", "pubkey2", 20, 20))

					request := &Request{
						RequestId: REQUEST_ID,
						Target:    TARGET,
						Messages: []*Message{
							{
								Descriptor_: &MessageDescriptor{
									Method:     COLLECTIONS_QUERY,
									ObjectId:   OBJECT_ID,
									Schema:     SCHEMA_ORGANIZATION,
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

				It("receives a response if a Message Descriptor has valid dateSort (createdAscending)", func() {

					mock.ExpectQuery("SELECT[a-zA-Z *]*").
						WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description", "escrowAmount", "owner_alias", "owner_pubkey", "price_per_message", "price_to_join"}).
							AddRow(1, "test", "desc", 0, "alias", "pubkey", 10, 10))

					request := &Request{
						RequestId: REQUEST_ID,
						Target:    TARGET,
						Messages: []*Message{
							{
								Descriptor_: &MessageDescriptor{
									Method:     COLLECTIONS_QUERY,
									ObjectId:   OBJECT_ID,
									Schema:     SCHEMA_ORGANIZATION,
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

				It("receives a response if a Message Descriptor has valid dateSort (createdDescending)", func() {

					mock.ExpectQuery("SELECT[a-zA-Z *]*").
						WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description", "escrowAmount", "owner_alias", "owner_pubkey", "price_per_message", "price_to_join"}).
							AddRow(1, "test", "desc", 0, "alias", "pubkey", 10, 10))

					request := &Request{
						RequestId: REQUEST_ID,
						Target:    TARGET,
						Messages: []*Message{
							{
								Descriptor_: &MessageDescriptor{
									Method:     COLLECTIONS_QUERY,
									ObjectId:   OBJECT_ID,
									Schema:     SCHEMA_ORGANIZATION,
									DataFormat: DATA_FORMAT,
									DateSort:   "createdDescending",
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

				It("receives a response if a Message Descriptor has valid dateSort (publishedAscending)", func() {

					mock.ExpectQuery("SELECT[a-zA-Z *]*").
						WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description", "escrowAmount", "owner_alias", "owner_pubkey", "price_per_message", "price_to_join"}).
							AddRow(1, "test", "desc", 0, "alias", "pubkey", 10, 10))

					request := &Request{
						RequestId: REQUEST_ID,
						Target:    TARGET,
						Messages: []*Message{
							{
								Descriptor_: &MessageDescriptor{
									Method:     COLLECTIONS_QUERY,
									ObjectId:   OBJECT_ID,
									Schema:     SCHEMA_ORGANIZATION,
									DataFormat: DATA_FORMAT,
									DateSort:   "publishedAscending",
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

				It("receives a response if a Message Descriptor has valid dateSort (publishedDescending)", func() {

					mock.ExpectQuery("SELECT[a-zA-Z *]*").
						WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description", "escrowAmount", "owner_alias", "owner_pubkey", "price_per_message", "price_to_join"}).
							AddRow(1, "test", "desc", 0, "alias", "pubkey", 10, 10))

					request := &Request{
						RequestId: REQUEST_ID,
						Target:    TARGET,
						Messages: []*Message{
							{
								Descriptor_: &MessageDescriptor{
									Method:     COLLECTIONS_QUERY,
									ObjectId:   OBJECT_ID,
									Schema:     SCHEMA_ORGANIZATION,
									DataFormat: DATA_FORMAT,
									DateSort:   "publishedDescending",
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

		Context("Write Tests", func() {

			It("receives an error if a Message Descriptor has missing objectID", func() {
				request := &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Descriptor_: &MessageDescriptor{
								Method: COLLECTIONS_WRITE,
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
								Method:   COLLECTIONS_WRITE,
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

			It("receives an error if a Message Descriptor has missing dateCreated", func() {
				request := &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Descriptor_: &MessageDescriptor{
								Method:   COLLECTIONS_WRITE,
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
								Method:      COLLECTIONS_WRITE,
								ObjectId:    OBJECT_ID,
								DateCreated: INVALID,
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

				mock.ExpectBegin()
				mock.ExpectExec("INSERT[a-zA-Z *]*").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()

				request := &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Data: `{ "Name": "test", "Description": "test", "OwnerUsername": "test_username", "OwnerDid": "test_did", "EscrowAmount": 10, "OwnerAlias": "test", "OwnerPubkey": "test", "PricePerMessage": 10, "PriceToJoin": 10 }`,
							Descriptor_: &MessageDescriptor{
								Method:      COLLECTIONS_WRITE,
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
				Expect(mock.ExpectationsWereMet()).To(BeNil())
			})

			It("receives an error if a Message Descriptor has invalid schema", func() {
				request := &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Descriptor_: &MessageDescriptor{
								Method:      COLLECTIONS_WRITE,
								ObjectId:    OBJECT_ID,
								DateCreated: DATE_CREATED,
								Schema:      INVALID,
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

				mock.ExpectBegin()
				mock.ExpectExec("INSERT[a-zA-Z *]*").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()

				request := &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Data: `{ "Name": "test", "Description": "test", "OwnerUsername": "test_username", "OwnerDid": "test_did", "EscrowAmount": 10, "OwnerAlias": "test", "OwnerPubkey": "test", "PricePerMessage": 10, "PriceToJoin": 10 }`,
							Descriptor_: &MessageDescriptor{
								Method:      COLLECTIONS_WRITE,
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
				Expect(mock.ExpectationsWereMet()).To(BeNil())
			})

			It("receives an error if a Message Descriptor has invalid datePublished", func() {
				request := &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Descriptor_: &MessageDescriptor{
								Method:        COLLECTIONS_WRITE,
								ObjectId:      OBJECT_ID,
								DateCreated:   DATE_CREATED,
								Schema:        SCHEMA,
								DatePublished: INVALID,
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

				mock.ExpectBegin()
				mock.ExpectExec("INSERT[a-zA-Z *]*").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()

				request := &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Data: `{ "Name": "test", "Description": "test", "OwnerUsername": "test_username", "OwnerDid": "test_did", "EscrowAmount": 10, "OwnerAlias": "test", "OwnerPubkey": "test", "PricePerMessage": 10, "PriceToJoin": 10 }`,
							Descriptor_: &MessageDescriptor{
								Method:        COLLECTIONS_WRITE,
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
				Expect(mock.ExpectationsWereMet()).To(BeNil())
			})

			It("receives a response if a Message Descriptor has valid data", func() {

				mock.ExpectBegin()
				mock.ExpectExec("INSERT[a-zA-Z *]*").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()

				request := &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Data: `{ "Name": "test", "Description": "test", "OwnerUsername": "test_username", "OwnerDid": "test_did", "EscrowAmount": 10, "OwnerAlias": "test", "OwnerPubkey": "test", "PricePerMessage": 10, "PriceToJoin": 10 }`,
							Descriptor_: &MessageDescriptor{
								Method:        COLLECTIONS_WRITE,
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
				Expect(mock.ExpectationsWereMet()).To(BeNil())
			})

		})

		Context("Commit Tests", func() {

			It("receives an error if a Message Descriptor has missing objectID", func() {
				request := &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Descriptor_: &MessageDescriptor{
								Method: COLLECTIONS_COMMIT,
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
								Method:   COLLECTIONS_COMMIT,
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

			It("receives an error if a Message Descriptor has missing dateCreated", func() {
				request := &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Descriptor_: &MessageDescriptor{
								Method:   COLLECTIONS_COMMIT,
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
								Method:      COLLECTIONS_COMMIT,
								ObjectId:    OBJECT_ID,
								DateCreated: INVALID,
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
								Method:      COLLECTIONS_COMMIT,
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
								Method:      COLLECTIONS_COMMIT,
								ObjectId:    OBJECT_ID,
								DateCreated: DATE_CREATED,
								Schema:      INVALID,
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
								Method:      COLLECTIONS_COMMIT,
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
								Method:        COLLECTIONS_COMMIT,
								ObjectId:      OBJECT_ID,
								DateCreated:   DATE_CREATED,
								Schema:        SCHEMA,
								DatePublished: INVALID,
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
								Method:        COLLECTIONS_COMMIT,
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

		Context("Delete Tests", func() {

			It("receives an error if a Message Descriptor has missing objectID", func() {
				request := &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Descriptor_: &MessageDescriptor{
								Method: COLLECTIONS_DELETE,
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
								Method:   COLLECTIONS_DELETE,
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

			It("receives a response if a Message Descriptor has valid objectID", func() {
				request := &Request{
					RequestId: REQUEST_ID,
					Target:    TARGET,
					Messages: []*Message{
						{
							Descriptor_: &MessageDescriptor{
								Method:   COLLECTIONS_DELETE,
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
