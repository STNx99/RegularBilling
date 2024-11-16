package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Notification struct {
	NotificationId primitive.ObjectID `bson:"_id"`
	Title          string             `bson:"title"`
	Description    string             `bson:"description"`
	Status         bool               `bson:"status"`
	Type           bool               `bson:"type"`
	CreatedAt      time.Time          `bson:"created_at"`
	BillId         string             `json:"bill_id"`
	UserId         string             `json:"user_id"`
}
