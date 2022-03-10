package community

import (
	"github.com/getzion/relay/api"
	v1 "github.com/getzion/relay/gen/proto/zion/v1"
	"github.com/sirupsen/logrus"
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

func (s *Service) GetAll() []v1.CommunityORM {
	var communities []v1.CommunityORM
	result := s.connection.DB.Find(&communities)
	//result := s.connection.DB.Table("communities").Select("*").Scan(&out)
	logrus.Infof("%v", result.Value)
	return communities
}

func (s *Service) Insert(value interface{}) error {
	result := s.connection.DB.Create(value)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
