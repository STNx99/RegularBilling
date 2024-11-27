package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
	ServiceId   primitive.ObjectID `bson:"_id"`
	ServiceName string             `bson:"service_name"`
	Price       float32            `bson:"price"`
	CreatedAt   time.Time          `bson:"created_at"`
	ExpireAt    time.Time          `bson:"expire_at"`
}

type AddUserService struct {
	UserId     primitive.ObjectID `bson:"_id"`
	Service  Service
}

type DeleteUserService struct {
	UserId     primitive.ObjectID `bson:"_id"`
	Service  Service
}
