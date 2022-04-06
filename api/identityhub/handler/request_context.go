package handler

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/decred/base58"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/getzion/relay/api/identityhub/errors"
	"github.com/getzion/relay/api/schema"
	hub "github.com/getzion/relay/gen/proto/identityhub/v1"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jws"
	"github.com/sirupsen/logrus"
)

type RequestContext struct {
	SchemaManager *schema.SchemaManager
	Message       *hub.Message
}

func (c *RequestContext) VerifyRequest(signedString string, publicKey *ecdsa.PublicKey) (bool, *errors.MessageLevelError) {

	//dotIndex := strings.Index(signedString, ".")
	//signatureString := signedString[:dotIndex]
	//payloadString := signedString[dotIndex+1:]
	//payloadString = strings.TrimSuffix(payloadString, ".")

	priv, _ := crypto.ToECDSA([]byte{126, 73, 27, 160, 104, 154, 68, 160, 31, 132, 149, 10, 124, 242, 99, 186, 1, 168, 169, 105, 253, 198, 35, 89, 109, 23, 251, 117, 82, 21, 23, 125})
	pKey := priv.PublicKey
	verified, err := jws.Verify([]byte(signedString), jwa.ES256K, pKey)

	verified, err = jws.Verify([]byte(signedString), jwa.ES256K, publicKey)
	fmt.Printf("verified: %v, error:%v", verified, err)
	return false, nil
}

func (c *RequestContext) GetPublicKey() (*ecdsa.PublicKey, *errors.MessageLevelError) {

	if c.Message.Authorization == nil {
		return nil, errors.NewMessageLevelError(400, "authorization cannot be null or empty", nil)
	} else if c.Message.Authorization.Protected == nil {
		return nil, errors.NewMessageLevelError(400, "authorization protected cannot be null", nil)
	} else if c.Message.Authorization.Protected.Alg != "SS256K" {
		return nil, errors.NewMessageLevelError(400, "Unsupported signing algorithm", nil)
	} else if strings.Trim(c.Message.Authorization.Protected.Kid, " ") == "" {
		return nil, errors.NewMessageLevelError(400, "Unsupported DID method", nil)
	} else if strings.HasPrefix(c.Message.Authorization.Protected.Kid, "did:key:") == false {
		return nil, errors.NewMessageLevelError(400, "Unsupported DID method", nil)
	}

	kid := strings.TrimPrefix(c.Message.Authorization.Protected.Kid, "did:key:z")
	didBytes := base58.Decode(kid)
	pubKeyBytes := make([]byte, 65)
	pubKeyBytes = didBytes[2:]

	pubKey, err := crypto.UnmarshalPubkey(pubKeyBytes)
	if err != nil {
		return nil, errors.NewMessageLevelError(400, "publicKey parse failed", err)
	}
	return pubKey, nil
}

func (c *RequestContext) SignPayload() (string, *errors.MessageLevelError) {
	payload, err := c.getMessageHash()
	if err != nil {
		logrus.Errorf("payload parse failed: %v", err)
		return "", errors.NewMessageLevelError(400, "payload parse failed", err)
	}

	stringifiedProtected, err := json.Marshal(&c.Message.Authorization.Protected)
	if err != nil {
		logrus.Errorf("stringifiedProtected parse failed: %v", err)
		return "", errors.NewMessageLevelError(400, "stringifiedProtected parse failed", err)
	}

	base64Protected := base64.RawURLEncoding.EncodeToString([]byte(string(stringifiedProtected)))
	signedString := fmt.Sprintf("%s.%s.", base64Protected, payload)
	return signedString, nil
}

func (c *RequestContext) getMessageHash() (string, error) {
	messageByte, err := c.serializeMessage()
	if err != nil {
		return "", err
	}
	hash := sha256.Sum256(messageByte)
	return fmt.Sprintf("%x", hash), nil
}

func (c *RequestContext) serializeMessage() ([]byte, error) {
	value := []string{
		c.Message.Authorization.Protected.Kid,
		c.Message.Descriptor_.DateCreated,
		c.Message.Descriptor_.ObjectId,
		c.Message.Descriptor_.Schema,
		c.Message.Data,
	}

	jsonString, err := json.Marshal(&value)
	if err != nil {
		return nil, err
	}

	return jsonString, nil
}
