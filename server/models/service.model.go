package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Service struct {
	ServiceId   primitive.ObjectID `bson:"_id"`
	ServiceName string             `bson:":service_name"`
	CreatedAt   time.Time          `bson:"created_at"`
}
