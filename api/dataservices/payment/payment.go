package payment

import (
	"github.com/getzion/relay/api"
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
	var payments []v1.PaymentORM
	result := s.connection.DB.Find(&payments)
	if result.Error != nil {
		return nil, result.Error
	}
	return payments, nil
}
