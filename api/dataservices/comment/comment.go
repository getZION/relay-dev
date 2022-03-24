package comment

import (
	"time"

	"github.com/getzion/relay/api"
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

func (s *Service) GetAll() ([]v1.CommentORM, error) {

	var comments []v1.CommentORM
	result := s.connection.DB.Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	}

	return comments, nil
}

func (s *Service) Insert(model api.Comment) (*v1.CommentORM, error) {

	currentTime := time.Now().Unix()
	comment := v1.CommentORM{
		Zid:     uuid.NewString(),
		UserDid: model.UserDid,
		Text:    model.Text,
		Created: currentTime,
		Updated: currentTime,
	}

	result := s.connection.DB.Create(&comment)
	if result.Error != nil {
		return nil, result.Error
	}

	return &comment, nil
}
