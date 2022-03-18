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

	return store, nil
}
