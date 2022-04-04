package handler

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/decred/base58"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/decred/dcrd/dcrec/secp256k1/v4/schnorr"
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

func (c *RequestContext) VerifyRequest(signedString string, publicKey *secp256k1.PublicKey) (bool, *errors.MessageLevelError) {

	//sigBytes, err := hex.DecodeString("970603d8ccd2475b1ff66cfb3ce7e622c5938348304c5a7bc2e6015fb98e3b457d4e912fcca6ca87c04390aa5e6e0e613bbbba7ffd6f15bc59f95bbd92ba50f0")
	sigBytes, err := hex.DecodeString(c.Message.Authorization.Signature)
	if err != nil {
		logrus.Errorf("signature decoding failed: %v", err)
		return false, errors.NewMessageLevelError(400, "signature decoding failed", err)
	}
	fmt.Printf("%v", sigBytes)

	//dotIndex := strings.Index(signedString, ".")
	//signatureString := signedString[dotIndex+1:]
	//payloadString := signedString[:dotIndex]
	verified, err := jws.Verify([]byte(signedString), jwa.ES256K, publicKey)
	fmt.Printf("%v", verified)
	return false, nil

	//payloadBytes, _ := hex.DecodeString(signatureString)
	//signatureBytes := sigBytes[:64]

	// signatureBytes := []byte(signatureString)
	// payloadBytes := []byte(payloadString)
	// fmt.Printf("%v", signatureBytes)

	// signature, err := schnorr.ParseSignature(signatureBytes)
	// if err != nil {
	// 	logrus.Errorf("signature parse failed: %v", err)
	// 	return false, errors.NewMessageLevelError(400, "signature parse failed", err)
	// }

	// verified := signature.Verify(payloadBytes, publicKey)
	// return verified, nil
}

func (c *RequestContext) GetPublicKey() (*secp256k1.PublicKey, *errors.MessageLevelError) {

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
	pubKeyBytes := make([]byte, 33)
	pubKeyBytes = didBytes[2:]
	pubKeyBytes = append([]byte{0x2}, pubKeyBytes...)

	pubKey, err := schnorr.ParsePubKey(pubKeyBytes)
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

	// protectedHeaderMap := map[string]string{
	// 	"alg": c.Message.Authorization.Protected.Alg,
	// 	"kid": c.Message.Authorization.Protected.Kid,
	// }

	stringifiedProtected, err := json.Marshal(&c.Message.Authorization.Protected)
	if err != nil {
		logrus.Errorf("stringifiedProtected parse failed: %v", err)
		return "", errors.NewMessageLevelError(400, "stringifiedProtected parse failed", err)
	}

	base64Protected := base64.URLEncoding.EncodeToString([]byte(string(stringifiedProtected)))
	signedString := fmt.Sprintf("%s.%s.", base64Protected, payload)
	return signedString, nil
}

func (c *RequestContext) getMessageHash() (string, error) {
	messageByte, err := c.serializeMessage()
	if err != nil {
		return "", err
	}
	hash := sha256.Sum256(messageByte)
	return fmt.Sprintf("%x", hash[:]), nil
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
