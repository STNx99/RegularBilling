package userstore

import (
	"context"
	"fmt"
	"server/models"
)


func (m *MongoStore) CreateUser(newUser *models.User) error{
	coll := m.db.Client().Database("database").Collection("users")
	existingUser, err := m.CheckUser(newUser.Email)
	if err == nil && existingUser != nil {
		return fmt.Errorf("user already exists")
	}
	_, err = coll.InsertOne(context.TODO(), newUser)
	return err
} 