package api

type (
	Community struct {
		Id              int64           `json:"Id"`
		Zid             string          `json:"Zid"`
		Name            string          `json:"Name" validate:"required,max=150"`
		OwnerDid        string          `json:"OwnerDid" validate:"required"`
		OwnerUsername   string          `json:"OwnerUsername" validate:"required"`
		Description     string          `json:"Description" validate:"max=250"`
		EscrowAmount    int64           `json:"EscrowAmount" validate:"gte=0,lt=100000"`
		Img             string          `json:"Img"`
		LastActive      int64           `json:"LastActive"`
		PricePerMessage int64           `json:"PricePerMessage" validate:"gte=0,lt=100000"`
		PriceToJoin     int64           `json:"PriceToJoin" validate:"gte=0,lt=100000"`
		Public          Bool            `json:"Public"`
		Created         int64           `json:"Created"`
		Updated         int64           `json:"Updated"`
		Deleted         Bool            `json:"Deleted"`
		Tags            []string        `json:"Tags" validate:"max=5"`
		Users           []UserCommunity `json:"Users"`
	}

	Conversation struct {
		Id           int64     `json:"Id"`
		Zid          string    `json:"Zid"`
		CommunityZid string    `json:"CommunityZid" validate:"required"`
		Text         string    `json:"Text" validate:"required_without=Link"`
		Link         string    `json:"Link" validate:"required_without=Text"`
		Img          string    `json:"Img"`
		Video        string    `json:"Video"`
		Public       Bool      `json:"Public"`
		PublicPrice  int64     `json:"PublicPrice"`
		Created      int64     `json:"Created"`
		Updated      int64     `json:"Updated"`
		Deleted      Bool      `json:"Deleted"`
		Comments     []Comment `json:"Comments"`
	}

	Comment struct {
		Id              int64  `json:"Id"`
		Zid             string `json:"Zid"`
		UserDid         string `json:"UserDid" validate:"required"`
		ConversationZid string `json:"Conversationzid" validate:"required"`
		Text            string `json:"Text" validate:"required_without=Link"`
		Link            string `json:"Link" validate:"required_without=Text"`
		Created         int64  `json:"Created"`
		Updated         int64  `json:"Updated"`
		Deleted         Bool   `json:"Deleted"`
	}

	JoinCommunity struct {
		UserDid      string `json:"user_did" validate:"required"`
		CommunityZid string `json:"community_zid" validate:"required"`
	}

	LeaveCommunity struct {
		UserDid      string `json:"user_did" validate:"required"`
		CommunityZid string `json:"community_zid" validate:"required"`
	}

	User struct {
		Id             int64  `json:"Id"`
		Did            string `json:"Did" validate:"required"`
		Username       string `json:"Username" validate:"required,username,min=6,max=16"`
		Email          string `json:"Email" validate:"omitempty,email"`
		Name           string `json:"Name" validate:"required"`
		Bio            string `json:"Bio"`
		Img            string `json:"Img"`
		PriceToMessage int64  `json:"PriceToMessage"`
		Created        int64  `json:"Created"`
		Updated        int64  `json:"Updated"`
	}

	UserCommunity struct {
		Id           string `json:"id"`
		UserDid      string `json:"user_did"`
		CommunityZid string `json:"community_zid"`
		JoinedDate   int64  `json:"joined_date"`
		LeftDate     int64  `json:"left_date"`
		LeftReason   string `json:"left_reason"`
	}

	Payment struct {
		Id                  int64
		Amount              int64
		Memo                string
		MessageZid          string
		RecipientDid        string
		RecipientNodePubkey string
		RecipientRelayUrl   string
		SenderDid           string
		Status              string
		Type                int64
		Zid                 string
	}
)

// Bool allows 0/1 to also become boolean.
type Bool bool

func (bit *Bool) UnmarshalJSON(b []byte) error {
	txt := string(b)
	*bit = Bool(txt == "1" || txt == "true")
	return nil
}
