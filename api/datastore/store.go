package datastore

import (
	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/dataservices/community"
)

// Store defines the implementation of api.DataStore using
type Store struct {
	connection *api.Connection

	CommunityService *community.Service
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

	return store, nil
}
