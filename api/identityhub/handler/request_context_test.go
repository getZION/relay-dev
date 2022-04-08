package handler

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/getzion/relay/api"
	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
)

func Test_RequetContext_GetPublicKey(t *testing.T) {

	context := RequestContext{
		Message: &api.Message{
			Data: "eyJjb21tdW5pdHkiOnsibmFtZSI6IlRlc3QgQ29tbXVuaXR5IiwiZGVzY3JpcHRpb24iOiJBd2Vzb21lIHRlc3QgY29tbXVuaXR5IiwicHJpY2VQZXJNZXNzYWdlIjo1LCJwcmljZVRvSm9pbiI6NX19",
			Descriptor: &api.MessageDescriptor{
				Cid:         "bagaaierapb3hfsapznktlc37grbcxlrlevueikala7pspw7dum4bzvpnegda",
				DataFormat:  "application/json",
				DateCreated: "1649421537831",
				Method:      "CollectionsWrite",
				ObjectId:    "d3c82418-b797-4d4a-b3c1-d521022d794c",
			},
			Attestation: &api.Attestation{
				Payload: "bagaaierablumuihbhsskqlvtqck5sh45dogpa3qvtvdd24qmxrnitta4df4a",
				Protected: &api.AttestationProtected{
					Alg: "ES256K",
					Kid: "did:key:z7r8orgCJdAbLdE37wXTa5gtwoDb28AmFKkPJXVgo9YgyzhT1wZs6nLwS6YGHAYucu8dAn1cophhkTqH7aS3sQtnohRCk",
				},
				Signature: "4a110804186cf5a3fbfcb90a372b660abfe33f76e9402531df6872bf895bc3ef2d24b8ca994a54d1b117fad4f716336dec648b036d61e784eaacd29d9008ed01",
			},
		},
	}

	messageByte, _ := json.Marshal(context.Message)
	prefix := cid.Prefix{
		Version:  1,
		Codec:    cid.Raw,
		MhType:   multihash.SHA2_256,
		MhLength: -1,
	}

	messageCid, _ := prefix.Sum(messageByte)
	fmt.Printf("%v\n", messageCid.String())

	cid, _ := cid.Decode(context.Message.Descriptor.Cid)
	fmt.Printf("%v\n", cid.String())

	if messageCid.Equals(cid) {
		fmt.Printf("equal")
	}

	publicKey, _ := context.GetPublicKey()
	signedString, _ := context.SignPayload()
	context.VerifyRequest(signedString, publicKey)
}
