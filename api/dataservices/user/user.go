package user

import (
	"fmt"
	"time"

	"github.com/getzion/relay/api"
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

func (s *Service) GetByDid(did string) (*v1.UserORM, error) {
	var user v1.UserORM
	result := s.connection.DB.Model(&user).First(&user, "did = ?", did)
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

func (s *Service) Insert(model api.User) (*v1.UserORM, error) {

	currentTime := time.Now().Unix()
	user := v1.UserORM{
		Name:           model.Name,
		Did:            model.Did,
		Email:          model.Email,
		Username:       model.Username,
		Created:        currentTime,
		Updated:        currentTime,
		Img:            model.Img,
		Bio:            model.Bio,
		PriceToMessage: model.PriceToMessage,
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
