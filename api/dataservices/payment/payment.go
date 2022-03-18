package payment

import (
	"encoding/json"

	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/validator"
	v1 "github.com/getzion/relay/gen/proto/zion/v1"
	"github.com/google/uuid"
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

func (s *Service) Insert(data []byte) (interface{}, error) {

	var payment v1.PaymentORM
	err := json.Unmarshal(data, &payment)
	if err != nil {
		return nil, err
	}

	payment.Zid = uuid.NewString()

	err = validator.ValidateStruct(&payment)
	if err != nil {
		return nil, err
	}

	result := s.connection.DB.Create(&payment)
	if result.Error != nil {
		return nil, result.Error
	}

	return &payment, nil
}