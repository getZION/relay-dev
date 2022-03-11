# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [proto/identityhub/v1/identityhub.proto](#proto/identityhub/v1/identityhub.proto)
    - [HubRequestService](#proto.identityhub.v1.HubRequestService)
  
- [proto/identityhub/v1/messages.proto](#proto/identityhub/v1/messages.proto)
    - [Attestation](#proto.identityhub.v1.Attestation)
    - [AttestationProtected](#proto.identityhub.v1.AttestationProtected)
    - [Authorization](#proto.identityhub.v1.Authorization)
    - [Message](#proto.identityhub.v1.Message)
    - [MessageDescriptor](#proto.identityhub.v1.MessageDescriptor)
  
- [proto/identityhub/v1/requests.proto](#proto/identityhub/v1/requests.proto)
    - [Request](#proto.identityhub.v1.Request)
  
- [proto/identityhub/v1/responses.proto](#proto/identityhub/v1/responses.proto)
    - [Reply](#proto.identityhub.v1.Reply)
    - [Response](#proto.identityhub.v1.Response)
    - [Status](#proto.identityhub.v1.Status)
  
- [Scalar Value Types](#scalar-value-types)



<a name="proto/identityhub/v1/identityhub.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/identityhub/v1/identityhub.proto


 

 

 


<a name="proto.identityhub.v1.HubRequestService"></a>

### HubRequestService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Process | [Request](#proto.identityhub.v1.Request) | [Response](#proto.identityhub.v1.Response) |  |

 



<a name="proto/identityhub/v1/messages.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/identityhub/v1/messages.proto



<a name="proto.identityhub.v1.Attestation"></a>

### Attestation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| protected | [AttestationProtected](#proto.identityhub.v1.AttestationProtected) |  |  |
| payload | [string](#string) |  |  |
| signature | [string](#string) |  |  |






<a name="proto.identityhub.v1.AttestationProtected"></a>

### AttestationProtected



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| alg | [string](#string) |  |  |
| kid | [string](#string) |  |  |






<a name="proto.identityhub.v1.Authorization"></a>

### Authorization
todo: https://identity.foundation/identity-hub/spec/#message-descriptors






<a name="proto.identityhub.v1.Message"></a>

### Message
All Identity Hub messaging is transacted via Messages JSON objects.
These objects contain message execution parameters, authorization material,
authorization signatures, and signing/encryption information.
For various purposes Messages rely on IPFS CIDs and DAG APIs.
https://identity.foundation/identity-hub/spec/#messages


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [string](#string) |  | Message objects MAY contain a data property, and if present its value MUST be a JSON value of the Message’s data. |
| descriptor | [MessageDescriptor](#proto.identityhub.v1.MessageDescriptor) |  | Message objects MUST contain a descriptor property, and its value MUST be an object, as defined by the Message Descriptors section of this specification. |
| attestation | [Attestation](#proto.identityhub.v1.Attestation) |  | Message objects MAY contain an attestation property, and if present its value MUST be an object, as defined by the Signed Data section of this specification. |
| authorization | [Authorization](#proto.identityhub.v1.Authorization) |  | If a Message object requires signatory and/or permission-evaluated authorization, it MUST include an authorization property, and its value MUST be a [RFC7515] JSON Web Signature composed as follows... |






<a name="proto.identityhub.v1.MessageDescriptor"></a>

### MessageDescriptor



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| method | [string](#string) |  |  |
| objectId | [string](#string) |  |  |
| schema | [string](#string) |  |  |
| dataFormat | [string](#string) |  |  |
| dateCreated | [string](#string) |  |  |
| datePublished | [string](#string) |  |  |
| dateSort | [string](#string) |  |  |
| root | [string](#string) |  |  |
| parent | [string](#string) |  |  |
| cid | [string](#string) |  |  |





 

 

 

 



<a name="proto/identityhub/v1/requests.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/identityhub/v1/requests.proto



<a name="proto.identityhub.v1.Request"></a>

### Request
Request Objects are JSON object envelopes used to pass messages to Identity Hub
instances.
https://identity.foundation/identity-hub/spec/#request-objects


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| requestId | [string](#string) |  | The Request Object MUST include a requestId property, and its value MUST be an [RFC4122] UUID Version 4 string to identify the request. |
| target | [string](#string) |  | The Request Object MUST include a target property, and its value MUST be the Decentralized Identifier base URI of the DID-relative URL. |
| messages | [Message](#proto.identityhub.v1.Message) | repeated | The Request Object MUST include a messages property, and its value MUST be an array composed of Message objects. |





 

 

 

 



<a name="proto/identityhub/v1/responses.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/identityhub/v1/responses.proto



<a name="proto.identityhub.v1.Reply"></a>

### Reply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| messageId | [string](#string) |  | The object MUST have a messageId property, and its value MUST be the stringified Version 1 CID of the associated message in the Request Object from which it was received. |
| status | [Status](#proto.identityhub.v1.Status) |  | The object MUST have a status property, and its value MUST be an object composed of the following properties: |
| entries | [string](#string) | repeated | The object MAY have a entries property if the message request was successful. If present, its value MUST be the resulting message entries returned from the invocation of the corresponding message. |






<a name="proto.identityhub.v1.Response"></a>

### Response
Responses from Interface method invocations are JSON objects that MUST be
constructed as follows:
https://identity.foundation/identity-hub/spec/#response-objects


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| requestId | [string](#string) |  | The object MUST include an requestId property, and its value MUST be the [RFC4122] UUID Version 4 string from the requestId property of the Request Object it is in response to. |
| status | [Status](#proto.identityhub.v1.Status) |  | The object MUST have a status property, and its value MUST be an object composed of the following properties: |
| replies | [Reply](#proto.identityhub.v1.Reply) | repeated | The object MAY have a replies property, and if present its value MUST be an array of Message Result Objects, which are constructed as follows: |






<a name="proto.identityhub.v1.Status"></a>

### Status



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [int64](#int64) |  | The status object MUST have a code property, and its value MUST be an integer set to the HTTP Status Code appropriate for the status of the response. |
| message | [string](#string) |  | The status object MAY have a message property, and if present its value MUST be a string that describes a terse summary of the status. It is RECOMMENDED that the implementer set the message text to the standard title of the HTTP Status Code, when a title/message has already been defined for that code. |





 

 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

