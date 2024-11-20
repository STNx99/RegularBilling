package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Bill struct {
	BillId    primitive.ObjectID `bson:"_id"`
	BillName  string             `bson:"bill_name"`
	UserId    primitive.ObjectID `bson:"user_id"`
	Price     float64            `bson:"price"`
	Paid      bool               `bson:"paid"`
	Expired   time.Time          `bson:"expired"`
	CreatedAt time.Time          `bson:"created_at"`
}

type MonthlyYearlyBill struct {
	Bills     []Bill
	YearTotal float64 `bson:"year_total"`
}
