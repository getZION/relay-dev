package user

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
	var users []v1.UserORM
	result := s.connection.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (s *Service) Insert(data []byte) (interface{}, error) {

	var user v1.UserORM
	json.Unmarshal(data, &user)
	err := validator.ValidateStruct(&user)
	if err != nil {
		return nil, err
	}

	result := s.connection.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
