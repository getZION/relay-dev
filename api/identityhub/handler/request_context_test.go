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
			Data: "{\"name\":\"Test Community\",\"description\":\"Test Description\",\"ownerDid\":\"did:key:z6DtUbeLv8CZ9qjieQYoRHczeqq6UfvTwMUDYqu5HaQKeAzS\",\"ownerUsername\":\"Testo Mang\",\"img\":\"https://placekitten.com/g/200/300\",\"priceToJoin\":1,\"pricePerMessage\":2,\"escrowAmount\":3,\"lastActive\":1648847624849,\"tags\":[],\"public\":true,\"zid\":\"d97d001e-abe2-483d-9591-c88e7212516b\"}",
			Descriptor_: &v1.MessageDescriptor{
				DateCreated: "1648847624849",
				ObjectId:    "4e26b6c3-25be-4fdd-bb9a-269eaf971d00",
				Schema:      "https://schema.org/Organization",
			},
			Authorization: &v1.Authorization{
				Protected: &v1.AttestationProtected{
					Alg: "SS256K",
					Kid: "did:key:z6DtUbeLv8CZ9qjieQYoRHczeqq6UfvTwMUDYqu5HaQKeAzS",
				},
				Payload:   "d522c9df4cfd90e8698f7ab286255ff5b8b27251932d22fa318b3003ab07e97b",
				Signature: "3733613332376163353138343032646138323864323461346634643738306166383138373434643162643630396331643530346137303532396538663335663032326161373338366130393531373836613265646561346638623334363637343764623436396635666264616163383965336538386231376437656666613264",
			},
		},
	}

	context.GetPublicKey()

}
