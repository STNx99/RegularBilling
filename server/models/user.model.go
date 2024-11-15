package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	UserId    primitive.ObjectID `bson:"_id"`
	UserName  string             `bson:"username"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	Credits   float64            `bson:"credits"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
}