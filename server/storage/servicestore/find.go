package servicestore

import (
	"context"
	"server/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (m *MongoStore) FindServicePrice(serviceId primitive.ObjectID) (float64, error){
	coll := m.db.Client().Database("database").Collection("services")
	var price float64
	err := coll.FindOne(context.TODO(), bson.D{{Key: "_id",Value: serviceId}})
	if err != nil{
		return 0, nil
	}
	return price, nil
}
