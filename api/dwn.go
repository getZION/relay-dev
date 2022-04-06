package api

type (

	// Request Objects are JSON object envelopes used to pass messages to Identity Hub
	// instances.
	// https://identity.foundation/identity-hub/spec/#request-objects
	Request struct {
		// The Request Object MUST include a requestId property, and its value MUST be
		// an [RFC4122] UUID Version 4 string to identify the request.
		RequestId string `protobuf:"bytes,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
		// The Request Object MUST include a target property, and its value MUST be the
		// Decentralized Identifier base URI of the DID-relative URL.
		Target string `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
		// The Request Object MUST include a messages property, and its value MUST be an
		// array composed of Message objects.
		Messages []*Message `protobuf:"bytes,3,rep,name=messages,proto3" json:"messages,omitempty"`
	}

	// Responses from Interface method invocations are JSON objects that MUST be
	// constructed as follows:
	// https://identity.foundation/identity-hub/spec/#response-objects
	Response struct {
		// The object MUST include an requestId property, and its value MUST be the
		// [RFC4122] UUID Version 4 string from the requestId property of the Request Object
		// it is in response to.
		RequestId string `protobuf:"bytes,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
		// The object MUST have a status property, and its value MUST be an object composed
		// of the following properties:
		Status *Status `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
		// The object MAY have a replies property, and if present its value MUST be an
		// array of Message Result Objects, which are constructed as follows:
		Replies []*Reply `protobuf:"bytes,3,rep,name=replies,proto3" json:"replies,omitempty"`
	}

	// All Identity Hub messaging is transacted via Messages JSON objects.
	// These objects contain message execution parameters, authorization material,
	// authorization signatures, and signing/encryption information.
	// For various purposes Messages rely on IPFS CIDs and DAG APIs.
	// https://identity.foundation/identity-hub/spec/#messages
	Message struct {
		// Message objects MAY contain a data property, and if present its value MUST
		// be a JSON value of the Message’s data.
		Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
		// Message objects MUST contain a descriptor property, and its value MUST be an
		// object, as defined by the Message Descriptors section of this specification.
		Descriptor *MessageDescriptor `protobuf:"bytes,2,opt,name=descriptor,proto3" json:"descriptor,omitempty"`
		// Message objects MAY contain an attestation property, and if present its value
		// MUST be an object, as defined by the Signed Data section of this specification.
		Attestation *Attestation `protobuf:"bytes,3,opt,name=attestation,proto3" json:"attestation,omitempty"`
		// If a Message object requires signatory and/or permission-evaluated authorization,
		// it MUST include an authorization property, and its value MUST be a [RFC7515]
		// JSON Web Signature composed as follows...
		Authorization *Authorization `protobuf:"bytes,4,opt,name=authorization,proto3" json:"authorization,omitempty"`
	}

	MessageDescriptor struct {
		Method        string `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
		ObjectId      string `protobuf:"bytes,2,opt,name=objectId,proto3" json:"objectId,omitempty"`
		Schema        string `protobuf:"bytes,3,opt,name=schema,proto3" json:"schema,omitempty"`
		DataFormat    string `protobuf:"bytes,4,opt,name=dataFormat,proto3" json:"dataFormat,omitempty"`
		DateCreated   string `protobuf:"bytes,5,opt,name=dateCreated,proto3" json:"dateCreated,omitempty"`
		DatePublished string `protobuf:"bytes,6,opt,name=datePublished,proto3" json:"datePublished,omitempty"`
		DateSort      string `protobuf:"bytes,7,opt,name=dateSort,proto3" json:"dateSort,omitempty"`
		Root          string `protobuf:"bytes,8,opt,name=root,proto3" json:"root,omitempty"`
		Parent        string `protobuf:"bytes,9,opt,name=parent,proto3" json:"parent,omitempty"`
		Cid           string `protobuf:"bytes,10,opt,name=cid,proto3" json:"cid,omitempty"`
	}

	Attestation struct {
		Protected *AttestationProtected `protobuf:"bytes,1,opt,name=protected,proto3" json:"protected,omitempty"`
		Payload   string                `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
		Signature string                `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	}

	AttestationProtected struct {
		Alg string `protobuf:"bytes,1,opt,name=alg,proto3" json:"alg,omitempty"`
		Kid string `protobuf:"bytes,2,opt,name=kid,proto3" json:"kid,omitempty"`
	}

	// todo: https://identity.foundation/identity-hub/spec/#message-descriptors
	Authorization struct {
	}

	Status struct {
		// The status object MUST have a code property, and its value MUST be an integer
		// set to the HTTP Status Code appropriate for the status of the response.
		Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
		// The status object MAY have a message property, and if present its value MUST
		// be a string that describes a terse summary of the status. It is RECOMMENDED
		// that the implementer set the message text to the standard title of the HTTP
		// Status Code, when a title/message has already been defined for that code.
		Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	}

	Reply struct {
		// The object MUST have a messageId property, and its value MUST be the stringified
		// Version 1 CID of the associated message in the Request Object from which it was
		// received.
		MessageId string `protobuf:"bytes,1,opt,name=messageId,proto3" json:"messageId,omitempty"`
		// The object MUST have a status property, and its value MUST be an object composed
		// of the following properties:
		Status *Status `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
		//
		// The object MAY have a entries property if the message request was successful.
		// If present, its value MUST be the resulting message entries returned from the
		// invocation of the corresponding message.
		Entries []string `protobuf:"bytes,3,rep,name=entries,proto3" json:"entries,omitempty"`
	}
)
