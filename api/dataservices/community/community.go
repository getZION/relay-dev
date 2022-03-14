package community

import (
	"encoding/json"

	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/validator"
	v1 "github.com/getzion/relay/gen/proto/zion/v1"
)

// Service represents a service for managing environment(endpoint) data.
type Service struct {
	connection *api.Connection
}

// NewService creates a new instance of a service.
func NewService(connection *api.Connection) (*Service, error) {
	return &Service{
		connection: connection,
	}, nil
}

func (s *Service) GetAll() (interface{}, error) {
	var communities []v1.CommunityORM
	result := s.connection.DB.Find(&communities)
	if result.Error != nil {
		return nil, result.Error
	}
	return communities, nil
}

func (s *Service) Insert(data []byte) error {

	var community v1.CommunityORM
	json.Unmarshal(data, &community)
	err := validator.ValidateStruct(&community)
	if err != nil {
		return err
	}

	result := s.connection.DB.Create(community)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
