package database

import (
	"context"
	"errors"
	"fmt"
	"golang-project/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Client     *mongo.Client
	Dbname     string
	Collection string
}

func (u *User) Insert(ctx context.Context, user *models.User) (any, error) {
	if u.Client == nil {
		return nil, fmt.Errorf("nil connection")
	}
	//Insert into the database
	result, err := u.Client.Database(u.Dbname).Collection(u.Collection).InsertOne(ctx, user)
	return result.InsertedID, err
}

func (u *User) Find(ctx context.Context, data *models.User) (*models.User, error) {
	if u.Client == nil {
		return nil, fmt.Errorf("nil connection")
	}

	//Find User
	filter := bson.D{{"name", data.Name}, {"email", data.Email}}
	results,err := u.Client.Database(u.Dbname).Collection(u.Collection).Find(context.TODO(),filter)
	if err!=nil{
		return nil,err
	}

	var users []models.User
	if err=results.All(context.TODO(),&users);err!=nil{
		return nil,err
	}

	for _,user:=range users{
		encryptionErr:=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(data.Password))
		if encryptionErr==nil{
			return data,nil
		}

	}
	
	return nil,errors.New("Encryption error")
}
