package handler

import (
	"encoding/json"
	"fmt"

	"github.com/getzion/relay/api/constants"
	"github.com/getzion/relay/api/datastore"
	v1 "github.com/getzion/relay/gen/proto/zion/v1"
)

type PaymentHandler struct {
	DataStore *datastore.Store
}

func (h *PaymentHandler) Execute(data []byte, method string) (interface{}, error) {
	switch method {
	case constants.COLLECTIONS_QUERY:

		return h.DataStore.PaymentService.GetAll()

	case constants.COLLECTIONS_WRITE:

		var payment v1.PaymentORM
		err := json.Unmarshal(data, &payment)
		if err != nil {
			return nil, err
		}

		return h.DataStore.PaymentService.Insert(payment)

	default:
		return nil, fmt.Errorf("unimplemented payment method: %s", method)
	}
}
