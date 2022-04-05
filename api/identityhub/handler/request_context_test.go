package handler

import (
	"testing"

	v1 "github.com/getzion/relay/gen/proto/identityhub/v1"
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
		Message: &v1.Message{
			Data: "{\"name\":\"Test Community\",\"description\":\"Awesome test community\",\"price_per_message\":5,\"price_to_join\":5,\"zid\":\"671f7355-201b-4b3c-80ef-32793eb3d884\"}",
			Descriptor_: &v1.MessageDescriptor{
				//DateCreated: "1648847624849",
				//ObjectId:    "4e26b6c3-25be-4fdd-bb9a-269eaf971d00",
				Schema: "https://schema.org/Organization",
			},
			Authorization: &v1.Authorization{
				Protected: &v1.AttestationProtected{
					Alg: "SS256K",
					Kid: "did:key:zQ3shr4enfaB31BsZtNovYYtuWfArzFbUiF5nofvtJbYdU5G5",
					//Kid: "did:key:z6DtNFZhbpJtsrN4x7sKEuMBpf1X3EwDqjfM7RVGYdh3qADy",
				},
				Payload:   "86143185119c926d401263858683074825fd5ae9b5051152bc27d1fa9470b9ff",
				Signature: "c4b14dcf038af59e45e9677b3ff3143ede713af23948be8fa6444974bc459b4371ce9c55aa1a03d0f8eb016f231f3fff528000b1fa8a6149e6cdefd8ab030170",
			},
		},
	}

	publicKey, _ := context.GetPublicKey()
	signedString, _ := context.SignPayload()
	context.VerifyRequest(signedString, publicKey)

}
