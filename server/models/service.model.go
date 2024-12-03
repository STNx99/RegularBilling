package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
	ServiceId   primitive.ObjectID `json:"_id" bson:"_id"`
	ServiceName string             `json:"service_name" bson:"service_name"`
	Price       float32            `json:"price" bson:"price"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	ExpireAt    time.Time          `json:"expire_at" bson:"expire_at"`
}

type AddUserService struct {
	UserId  primitive.ObjectID `json:"user_id" bson:"user_id"`
	Service Service            `json:"service" bson:"service"`
}

type DeleteUserService struct {
	UserId  primitive.ObjectID `json:"user_id" bson:"user_id"`
	Service Service            `json:"service" bson:"service"`
}

type ServicesData struct {
	Services     []Service
	ServiceTotal float32 `bson:"total"`
}
