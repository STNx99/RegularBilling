package userstore

import (
	"context"
	"server/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (m *MongoStore) UpdateUserServices(name string, service models.Service) error {
	_, err := m.db.Collection("users").UpdateOne(
		context.TODO(),
		bson.M{"username": name},
		bson.M{"$push": bson.M{"service_ids": service}},
	)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoStore) UpdateUserBill(name string, bill models.Bill) error {
	_, err := m.db.Collection("users").UpdateOne(
		context.TODO(),
		bson.M{"username": name},
		bson.M{"$push": bson.M{"bill_ids": bill}},
	)
	if err != nil {
		return err
	}
	return nil
}
