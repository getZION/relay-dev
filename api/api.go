package api

type (
	Community struct {
		Name            string   `json:"Name" validate:"required,max=150"`
		Description     string   `json:"Description" validate:"max=250"`
		OwnerDid        string   `json:"OwnerDid" validate:"required"`
		OwnerUsername   string   `json:"OwnerUsername" validate:"required"`
		PricePerMessage int64    `json:"PricePerMessage" validate:"gte=0,lt=100000"`
		PriceToJoin     int64    `json:"PriceToJoin" validate:"gte=0,lt=100000"`
		EscrowAmount    int64    `json:"EscrowAmount" validate:"gte=0,lt=100000"`
		Tags            []string `json:"Tags"`
	}
)
