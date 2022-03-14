package community

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

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
	var communities []v1.CommunityORM
	result := s.connection.DB.Find(&communities)
	if result.Error != nil {
		return nil, result.Error
	}
	return communities, nil
}

func (s *Service) Insert(data []byte) (interface{}, error) {

	var community v1.CommunityORM
	json.Unmarshal(data, &community)
	err := validator.ValidateStruct(&community)
	if err != nil {
		return nil, err
	}

	var exists int64
	err = s.connection.DB.Model(v1.CommunityORM{}).Where("name = ?", strings.ToLower(community.Name)).Count(&exists).Error
	if err != nil || exists > 0 {
		return nil, fmt.Errorf("the specified community already exist: %s", community.Name)
	}

	community.Zid = uuid.NewString()
	community.Created = time.Now().Unix()
	community.Updated = community.Created
	community.LastActive = community.Created

	result := s.connection.DB.Create(&community)
	if result.Error != nil {
		return nil, result.Error
	}

	return &community, nil
}
