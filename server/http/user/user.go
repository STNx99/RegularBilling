package user

import (
	"log"
	"server/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)



func hashPassword(password string) string{
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Fatal("Error hashing password")
	}
	return string(encryptedPassword)
}


func CreateNewUser(newUser *models.User) (error, models.User){
	if newUser.Bills == nil{
		newUser.Bills = []models.Bill{}
	}
	if newUser.ServiceIds == nil{
		newUser.ServiceIds = []models.Service{}
	}
	userId := primitive.NewObjectID()
	return nil, models.User{
		UserId: userId,
		UserName: newUser.UserName,
		Password: hashPassword(newUser.Password),
		Email: newUser.Email,
		Credits: 0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Bills: newUser.Bills,
		ServiceIds: newUser.ServiceIds,
	}
}
