package smtp

import (
	"fmt"
	"server/models"
	"server/storage/billstore"
	"server/storage/userstore"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func newSMTPBill(userId primitive.ObjectID) (error, models.Bill) {
	return nil, models.Bill{
		BillId:    primitive.NewObjectID(),
		BillName:  "Thanh toan",
		UserId:    userId,
		Price:     0,
		Paid:      false,
		Expired:   time.Now().Add(time.Hour * 72),
		CreatedAt: time.Now(),
	}
}

func CalculateUserBill(store *mongo.Database) {
	userdb := userstore.NewMongoStore(store)
	
	users, err := userdb.FindAll()
	if err != nil {
		fmt.Errorf("Error getting users" + err.Error())
	}
	for _, user := range users {
		err, newUserBill := newSMTPBill(user.UserId)
		if err != nil {
			fmt.Println("Error creating newSMTPBill" + err.Error())
			continue
		}
		err = billstore.NewMongoStore(store).CreateBill(newUserBill)
		if err != nil {
			fmt.Println("Error creating new Bill" + err.Error())
			continue
		}
	}
}
