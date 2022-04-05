package common

import (
	"github.com/getzion/relay/api"
	"github.com/google/uuid"
)

func (c *Connection) GetPayments() ([]api.Payment, error) {
	// var payments []v1.PaymentORM
	// result := s.connection.DB.Find(&payments)
	// if result.Error != nil {
	// 	return nil, result.Error
	// }
	// return payments, nil
	return nil, nil
}

func (c *Connection) InsertPayment(payment *api.Payment) error {

	payment.Zid = uuid.NewString()

	// result := s.connection.DB.Create(&payment)
	// if result.Error != nil {
	// 	return nil, result.Error
	// }

	// return &payment, nil
	return nil
}
