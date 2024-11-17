package billstore

import (
	"context"
	"server/models"

)

func (m *MongoStore) CreateBill(newBill models.Bill) error{
	coll := m.db.Collection("bills")
	_, err := coll.InsertOne(context.TODO(), newBill)
	if err != nil{
		return err
	}
	return nil
}