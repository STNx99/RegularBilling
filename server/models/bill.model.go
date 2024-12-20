package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Bill struct {
	BillId    primitive.ObjectID `bson:"_id"`
	BillName  string             `bson:"bill_name"`
	UserId    string             `bson:"user_id"`
	ServiceIds []primitive.ObjectID `bson:"service_ids"` 
	Price     float64            `bson:"price"`
	Paid      bool               `bson:"paid"`
	Expired   time.Time          `bson:"expired"`
	CreatedAt time.Time          `bson:"created_at"`
}
