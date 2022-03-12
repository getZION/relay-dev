package datastore

import (
	"fmt"

	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/dataservices"
	"github.com/getzion/relay/api/dataservices/community"
	"github.com/getzion/relay/api/dataservices/conversation"
	"github.com/getzion/relay/api/dataservices/payment"
	"github.com/getzion/relay/api/dataservices/user"
)

const (
	organizationSchema = "https://schema.org/Organization"
	conversationSchema = "https://schema.org/Conversation"
	messageSchema      = "https://schema.org/SocialMediaPosting"
	paymentSchema      = "https://schema.org/Invoice"
	personSchema       = "https://schema.org/Person"
)

// Store defines the implementation of api.DataStore using
type Store struct {
	connection *api.Connection
	services   map[string]dataservices.Service

	CommunityService    *community.Service
	ConversationService *conversation.Service
	PaymentService      *payment.Service
	UserService         *user.Service
}

// NewStore initializes a new Store and the associated services
func NewStore(connection *api.Connection) (*Store, error) {
	store := &Store{
		connection: connection,
		services:   make(map[string]dataservices.Service),
	}

	communityService, err := community.NewService(connection)
	if err != nil {
		return nil, err
	}
	store.CommunityService = communityService
	store.services[organizationSchema] = communityService

	conversationService, err := conversation.NewService(connection)
	if err != nil {
		return nil, err
	}
	store.ConversationService = conversationService
	store.services[conversationSchema] = conversationService

	paymentService, err := payment.NewService(connection)
	if err != nil {
		return nil, err
	}
	store.PaymentService = paymentService
	store.services[paymentSchema] = paymentService

	userService, err := user.NewService(connection)
	if err != nil {
		return nil, err
	}
	store.UserService = userService
	store.services[personSchema] = userService

	return store, nil
}

func (s *Store) GetServiceBySchema(schema string) (dataservices.Service, error) {
	if service, ok := s.services[schema]; ok {
		return service, nil
	}

	return nil, fmt.Errorf("unknown schema: %s", schema)
}
