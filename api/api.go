package api

type (
	Community struct {
		Id              int64
		Zid             string
		Name            string   `json:"Name" validate:"required,max=150"`
		Description     string   `json:"Description" validate:"max=250"`
		OwnerDid        string   `json:"OwnerDid" validate:"required"`
		OwnerUsername   string   `json:"OwnerUsername" validate:"required"`
		PricePerMessage int64    `json:"PricePerMessage" validate:"gte=0,lt=100000"`
		PriceToJoin     int64    `json:"PriceToJoin" validate:"gte=0,lt=100000"`
		EscrowAmount    int64    `json:"EscrowAmount" validate:"gte=0,lt=100000"`
		Img             string   `json:"Img"`
		Created         int64    `json:"Created"`
		LastActive      int64    `json:"LastActive"`
		Public          bool     `json:"Public"`
		Tags            []string `json:"Tags" validate:"max=5"`
		Users           []string
	}

	Conversation struct {
		Id          int64
		Zid         string
		Text        string
		Link        string
		Img         string
		Video       string
		Public      bool
		PublicPrice int64
		Created     int64
		Updated     int64
		Deleted     bool
	}

	JoinCommunity struct {
		UserDid      string `json:"user_did" validate:"required"`
		CommunityZid string `json:"community_zid" validate:"required"`
	}

	LeaveCommunity struct {
		UserDid      string `json:"user_did" validate:"required"`
		CommunityZid string `json:"community_zid" validate:"required"`
	}
)
