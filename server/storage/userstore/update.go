package userstore

import (
	"context"
	"server/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (m *MongoStore) AddUserServices(name string, service models.Service) error {
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
func (m *MongoStore) DeleteUSerServices(name string, service models.Service) error {
	_, err := m.db.Collection("users").UpdateOne(
		context.TODO(),
		bson.M{"username": name},
		bson.M{"$pull": bson.M{"service_ids": service}},
	)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoStore) UpdateUserServices(name string, service models.Service) error {
	
	_, err := m.db.Collection("users").UpdateOne(
		context.TODO(),
		bson.M{"username": name, "service_ids._id": service.ServiceId},
		bson.M{
			"$set": bson.M{
				"service_ids.$.service_name": service.ServiceName,
				"service_ids.$.price":       service.Price,
				"service_ids.$.created_at":  service.CreatedAt,
			},
		},
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
