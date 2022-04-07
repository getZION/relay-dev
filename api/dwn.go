package api

func (e *ErrorHandler) Error() string {
	return e.Message
}

type (
	ErrorHandler struct {
		StatusCode int
		Message    string
		RequestId  string
		Err        error
	}

	// Request Objects are JSON object envelopes used to pass messages to Identity Hub instances.
	// https://identity.foundation/identity-hub/spec/#request-objects
	Request struct {
		// The Request Object MUST include a requestId property, and its value MUST be
		// an [RFC4122] UUID Version 4 string to identify the request.
		RequestId string `json:"requestId,omitempty"`
		// The Request Object MUST include a target property, and its value MUST be the
		// Decentralized Identifier base URI of the DID-relative URL.
		Target string `json:"target,omitempty"`
		// The Request Object MUST include a messages property, and its value MUST be an
		// array composed of Message objects.
		Messages []*Message `json:"messages,omitempty"`
	}

	// Responses from Interface method invocations are JSON objects that MUST be constructed as follows:
	// https://identity.foundation/identity-hub/spec/#response-objects
	Response struct {
		// The object MUST include an requestId property, and its value MUST be the
		// [RFC4122] UUID Version 4 string from the requestId property of the Request Object it is in response to.
		RequestId string `json:"requestId,omitempty"`
		// The object MUST have a status property, and its value MUST be an object composed of the following properties:
		Status *Status `json:"status,omitempty"`
		// The object MAY have a replies property, and if present its value MUST be an
		// array of Message Result Objects, which are constructed as follows:
		Replies []*Reply `json:"replies,omitempty"`
	}

	// All Identity Hub messaging is transacted via Messages JSON objects.
	// These objects contain message execution parameters, authorization material,
	// authorization signatures, and signing/encryption information.
	// For various purposes Messages rely on IPFS CIDs and DAG APIs.
	// https://identity.foundation/identity-hub/spec/#messages
	Message struct {
		// Message objects MAY contain a data property, and if present its value MUST
		// be a JSON value of the Message’s data.
		Data string `json:"data,omitempty"`
		// Message objects MUST contain a descriptor property, and its value MUST be an
		// object, as defined by the Message Descriptors section of this specification.
		Descriptor *MessageDescriptor `json:"descriptor,omitempty"`
		// Message objects MAY contain an attestation property, and if present its value
		// MUST be an object, as defined by the Signed Data section of this specification.
		Attestation *Attestation `json:"attestation,omitempty"`
		// If a Message object requires signatory and/or permission-evaluated authorization,
		// it MUST include an authorization property, and its value MUST be a [RFC7515]
		// JSON Web Signature composed as follows...
		Authorization *Authorization `json:"authorization,omitempty"`
	}

	MessageDescriptor struct {
		Method        string `json:"method,omitempty"`
		ObjectId      string `json:"objectId,omitempty"`
		Schema        string `json:"schema,omitempty"`
		DataFormat    string `json:"dataFormat,omitempty"`
		DateCreated   string `json:"dateCreated,omitempty"`
		DatePublished string `json:"datePublished,omitempty"`
		DateSort      string `json:"dateSort,omitempty"`
		Root          string `json:"root,omitempty"`
		Parent        string `json:"parent,omitempty"`
		Cid           string `json:"cid,omitempty"`
	}

	Attestation struct {
		Protected *AttestationProtected `json:"protected,omitempty"`
		Payload   string                `json:"payload,omitempty"`
		Signature string                `json:"signature,omitempty"`
	}

	AttestationProtected struct {
		Alg string `json:"alg,omitempty"`
		Kid string `json:"kid,omitempty"`
	}

	// todo: https://identity.foundation/identity-hub/spec/#message-descriptors
	Authorization struct {
		Protected *AuthorizationProtected `json:"protected,omitempty"`
		Payload   string                  `json:"payload,omitempty"`
		Signature string                  `json:"signature,omitempty"`
	}

	AuthorizationProtected struct {
		Alg string `json:"alg,omitempty"`
		Kid string `json:"kid,omitempty"`
	}

	Status struct {
		// The status object MUST have a code property, and its value MUST be an integer
		// set to the HTTP Status Code appropriate for the status of the response.
		Code int `json:"code,omitempty"`
		// The status object MAY have a message property, and if present its value MUST
		// be a string that describes a terse summary of the status. It is RECOMMENDED
		// that the implementer set the message text to the standard title of the HTTP
		// Status Code, when a title/message has already been defined for that code.
		Message string `json:"message,omitempty"`
	}

	Reply struct {
		// The object MUST have a messageId property, and its value MUST be the stringified
		// Version 1 CID of the associated message in the Request Object from which it was
		// received.
		MessageId string `json:"messageId,omitempty"`
		// The object MUST have a status property, and its value MUST be an object composed
		// of the following properties:
		Status *Status `json:"status,omitempty"`
		//
		// The object MAY have a entries property if the message request was successful.
		// If present, its value MUST be the resulting message entries returned from the
		// invocation of the corresponding message.
		Entries []string `json:"entries,omitempty"`
	}
)
