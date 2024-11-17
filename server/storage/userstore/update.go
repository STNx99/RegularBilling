package userstore

import (
	"context"
	"server/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (m *MongoStore) UpdateUserBill(name string, billId primitive.ObjectID) error {
	_, err := m.db.Collection("users").UpdateOne(
		context.TODO(),
		bson.M{"username": name},
		bson.M{"$push": bson.M{"bill_ids": billId}},
	)
	if err != nil {
		return err
	}
	return nil
}
