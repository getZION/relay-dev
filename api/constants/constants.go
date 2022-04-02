package constants

const (
	SCHEMA_COMMUNITY       = "https://schema.org/Organization"
	SCHEMA_CONVERSATION    = "https://schema.org/Conversation"
	SCHEMA_MESSAGE         = "https://schema.org/SocialMediaPosting"
	SCHEMA_PAYMENT         = "https://schema.org/Invoice"
	SCHEMA_PERSON          = "https://schema.org/Person"
	SCHEMA_JOIN_COMMUNITY  = "https://schema.org/JoinAction"
	SCHEMA_LEAVE_COMMUNITY = "https://schema.org/LeaveAction"
	SCHEMA_COMMENT         = "https://schema.org/Comment"
	SCHEMA_KICK_USER       = "https://schema.org/Organization/KickUser"

	COLLECTIONS_QUERY  = "CollectionsQuery"
	COLLECTIONS_WRITE  = "CollectionsWrite"
	COLLECTIONS_COMMIT = "CollectionsCommit"
	COLLECTIONS_DELETE = "CollectionsDelete"

	THREADS_QUERY  = "ThreadsQuery"
	THREADS_CREATE = "ThreadsCreate"
	THREADS_REPLY  = "ThreadsReply"
	THREADS_CLOSE  = "ThreadsClose"
	THREADS_DELETE = "ThreadsDelete"

	PERMISSIONS_REQUEST = "PermissionsRequest"
	PERMISSIONS_QUERY   = "PermissionsQuery"
	PERMISSIONS_GRANT   = "PermissionsGrant"
	PERMISSIONS_REVOKE  = "PermissionsRevoke"
)
