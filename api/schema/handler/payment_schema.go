package handler

import (
	"fmt"

	"github.com/getzion/relay/api/constants"
	"github.com/getzion/relay/api/datastore"
)

type PaymentHandler struct {
	DataStore *datastore.Store
}

func (h *PaymentHandler) Execute(data []byte, method string) (interface{}, error) {
	switch method {
	case constants.COLLECTIONS_QUERY:

		return h.DataStore.PaymentService.GetAll()

	case constants.COLLECTIONS_WRITE:

		return h.DataStore.PaymentService.Insert(data)

	default:
		return nil, fmt.Errorf("unimplemented payment method: %s", method)
	}
}
