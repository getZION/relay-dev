package user

import (
	"fmt"

	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/validator"
	v1 "github.com/getzion/relay/gen/proto/zion/v1"
	"github.com/go-sql-driver/mysql"
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

func (s *Service) GetById(id int64) (*v1.UserORM, error) {
	var user v1.UserORM
	result := s.connection.DB.Model(&user).First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (s *Service) GetAll() ([]v1.UserORM, error) {
	var users []v1.UserORM
	result := s.connection.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (s *Service) Insert(user v1.UserORM) (*v1.UserORM, error) {

	err := validator.ValidateStruct(&user)
	if err != nil {
		return nil, err
	}

	result := s.connection.DB.Create(&user)
	if result.Error != nil {
		mySqlError := result.Error.(*mysql.MySQLError)
		if mySqlError != nil && mySqlError.Number == 1062 {
			return nil, fmt.Errorf("the specified username already exist: %s", user.Username)
		}
		return nil, result.Error
	}

	return &user, nil
}
