package conversation

import (
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

func (s *Service) GetAll() ([]api.Conversation, error) {
	var conversations []v1.ConversationORM
	result := s.connection.DB.
		Preload("Comments").
		Find(&conversations)
	if result.Error != nil {
		return nil, result.Error
	}

	var response []api.Conversation
	for _, conversation := range conversations {
		con := api.Conversation{
			Id:           conversation.Id,
			Zid:          conversation.Zid,
			CommunityZid: conversation.CommunityZid,
			Text:         conversation.Text,
			Link:         conversation.Link,
			Img:          conversation.Img,
			Video:        conversation.Video,
			Public:       conversation.Public,
			PublicPrice:  conversation.PublicPrice,
			Created:      conversation.Created,
			Updated:      conversation.Updated,
			Deleted:      conversation.Deleted,
		}

		for _, comment := range conversation.Comments {
			con.Comments = append(con.Comments, api.Comment{
				Id:      comment.Id,
				Zid:     comment.Zid,
				UserDid: comment.UserDid,
				Text:    comment.Text,
				Created: comment.Created,
				Updated: comment.Updated,
				Deleted: comment.Deleted,
			})
		}

		response = append(response, con)
	}

	return response, nil
}

func (s *Service) Insert(model api.Conversation) (*v1.ConversationORM, error) {

	currentTime := time.Now().Unix()
	conversation := v1.ConversationORM{
		Zid:          uuid.NewString(),
		CommunityZid: model.CommunityZid,
		Text:         model.Text,
		Link:         model.Link,
		Img:          model.Img,
		Video:        model.Video,
		Public:       model.Public,
		PublicPrice:  model.PublicPrice,
		Created:      currentTime,
		Updated:      currentTime,
	}

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
