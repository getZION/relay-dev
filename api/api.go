package api

type (
	Community struct {
		Id              int64    `json:"Id"`
		Zid             string   `json:"Zid"`
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
		Deleted         bool     `json:"Deleted"`
		Updated         int64    `json:"Updated"`
		Tags            []string `json:"Tags" validate:"max=5"`
		Users           []string `json:"Users"`
	}

	//todo: write validation tests
	Conversation struct {
		Id           int64     `json:"Id"`
		Zid          string    `json:"Zid"`
		CommunityZid string    `json:"CommunityZid" validate:"required"`
		Text         string    `json:"Text" validate:"required_without=Link"`
		Link         string    `json:"Link" validate:"required_without=Text"`
		Img          string    `json:"Img"`
		Video        string    `json:"Video"`
		Public       bool      `json:"Public"`
		PublicPrice  int64     `json:"PublicPrice"`
		Created      int64     `json:"Created"`
		Updated      int64     `json:"Updated"`
		Deleted      bool      `json:"Deleted"`
		Comments     []Comment `json:"Comments"`
	}

	//todo: write validation tests
	Comment struct {
		Id              int64  `json:"Id"`
		Zid             string `json:"Zid"`
		UserDid         string `json:"UserDid" validate:"required"`
		ConversationZid string `json:"Conversationzid" validate:"required"`
		Text            string `json:"Text" validate:"required_without=Link"`
		Link            string `json:"Link" validate:"required_without=Text"`
		Created         int64  `json:"Created"`
		Updated         int64  `json:"Updated"`
		Deleted         bool   `json:"Deleted"`
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
		Bio            string `json:"Bio"`
		Created        int64  `json:"Created"`
		Did            string `json:"Did" validate:"required"`
		Email          string `json:"Email" validate:"omitempty,email"`
		Id             int64  `json:"Id"`
		Name           string `json:"Name" validate:"required"`
		Picture        string `json:"Price"`
		PriceToMessage int64  `json:"PriceToMessage"`
		Updated        int64  `json:"Updated"`
		Username       string `json:"Username" validate:"required,username,min=6,max=16"`
	}
)
