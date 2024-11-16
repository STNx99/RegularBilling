package bill

import (
	"fmt"
	"server/models"
	"server/storage/servicestore"

	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateNewBill(newBill *models.Bill, user *models.User) (error, models.Bill) {
	return nil, models.Bill{
		BillId:     newBill.BillId,
		BillName:   newBill.BillName,
		UserId:     user.UserId.String(),
		Price:      0,
		Paid:       false,
		Expired:    time.Now().Add(time.Hour * 72),
		CreatedAt:  time.Now(),
		ServiceIds: newBill.ServiceIds,
	}
}
// Calculate the service prices
func CalculatePrice(servicesId []primitive.ObjectID, store *servicestore.MongoStore) float64 {
	var total float64

	for _, serviceId := range servicesId {
		price, err := store.FindServicePrice(serviceId)
		if err != nil {
			fmt.Errorf("No service found")
			continue
		}
		total += price
	}
	return total
}
