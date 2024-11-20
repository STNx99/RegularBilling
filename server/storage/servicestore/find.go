package servicestore

import (
	"context"
	"fmt"
	"server/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m *MongoStore) FindService(s models.AddUserService) (models.Service, error) {
	coll := m.db.Collection("services")

	var foundService models.Service
	err := coll.FindOne(context.TODO(), bson.D{
		{Key: "service_name", Value: bson.D{
			{Key: "$regex", Value: s.ServiceName},
			{Key: "$options", Value: "i"},
		}},
	}).Decode(&foundService)
	if err != nil {
		return models.Service{}, fmt.Errorf("No document found")
	}
	fmt.Println(foundService)
	return foundService, nil
}

// func (m *MongoStore) UserHasService(s models.AddUserService) (bool, error) {
// 	userColl := m.db.Collection("users")
// 	serviceColl := m.db.Collection("services")
// 	var user models.User
// 	var service models.Service
//     err := serviceColl.FindOne(context.TODO(), bson.D{
//         {Key: "service_name", Value: s.ServiceName},
//     }).Decode(&service)

//     if err != nil {
//         if err == mongo.ErrNoDocuments {
//             return false, nil
//         }
//         return false, err
//     }
// 	fmt.Println(service)
// 	err = userColl.FindOne(context.TODO(), bson.D{
// 		{Key: "username", Value: s.Username},
// 		{Key: "service_ids", Value: bson.D{
// 			{Key: "$elemMatch", Value: bson.D{
//                 {Key: "service_id", Value: service.ServiceId},
//             }},
// 		}},
// 	}).Decode(&user)
// 	fmt.Println(user)
// 	if err != nil {
//         if err == mongo.ErrNoDocuments {
//             return false, nil
//         }
//         return false, err
//     }
//     return true, nil
// }

func (m *MongoStore) FindServicePrice(serviceId primitive.ObjectID) (float64, error) {
	coll := m.db.Collection("services")
	var price float64
	err := coll.FindOne(context.TODO(), bson.D{{Key: "_id", Value: serviceId}})
	if err != nil {
		return 0, nil
	}
	return price, nil
}


func (m *MongoStore) FindAll() ([]models.Service, error) {
	var services []models.Service
	coll := m.db.Collection("services")

	cursor, err := coll.Find(context.TODO(), bson.D{})

	if err != nil {
		return services, err
	}
	if err := cursor.All(context.TODO(), &services); err != nil {
		return services, err
	}
	return services, nil

}
