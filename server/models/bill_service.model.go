package models

import "time"

type ServiceBill struct {
	BillId    string    `json:"bill_id"`
	ServiceId string    `json:"service_id"`
	Price     float64   `json:"price"`
	Expired   time.Time `json:"expired"`
	CreatedAt time.Time `json:"created_at"`
}
