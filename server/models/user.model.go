package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	UserId     primitive.ObjectID `bson:"_id"`
	UserName   string             `bson:"username"`
	Email      string             `bson:"email"`
	Password   string             `bson:"password"`
	Credits    float64            `bson:"credits"`
	CreatedAt  time.Time          `bson:"created_at,omitempty"`
	UpdatedAt  time.Time          `bson:"updated_at,omitempty"`
	Bills      []Bill             `bson:"bill_ids"`
	ServiceIds []Service          `bson:"service_ids"`
}
