package userstore

import (
	"context"
	"fmt"
	"server/models"
)


func (m *MongoStore) CreateUser(newUser *models.User) error {
	coll := m.db.Collection("users")

	// Check if the user already exists
	exists, err := m.CheckUser(newUser, coll)
	if err != nil {
		return fmt.Errorf("error checking user: %w", err)
	}
	if exists {
		return fmt.Errorf("user already exists")
	}

	// Insert the new user into the database
	_, err = coll.InsertOne(context.TODO(), newUser)
	if err != nil {
		return fmt.Errorf("error inserting user: %w", err)
	}

	return nil
}