package conversation

import (
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

func (s *Service) GetAll() ([]v1.ConversationORM, error) {
	var conversations []v1.ConversationORM
	result := s.connection.DB.Find(&conversations)
	if result.Error != nil {
		return nil, result.Error
	}
	return conversations, nil
}

func (s *Service) Insert(conversation v1.ConversationORM) (*v1.ConversationORM, error) {

	conversation.Zid = uuid.NewString()

	err := validator.ValidateStruct(&conversation)
	if err != nil {
		return nil, err
	}

	result := s.connection.DB.Create(&conversation)
	if result.Error != nil {
		return nil, result.Error
	}

	return &conversation, nil
}
