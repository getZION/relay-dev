package community

import (
	"fmt"
	"time"

	"github.com/getzion/relay/api"
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

func (s *Service) GetByZid(zid string) (*v1.CommunityORM, error) {
	var community v1.CommunityORM
	result := s.connection.DB.Model(&community).First(&community, "zid = ?", zid)
	if result.Error != nil {
		return nil, result.Error
	}

	return &community, nil
}

func (s *Service) GetAll() ([]api.Community, error) {

	var communities []v1.CommunityORM
	result := s.connection.DB.
		Preload("Tags").
		Preload("Users").
		Find(&communities)
	if result.Error != nil {
		return nil, result.Error
	}

	var response []api.Community
	for _, community := range communities {

		comm := api.Community{
			Id:              community.Id,
			Zid:             community.Zid,
			Name:            community.Name,
			Description:     community.Description,
			OwnerDid:        community.OwnerDid,
			OwnerUsername:   community.OwnerUsername,
			PricePerMessage: community.PricePerMessage,
			PriceToJoin:     community.PriceToJoin,
			EscrowAmount:    community.EscrowAmount,
			Img:             community.Img,
			Created:         community.Created,
			Updated:         community.Updated,
			LastActive:      community.LastActive,
			Public:          community.Public,
			Deleted:         community.Deleted,
		}

		for _, tag := range community.Tags {
			comm.Tags = append(comm.Tags, tag.Tag)
		}

		for _, user := range community.Users {
			comm.Users = append(comm.Users, user.Did)
		}

		response = append(response, comm)
	}

	return response, nil
}

func (s *Service) Insert(model api.Community) (*v1.CommunityORM, error) {

	currentTime := time.Now().Unix()
	community := v1.CommunityORM{
		Zid:             uuid.NewString(),
		Name:            model.Name,
		Description:     model.Description,
		OwnerDid:        model.OwnerDid,
		OwnerUsername:   model.OwnerUsername,
		PricePerMessage: model.PricePerMessage,
		PriceToJoin:     model.PriceToJoin,
		EscrowAmount:    model.EscrowAmount,
		Img:             model.Img,
		Tags:            make([]*v1.TagORM, 0),
		Created:         currentTime,
		Updated:         currentTime,
		LastActive:      currentTime,
		Public:          model.Public,
	}

	for _, tagString := range model.Tags {

		tag := v1.TagORM{Tag: tagString}
		result := s.connection.DB.Model(&v1.TagORM{}).First(&tag, "tag = ?", tagString)
		if result.Error != nil && result.Error.Error() != "record not found" {
			return nil, result.Error
		}

		community.Tags = append(community.Tags, &tag)
	}

	result := s.connection.DB.Create(&community)
	if result.Error != nil {
		mySqlError, ok := result.Error.(*mysql.MySQLError)
		if !ok {
			return nil, result.Error
		}

		if mySqlError != nil && mySqlError.Number == 1062 {
			return nil, fmt.Errorf("the specified community already exist: %s", community.Name)
		}
		return nil, result.Error
	}

	return &community, nil
}

func (s *Service) AddUserToCommunity(community *v1.CommunityORM, user *v1.UserORM) error {
	association := s.connection.DB.Model(community).Omit("users").Association("users").Append(user)
	return association.Error
}

func (s *Service) RemoveUserToCommunity(community *v1.CommunityORM, user *v1.UserORM) error {
	association := s.connection.DB.Model(community).Omit("users").Association("users").Delete(user)
	return association.Error
}
