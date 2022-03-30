package handler

import (
	"encoding/json"
	"fmt"

	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/constants"
)

type PaymentHandler struct {
	storage api.Storage
}

func InitPaymentHandler(storage api.Storage) *PaymentHandler {
	return &PaymentHandler{
		storage: storage,
	}
}

func (h *PaymentHandler) Execute(data []byte, method string) (interface{}, error) {
	switch method {
	case constants.COLLECTIONS_QUERY:

		return h.storage.GetPayments()

	case constants.COLLECTIONS_WRITE:

		var payment api.Payment
		err := json.Unmarshal(data, &payment)
		if err != nil {
			return nil, err
		}

		err = h.storage.InsertPayment(&payment)
		if err != nil {
			return nil, err
		}

		return payment, nil

	default:
		return nil, fmt.Errorf("unimplemented payment method: %s", method)
	}
}
