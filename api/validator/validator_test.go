package validator

import (
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/getzion/relay/api"
	v1 "github.com/getzion/relay/gen/proto/zion/v1"
	"github.com/stretchr/testify/require"
)

func Test_ShouldValidate_Community(t *testing.T) {
	InitValidator()

	faker.SetGenerateUniqueValues(false)

	tests := []struct {
		name               string
		generate           func() interface{}
		expectedError      bool
		expectedErrorCount int
	}{
		{
			name: "Name field should be required",
			generate: func() interface{} {
				model := api.Community{}
				faker.FakeData(&model)
				model.Name = ""
				return model
			},
			expectedError:      true,
			expectedErrorCount: 1,
		},
		{
			name: "Description field should be maximum 250 character",
			generate: func() interface{} {
				model := api.Community{}
				faker.FakeData(&model)
				faker.SetRandomStringLength(300)
				faker.FakeData(&model.Description)
				faker.SetRandomStringLength(25)
				return model
			},
			expectedError:      true,
			expectedErrorCount: 1,
		},
		{
			name: "OwnerDid field should be required",
			generate: func() interface{} {
				model := api.Community{}
				faker.FakeData(&model)
				model.OwnerDid = ""
				return model
			},
			expectedError:      true,
			expectedErrorCount: 1,
		},
		{
			name: "OwnerUsername field should be required",
			generate: func() interface{} {
				model := api.Community{}
				faker.FakeData(&model)
				model.OwnerUsername = ""
				return model
			},
			expectedError:      true,
			expectedErrorCount: 1,
		},
		{
			name: "EscrowAmount field should be greater or equal to zero",
			generate: func() interface{} {
				model := api.Community{}
				faker.FakeData(&model)
				model.EscrowAmount = -1
				return model
			},
			expectedError:      true,
			expectedErrorCount: 1,
		},
		{
			name: "EscrowAmount field should be less than 100000",
			generate: func() interface{} {
				model := api.Community{}
				faker.FakeData(&model)
				model.EscrowAmount = 100000
				return model
			},
			expectedError:      true,
			expectedErrorCount: 1,
		},
		{
			name: "PricePerMessage field should be greater or equal to zero",
			generate: func() interface{} {
				model := api.Community{}
				faker.FakeData(&model)
				model.PricePerMessage = -1
				return model
			},
			expectedError:      true,
			expectedErrorCount: 1,
		},
		{
			name: "PricePerMessage field should be less than 100000",
			generate: func() interface{} {
				model := api.Community{}
				faker.FakeData(&model)
				model.PricePerMessage = 100000
				return model
			},
			expectedError:      true,
			expectedErrorCount: 1,
		},
		{
			name: "PriceToJoin field should be greater or equal to zero",
			generate: func() interface{} {
				model := api.Community{}
				faker.FakeData(&model)
				model.PriceToJoin = -1
				return model
			},
			expectedError:      true,
			expectedErrorCount: 1,
		},
		{
			name: "PriceToJoin field should be less than 100000",
			generate: func() interface{} {
				model := api.Community{}
				faker.FakeData(&model)
				model.PriceToJoin = 100000
				return model
			},
			expectedError:      true,
			expectedErrorCount: 1,
		},
		{
			name: "should be valid",
			generate: func() interface{} {
				model := api.Community{
					Name:            "test",
					Description:     "test",
					EscrowAmount:    10,
					OwnerDid:        "test_did",
					OwnerUsername:   "test_username",
					PricePerMessage: 2,
					PriceToJoin:     5,
				}
				return model
			},
			expectedError:      false,
			expectedErrorCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			model := tt.generate()
			err := ValidateStruct(model)

			if tt.expectedError {
				require.NotNil(t, err)
				require.Len(t, err, tt.expectedErrorCount)
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func Test_ShouldValidate_Conversation(t *testing.T) {
	InitValidator()

	faker.SetGenerateUniqueValues(false)

	tests := []struct {
		name               string
		generate           func() interface{}
		expectedError      bool
		expectedErrorCount int
	}{
		{
			name: "Zid field should be required",
			generate: func() interface{} {
				model := v1.ConversationORM{}
				faker.FakeData(&model)
				model.Zid = ""
				return model
			},
			expectedError:      true,
			expectedErrorCount: 1,
		},
		{
			name: "should be valid",
			generate: func() interface{} {
				model := v1.ConversationORM{}
				faker.FakeData(&model)
				return model
			},
			expectedError:      false,
			expectedErrorCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			model := tt.generate()
			err := ValidateStruct(model)

			if tt.expectedError {
				require.NotNil(t, err)
				require.Len(t, err, tt.expectedErrorCount)
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func Test_ShouldValidate_User(t *testing.T) {
	InitValidator()

	faker.SetGenerateUniqueValues(false)

	tests := []struct {
		name               string
		generate           func() interface{}
		expectedError      bool
		expectedErrorCount int
	}{
		{
			name: "Name field should be required",
			generate: func() interface{} {
				model := v1.UserORM{}
				faker.FakeData(&model)
				model.Email = faker.Email()
				model.Username = faker.Username()
				model.Name = ""
				return model
			},
			expectedError:      true,
			expectedErrorCount: 1,
		},
		{
			name: "Email field should be invalid",
			generate: func() interface{} {
				model := v1.UserORM{}
				faker.FakeData(&model)
				model.Username = faker.Username()
				model.Email = "test-"
				return model
			},
			expectedError:      true,
			expectedErrorCount: 1,
		},
		{
			name: "Username field should be at least 6 character",
			generate: func() interface{} {
				model := v1.UserORM{}
				faker.FakeData(&model)
				model.Username = "usern"
				model.Email = faker.Email()
				return model
			},
			expectedError:      true,
			expectedErrorCount: 1,
		},
		{
			name: "Username field should be maximum 16 character",
			generate: func() interface{} {
				model := v1.UserORM{}
				faker.FakeData(&model)
				model.Username = "maximumsixteencharacter"
				model.Email = faker.Email()
				return model
			},
			expectedError:      true,
			expectedErrorCount: 1,
		},
		{
			name: "Username field should be alphanumeric and underscore",
			generate: func() interface{} {
				model := v1.UserORM{}
				faker.FakeData(&model)
				model.Username = "test*user!"
				model.Email = faker.Email()
				return model
			},
			expectedError:      true,
			expectedErrorCount: 1,
		},
		{
			name: "should be valid",
			generate: func() interface{} {
				model := v1.UserORM{}
				faker.FakeData(&model)
				model.Username = "test_user"
				model.Email = faker.Email()
				return model
			},
			expectedError:      false,
			expectedErrorCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			model := tt.generate()
			err := ValidateStruct(model)

			if tt.expectedError {
				require.NotNil(t, err)
				require.Len(t, err, tt.expectedErrorCount)
			} else {
				require.Nil(t, err)
			}
		})
	}
}
