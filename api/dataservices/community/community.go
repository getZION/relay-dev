package community

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

func (s *Service) Find(out v1.CommunityORM, where ...interface{}) {
	s.connection.DB.Find(&out, where...)
}

func (s *Service) Insert(value interface{}) {
	s.connection.DB.Create(value)
}
