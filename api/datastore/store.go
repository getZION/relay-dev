package datastore

import (
	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/dataservices/community"
	"github.com/getzion/relay/api/dataservices/conversation"
	"github.com/getzion/relay/api/dataservices/payment"
	"github.com/getzion/relay/api/dataservices/user"
)

// Store defines the implementation of api.DataStore using
type Store struct {
	connection *api.Connection

	CommunityService    *community.Service
	ConversationService *conversation.Service
	PaymentService      *payment.Service
	UserService         *user.Service
}

// NewStore initializes a new Store and the associated services
func NewStore(connection *api.Connection) (*Store, error) {
	store := &Store{
		connection: connection,
	}

	communityService, err := community.NewService(connection)
	if err != nil {
		return nil, err
	}
	store.CommunityService = communityService

	conversationService, err := conversation.NewService(connection)
	if err != nil {
		return nil, err
	}
	store.ConversationService = conversationService

	paymentService, err := payment.NewService(connection)
	if err != nil {
		return nil, err
	}
	store.PaymentService = paymentService

	userService, err := user.NewService(connection)
	if err != nil {
		return nil, err
	}
	store.UserService = userService

	//_, err = communityService.Insert([]byte(`{ "id": 1, "Name": "test", "Description": "test", "OwnerDid": "did", "OwnerUsername": "username", "PricePerMessage": 10, "PriceToJoin": 5 }`))
	//userService.Insert([]byte(`{ "Id": 1, "Name": "test", "Username": "test1234", "Email": "test@test.org" }`))

	// todo: make new for these schemas => `https://schema.org/JoinAction`, `https://schema.org/LeaveAction`
	// todo: data will contain UserId and CommunityId ... schemas == dataservices is not cover our requirements. Because we need to get Data from Users and also Communities. maybe we can write some SchemaHandlers ???
	// comm, err := communityService.GetAll()
	// communities := comm.([]v1.CommunityORM)
	// fmt.Printf("%d", len(communities))

	// todo: https://gorm.io/docs/associations.html#Association-Mode
	// ass := connection.DB.Model(&communities[0]).Association("users").Delete(&v1.UserORM{Id: 1})
	// fmt.Printf("%v", ass.Error)

	return store, nil
}
