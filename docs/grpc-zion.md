# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [proto/zion/v1/communities.proto](#proto/zion/v1/communities.proto)
    - [Community](#proto.zion.v1.Community)
  
- [proto/zion/v1/conversations.proto](#proto/zion/v1/conversations.proto)
    - [Conversation](#proto.zion.v1.Conversation)
  
- [proto/zion/v1/messages.proto](#proto/zion/v1/messages.proto)
    - [Message](#proto.zion.v1.Message)
  
- [proto/zion/v1/nodeinfo.proto](#proto/zion/v1/nodeinfo.proto)
    - [GetNodeInfoRequest](#proto.zion.v1.GetNodeInfoRequest)
    - [GetNodeInfoResponse](#proto.zion.v1.GetNodeInfoResponse)
  
    - [NodeInfoService](#proto.zion.v1.NodeInfoService)
  
- [proto/zion/v1/payments.proto](#proto/zion/v1/payments.proto)
    - [Payment](#proto.zion.v1.Payment)
  
- [proto/zion/v1/users.proto](#proto/zion/v1/users.proto)
    - [User](#proto.zion.v1.User)
  
- [Scalar Value Types](#scalar-value-types)



<a name="proto/zion/v1/communities.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/zion/v1/communities.proto



<a name="proto.zion.v1.Community"></a>

### Community



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | The community&#39;s unique ID (primary key) |
| zid | [string](#string) |  | Community id unique across Zion

Required; must be uuid v4 |
| name | [string](#string) |  | The name of the community

Required; unique |
| owner_did | [string](#string) |  | Owner DID

Required |
| owner_username | [string](#string) |  | Owner username

Required |
| description | [string](#string) |  | The community&#39;s description or mission statement

Required |
| img | [string](#string) |  | Image URL

Optional |
| tags | [string](#string) | repeated | User-defined tags

Optional |
| price_to_join | [int64](#int64) |  | How much it costs to join

Required |
| price_per_message | [int64](#int64) |  | How much each message costs

Required |
| escrow_amount | [int64](#int64) |  | Escrow amount

Required |
| last_active | [int64](#int64) |  | When last active

Optional |
| public | [bool](#bool) |  | Is this a public community?

Required; defaults true |
| deleted | [bool](#bool) |  | Deleted

Optional, defaults 0 |
| created | [int64](#int64) |  | Created when?

Required |
| updated | [int64](#int64) |  | Updated when?

Optional |





 

 

 

 



<a name="proto/zion/v1/conversations.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/zion/v1/conversations.proto



<a name="proto.zion.v1.Conversation"></a>

### Conversation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | The community&#39;s unique ID (primary key) |
| zid | [string](#string) |  | Conversation ID unique across Zion

Required; must be uuid v4 |
| community_zid | [string](#string) |  | Zion ID of the community this conversation belongs to

Required; must be uuid v4 |
| message | [Message](#proto.zion.v1.Message) |  | The Message that starts the conversation

Required |
| public | [bool](#bool) |  | Is this conversation publicly visible?

Optional; defaults to false |
| public_price | [int64](#int64) |  | Price in sats for non-community members to read, if public = true

Optional; defaults to 0 |





 

 

 

 



<a name="proto/zion/v1/messages.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/zion/v1/messages.proto



<a name="proto.zion.v1.Message"></a>

### Message



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | Message primary ID &amp; key - unique only to this relay |
| zid | [string](#string) |  | Message id unique across Zion

Required; must be uuid v4 |
| user_did | [int64](#int64) |  | DID of user who sent this message

Required |
| receiving_user_did | [int64](#int64) |  | DID of user recipient of this message (if DM)

Optional if conversation_id is non-null, otherwise required |
| conversation_zid | [int64](#int64) |  | Zion ID of conversation of this message (if community/conversation message)

Optional if receiving_user_id is non-null, otherwise required |
| reply_to_message_zid | [int64](#int64) |  | Zion ID of message this message is a reply to

Optional |
| text | [string](#string) |  | Main message body

Optional if link is non-null, otherwise required |
| link | [string](#string) |  | URL link to external piece of content

Optional if text is non-null, otherwise required |
| img | [string](#string) |  | URL for associated image

Optional |
| video | [string](#string) |  | URL for associated video

Optional |





 

 

 

 



<a name="proto/zion/v1/nodeinfo.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/zion/v1/nodeinfo.proto



<a name="proto.zion.v1.GetNodeInfoRequest"></a>

### GetNodeInfoRequest







<a name="proto.zion.v1.GetNodeInfoResponse"></a>

### GetNodeInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pubkey | [string](#string) |  |  |
| balance | [uint64](#uint64) |  |  |





 

 

 


<a name="proto.zion.v1.NodeInfoService"></a>

### NodeInfoService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetNodeInfo | [GetNodeInfoRequest](#proto.zion.v1.GetNodeInfoRequest) | [GetNodeInfoResponse](#proto.zion.v1.GetNodeInfoResponse) | Return node pubkey and balance |

 



<a name="proto/zion/v1/payments.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/zion/v1/payments.proto



<a name="proto.zion.v1.Payment"></a>

### Payment



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | The community&#39;s unique ID (primary key) |
| zid | [string](#string) |  | Payment id unique across Zion

Required; must be uuid v4 |
| sender_did | [string](#string) |  | DID of payment sender

Required |
| recipient_did | [string](#string) |  | DID of payment recipient

Required |
| recipient_node_pubkey | [string](#string) |  | Pubkey of recipient LND node

Required |
| recipient_relay_url | [string](#string) |  | URL of recipient relay

Required |
| status | [string](#string) |  | Status

Required; default &#34;pending&#34; |
| amount | [int64](#int64) |  | Amount in sats

Required |
| type | [int64](#int64) |  | Type of payment: boost, P2P send, community join, stake, etc. |
| memo | [string](#string) |  | Memo describing transaction

Optional |





 

 

 

 



<a name="proto/zion/v1/users.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/zion/v1/users.proto



<a name="proto.zion.v1.User"></a>

### User



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | User primary ID &amp; key |
| did | [string](#string) |  | User&#39;s decentralized identifier (DID), e.g. did:ion:[fingerprint]

Required; unique |
| username | [string](#string) |  | User&#39;s username, unique in Zion

Required; unique |
| name | [string](#string) |  | User&#39;s display name

Optional |
| email | [string](#string) |  | User&#39;s email address for opt-in marketing updates

Optional |
| bio | [string](#string) |  | User&#39;s personal biography that shows on their profile

Optional |
| picture | [string](#string) |  | URL for user&#39;s profile picture

Optional |
| price_to_message | [int64](#int64) |  | How many sats it costs to direct-message this user

Optional - Default to 0 |





 

 

 

 



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

