package userstore

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func (m *MongoStore) CheckUser(email string) (*User, error)  {
	coll := m.db.Client().Database("database").Collection("users")

	var foundUser User
	err := coll.FindOne(context.TODO(), bson.D{{Key:"email", Value: email}}).Decode(&foundUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		return nil, err
	}
	return &foundUser, nil
}
func (m *MongoStore) FindUser(email string, password string) (error){
	coll := m.db.Client().Database("database").Collection("users")

	var foundUser User
	err := coll.FindOne(context.TODO(), bson.D{{Key:"email", Value: email}}).Decode(&foundUser)
	if err != nil{
		return mongo.ErrNoDocuments
	}
	if err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password)); err != nil{
		return err 
	}
	return nil
}