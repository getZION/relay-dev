package datastore

import (
	"testing"

	"github.com/getzion/relay/api/dataservices"
	"github.com/getzion/relay/api/dataservices/community"
	"github.com/getzion/relay/api/dataservices/conversation"
	"github.com/getzion/relay/api/dataservices/payment"
	"github.com/getzion/relay/api/dataservices/user"
	"github.com/stretchr/testify/require"
)

func Test_ShouldReturnCorrectService(t *testing.T) {
	tests := []struct {
		name            string
		schema          string
		expectedService dataservices.Service
	}{
		{
			name:            "should return community service",
			schema:          organizationSchema,
			expectedService: &community.Service{},
		},
		{
			name:            "should return conversation service",
			schema:          conversationSchema,
			expectedService: &conversation.Service{},
		},
		{
			name:            "should return payment service",
			schema:          paymentSchema,
			expectedService: &payment.Service{},
		},
		{
			name:            "should return user service",
			schema:          personSchema,
			expectedService: &user.Service{},
		},
	}

	store, _ := NewStore(nil)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service, err := store.GetServiceBySchema(tt.schema)

			require.Nil(t, err)
			require.IsType(t, tt.expectedService, service)
		})
	}
}
