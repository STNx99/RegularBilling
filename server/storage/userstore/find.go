package userstore

import (
	"context"
	"fmt"
	"server/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)
func(m *MongoStore) CheckUser(newUser *models.User, coll *mongo.Collection) (bool, error){
	exist, err := m.CheckUserMail(newUser.Email, coll)
	if err != nil{
		return true, err
	}
	if exist != nil {
		return true, nil
	} 
	exist, err = m.CheckUserName(newUser.UserName, coll)
	if err != nil{
		return true, err
	}
	if exist != nil{
		return true, nil
	} 
	
	return false, nil
}

func (m *MongoStore) CheckUserMail(email string, coll *mongo.Collection) (*models.User, error)  {
	var foundUser models.User
	err := coll.FindOne(context.TODO(), bson.D{{Key:"email", Value: email}}).Decode(&foundUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	fmt.Println(foundUser)
	return &foundUser, err
}

func (m *MongoStore) CheckUserName(name string, coll *mongo.Collection) (*models.User, error){
	var foundUser models.User
	err := coll.FindOne(context.TODO(), bson.D{{Key:"username", Value: name}}).Decode(&foundUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	fmt.Println(foundUser)
	return &foundUser, nil
}

func (m *MongoStore) FindUser(email string, password string) (error){
	coll := m.db.Collection("users")

	var foundUser models.User
	err := coll.FindOne(context.TODO(), bson.D{{Key:"email", Value: email}}).Decode(&foundUser)
	if err != nil{
		return mongo.ErrNoDocuments
	}
	if err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password)); err != nil{
		return err 
	}
	return nil
}
func (m *MongoStore) FindAll() ([]models.User, error){
	coll := m.db.Collection("users")
	var users []models.User
	
	cursor, err := coll.Find(context.TODO(), bson.D{})

	if err != nil {
		return users, err
	}
	if err := cursor.All(context.TODO(), &users); err != nil{
		return users, err
	}
	return users, nil
}