package community

import (
	"fmt"
	"time"

	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/validator"
	v1 "github.com/getzion/relay/gen/proto/zion/v1"
	"github.com/go-sql-driver/mysql"
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

func (s *Service) GetById(id int64) (*v1.CommunityORM, error) {
	var community v1.CommunityORM
	result := s.connection.DB.Model(&community).First(&community, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &community, nil
}

func (s *Service) GetAll() ([]v1.CommunityORM, error) {
	var communities []v1.CommunityORM
	result := s.connection.DB.Find(&communities)
	if result.Error != nil {
		return nil, result.Error
	}
	return communities, nil
}

func (s *Service) Insert(community v1.CommunityORM) (*v1.CommunityORM, error) {

	community.Zid = uuid.NewString()
	community.Created = time.Now().Unix()
	community.Updated = community.Created
	community.LastActive = community.Created

	err := validator.ValidateStruct(&community)
	if err != nil {
		return nil, err
	}

	result := s.connection.DB.Create(&community)
	if result.Error != nil {
		mySqlError := result.Error.(*mysql.MySQLError)
		if mySqlError != nil && mySqlError.Number == 1062 {
			return nil, fmt.Errorf("the specified community already exist: %s", community.Name)
		}
		return nil, result.Error
	}

	return &community, nil
}

func (s *Service) AddUserToCommunity(community *v1.CommunityORM, userId int64) error {
	association := s.connection.DB.Model(community).Association("users").Append(&v1.UserORM{Id: userId})
	return association.Error
}

func (s *Service) RemoveUserToCommunity(community *v1.CommunityORM, userId int64) error {
	association := s.connection.DB.Model(community).Association("users").Delete(&v1.UserORM{Id: userId})
	return association.Error
}
