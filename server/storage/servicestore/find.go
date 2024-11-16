package servicestore

import (
	"context"
	"server/models"

	"go.mongodb.org/mongo-driver/bson"
)


func (m *MongoStore) FindService(service models.Service) (models.Service, error) {
	coll := m.db.Client().Database("database").Collection("services")
	var foundService models.Service
	err := coll.FindOne(context.TODO(), bson.D{{Key: "service_name", Value: service.ServiceName}}).Decode(&foundService)
	if err != nil{
		return models.Service{}, err
	}
	return foundService,  nil
}