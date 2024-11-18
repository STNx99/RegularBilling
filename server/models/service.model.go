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
}

type AddUserService struct {
	ServiceName string `bson:"service_name"`
	Username    string `bson:"username"`
}

type DeleteUserService struct{
	Username string `bson:"username"`
	Service Service
}