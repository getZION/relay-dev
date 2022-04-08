package api

type (
	Community struct {
		Id              int64           `json:"id,omitempty"`
		Zid             string          `json:"zid"`
		Name            string          `json:"name" validate:"required,max=150"`
		OwnerDid        string          `json:"ownerDid" validate:"required"`
		OwnerUsername   string          `json:"ownerUsername" validate:"required"`
		Description     string          `json:"description" validate:"max=250"`
		EscrowAmount    int64           `json:"escrowAmount" validate:"gte=0,lt=100000"`
		Img             string          `json:"img,omitempty"`
		LastActive      int64           `json:"lastActive,omitempty"`
		PricePerMessage float64         `json:"pricePerMessage" validate:"gte=0,lt=100000"`
		PriceToJoin     float64         `json:"priceToJoin" validate:"gte=0,lt=100000"`
		Public          Bool            `json:"public,omitempty"`
		Created         int64           `json:"created,omitempty"`
		Updated         int64           `json:"updated,omitempty"`
		Deleted         Bool            `json:"deleted,omitempty"`
		Tags            []string        `json:"tags,omitempty" validate:"max=5"`
		Users           []UserCommunity `json:"users,omitempty"`
	}

	Conversation struct {
		Id           int64     `json:"id,omitempty"`
		Zid          string    `json:"zid"`
		CommunityZid string    `json:"communityZid" validate:"required,uuid4"`
		UserDid      string    `json:"userDid" validate:"required"`
		Text         string    `json:"text,omitempty" validate:"required_without=Link"`
		Link         string    `json:"link,omitempty" validate:"required_without=Text"`
		Img          string    `json:"img,omitempty"`
		Video        string    `json:"video,omitempty"`
		Public       Bool      `json:"public,omitempty"`
		PublicPrice  int64     `json:"publicPrice,omitempty"`
		Created      int64     `json:"created,omitempty"`
		Updated      int64     `json:"updated,omitempty"`
		Deleted      Bool      `json:"deleted,omitempty"`
		Comments     []Comment `json:"comments,omitempty"`
	}

	Comment struct {
		Id              int64  `json:"id,omitempty"`
		Zid             string `json:"zid"`
		UserDid         string `json:"userDid" validate:"required"`
		ConversationZid string `json:"conversationzid" validate:"required,uuid4"`
		Text            string `json:"text,omitempty" validate:"required_without=Link"`
		Link            string `json:"link,omitempty" validate:"required_without=Text"`
		Created         int64  `json:"created,omitempty"`
		Updated         int64  `json:"updated,omitempty"`
		Deleted         Bool   `json:"deleted,omitempty"`
	}

	JoinCommunity struct {
		UserDid      string `json:"userDid" validate:"required"`
		CommunityZid string `json:"communityZid" validate:"required,uuid4"`
	}

	LeaveCommunity struct {
		UserDid      string `json:"userDid" validate:"required"`
		CommunityZid string `json:"communityZid" validate:"required,uuid4"`
	}

	User struct {
		Id             int64   `json:"id,omitempty"`
		Did            string  `json:"did" validate:"required"`
		Username       string  `json:"username" validate:"required,username,min=6,max=16"`
		Email          string  `json:"email,omitempty" validate:"omitempty,email"`
		Name           string  `json:"name" validate:"required"`
		Amount         float64 `json:"amount"`
		Bio            string  `json:"bio,omitempty"`
		Img            string  `json:"img,omitempty"`
		PriceToMessage int64   `json:"priceToMessage,omitempty"`
		Created        int64   `json:"created,omitempty"`
		Updated        int64   `json:"updated,omitempty"`
	}

	UserCommunity struct {
		Id           string `json:"id,omitempty"`
		UserDid      string `json:"userDid"`
		CommunityZid string `json:"communityZid"`
		JoinedDate   int64  `json:"joinedDate"`
		LeftDate     int64  `json:"leftDate,omitempty"`
		LeftReason   string `json:"leftReason,omitempty"`
	}

	Payment struct {
		Id                  int64   `json:"id"`
		Zid                 string  `json:"zid"`
		RecipientDid        string  `json:"recipientDid" validate:"required"`
		SenderDid           string  `json:"senderDid" validate:"required"`
		Amount              float64 `json:"amount" validate:"required"`
		Type                string  `json:"type" validate:"oneof=boost p2p community_join stake"`
		Status              string  `json:"status" validate:"oneof=pending completed failed"`
		RelevantType        string  `json:"relevantType" validate:"required_if=Type boost,oneof=Conversation Comment Message"`
		RelevantZid         string  `json:"relevantZid" validate:"required_if=Type boost,uuid4"`
		RecipientNodePubkey string  `json:"recipientNodePubkey" validate:"required"`
		RecipientRelayUrl   string  `json:"recipientRelayUrl" validate:"required"`
		Memo                string  `json:"memo"`
		TypeInt             int16
		StatusInt           int16
	}
)

// Bool allows 0/1 to also become boolean.
type Bool bool

func (bit *Bool) UnmarshalJSON(b []byte) error {
	txt := string(b)
	*bit = Bool(txt == "1" || txt == "true")
	return nil
}
