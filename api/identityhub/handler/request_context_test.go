package handler

import (
	"testing"

	"github.com/getzion/relay/api"
)

func Test_RequetContext_GetPublicKey(t *testing.T) {

	// tests := []struct {
	// 	name                 string
	// 	context              *RequestContext
	// 	errorExpected        bool
	// 	expectedStatusCode   int64
	// 	expectedErrorMessage string
	// }{
	// 	{
	// 		name: "",
	// 	},
	// 	// should return authorization cannot be null
	// 	// authorization protected cannot be null
	// 	// Unsupported signing algorithm
	// 	// Unsupported DID method
	// 	// Unsupported DID method
	// }

	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		publicKey, err := tt.context.GetPublicKey()

	// 		if tt.errorExpected {
	// 			require.Nil(t, publicKey)
	// 			require.NotNil(t, err)
	// 		}
	// 	})
	// }

	context := RequestContext{
		Message: &api.Message{
			Data: "{\"name\":\"Test Community\",\"description\":\"Awesome test community\",\"pricePerMessage\":5,\"priceToJoin\":5,\"zid\":\"abd122ad-4a9a-417d-a96f-7096fbb331b8\"}",
			Descriptor: &api.MessageDescriptor{
				DateCreated: "1649267340735",
				ObjectId:    "2f4d42de-96da-453e-ba57-c35ef9a323b2",
				Schema:      "https://schema.org/Organization",
			},
			Authorization: &api.Authorization{
				Protected: &api.AuthorizationProtected{
					Alg: "SS256K",
					Kid: "did:key:z7r8orgCJdAbLdE37wXTa5gtwoDb28AmFKkPJXVgo9YgyzhT1wZs6nLwS6YGHAYucu8dAn1cophhkTqH7aS3sQtnohRCk",
				},
				Payload:   "585d655f2f7fccb97db688105ed7355dec93a25990369fd15565779667bd9f6b",
				Signature: "a38ab25f3c025f92507ab2ac137fb97f92e62dd0594c53330196dd05bc8f4cfa6e010bb962b99c1ce00032efcaf9b16e0e83e80b30a2d907d8f921c1c2519d3f",
			},
		},
	}

	publicKey, _ := context.GetPublicKey()
	signedString, _ := context.SignPayload()
	context.VerifyRequest(signedString, publicKey)
}
