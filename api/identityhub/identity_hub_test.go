package identityhub

import (
	"fmt"

	"github.com/getzion/relay/api/schema"
	"github.com/getzion/relay/api/storage"
	. "github.com/getzion/relay/gen/proto/identityhub/v1"
	"github.com/golang/mock/gomock"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"golang.org/x/net/context"
)

const (
	REQUEST_ID     = "3eb8ea70-7ea5-4069-a153-cfb0ea682df9"
	TARGET         = "TheTarget"
	OBJECT_ID      = "d82c0026-ed42-4b26-81f3-94805958a75c"
	DATE_CREATED   = "1645917431"
	DATE_PUBLISHED = "1645917431"
	SCHEMA         = "https://test.com"
	DATA_FORMAT    = "application/json"
	ROOT           = "e23ea8cf-5e64-42d0-b3c6-54e5ab1dcf25"
	PARENT         = "654a4593-4c01-4a6c-9cd9-6bf04bd3d441"
	INVALID        = "<invalid>"
)

type GinkgoTestReporter struct{}

func (g GinkgoTestReporter) Errorf(format string, args ...interface{}) {
	Fail(fmt.Sprintf(format, args...))
}

func (g GinkgoTestReporter) Fatalf(format string, args ...interface{}) {
	Fail(fmt.Sprintf(format, args...))
}

var _ = Describe("IdentityHub", func() {
	var (
		t                GinkgoTestReporter
		gomockController *gomock.Controller
		client           *IdentityHubService
		ctx              context.Context
		st               *storage.MockStorage
	)

	BeforeEach(func() {
		gomockController = gomock.NewController(t)
		st = storage.NewMockStorage(gomockController)
		schemaManager := schema.NewSchemaManager(st)

		client = &IdentityHubService{
			prefix:                   prefix,
			validHubInterfaceMethods: validHubInterfaceMethods,
			schemaManager:            schemaManager,
		}
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
				Target:   TARGET,
				Messages: make([]*Message, 1),
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
				RequestId: INVALID,
				Messages:  make([]*Message, 1),
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
				Messages:  make([]*Message, 1),
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
							Method: INVALID,
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

	})
})
