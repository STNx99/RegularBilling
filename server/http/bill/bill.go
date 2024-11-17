package bill

import (
	"fmt"
	"server/models"
	"server/storage/servicestore"

	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateNewBill(newBill *models.Bill, user *models.User) (error, models.Bill) {
	return nil, models.Bill{
		BillId:    newBill.BillId,
		BillName:  newBill.BillName,
		UserId:    user.UserId,
		Price:     0,
		Paid:      false,
		Expired:   time.Now().Add(time.Hour * 72),
		CreatedAt: time.Now(),
	}
}

// Calculate the service prices
func CalculatePrice(servicesId []primitive.ObjectID, store *mongo.Database) float64 {
	var total float64
	db := servicestore.NewMongoStore(store)
	for _, serviceId := range servicesId {
		price, err := db.FindServicePrice(serviceId)
		if err != nil {
			fmt.Errorf("No service found")
			continue
		}
		total += price
	}
	return total
}
