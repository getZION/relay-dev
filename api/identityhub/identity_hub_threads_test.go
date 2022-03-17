package identityhub

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/datastore"
	"github.com/getzion/relay/api/schema"
	. "github.com/getzion/relay/gen/proto/identityhub/v1"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

const (
	THREADS_QUERY  = "ThreadsQuery"
	THREADS_CREATE = "ThreadsCreate"
	THREADS_REPLY  = "ThreadsReply"
	THREADS_CLOSE  = "ThreadsClose"
	THREADS_DELETE = "ThreadsDelete"
)

var _ = Describe("IdentityHub Threads", func() {
	var client *IdentityHubService
	var ctx context.Context
	//var mock sqlmock.Sqlmock

	BeforeEach(func() {
		var err error
		var db *sql.DB

		db, _, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
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

		schemaManager := schema.NewSchemaManager(store)

		client = &IdentityHubService{
			prefix:                   prefix,
			validHubInterfaceMethods: validHubInterfaceMethods,
			schemaManager:            schemaManager,
		}
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
