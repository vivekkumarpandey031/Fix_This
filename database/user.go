package database

import (
	"context"
	"errors"
	"fmt"
	"golang-project/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	results, err := u.Client.Database(u.Dbname).Collection(u.Collection).Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	var users []models.User
	if err = results.All(context.TODO(), &users); err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.New("No documents found")
	}

	for _, user := range users {
		encryptionErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
		if encryptionErr == nil {
			return &user, encryptionErr
		}
	}

	return nil, errors.New("Encryption error")
}

func (u *User) GetById(ctx context.Context, id string) (*models.User, error) {
	if u.Client == nil {
		return nil, errors.New("nil connection")
	}

	objid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	result := u.Client.Database(u.Dbname).Collection(u.Collection).FindOne(ctx, bson.D{{"_id", objid}})
	user := new(models.User)
	err = result.Decode(user)

	if err != nil {
		return nil, err
	}

	return user, nil

}

func (u *User) TimeAndCost(ctx context.Context, id string, category string) (int64, int64, error) {
	if u.Client == nil {
		return -1, -1, errors.New("nil connection")
	}

	//Query the userproblems database to find cost and time
	problems := []models.UserProblem{}

	results, err := u.Client.Database("userproblemdb").Collection("userproblems").Find(ctx, bson.D{{"uid", id}, {"type", category}})
	if err != nil {
		return -1, -1, err
	}

	if err = results.All(ctx, &problems); err != nil {
		return -1, -1, err
	}

	if len(problems) == 0 {
		return -1, -1, fmt.Errorf("Internal Server Error")
	}

	var cost int64
	var time int64

	for _, problem := range problems {
		c, err := u.Cost(ctx, problem.Problem, category)
		if err != nil {
			return -1, -1, err
		}
		t, err := u.Time(ctx, problem.Problem, category)
		if err != nil {
			return -1, -1, err
		}
		cost += c
		time += t
	}

	return cost, time, nil

}

func (u *User) Cost(ctx context.Context, problem string, category string) (int64, error) {
	if u.Client == nil {
		return -1, errors.New("nil connection")
	}

	p := new(models.Problem)
	err := u.Client.Database("problemdb").Collection("problems").FindOne(ctx, bson.D{{"name", problem}, {"type", category}}).Decode(p)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return -1, err
		} else {
			return -1, fmt.Errorf("Internal server error")

		}
	}
	return p.Cost, nil
}

func (u *User) Time(ctx context.Context, problem string, category string) (int64, error) {
	if u.Client == nil {
		return -1, errors.New("nil connection")
	}

	p := new(models.Problem)
	err := u.Client.Database("problemdb").Collection("problems").FindOne(ctx, bson.D{{"name", problem}, {"type", category}}).Decode(p)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return -1, err
		} else {
			return -1, fmt.Errorf("Internal server error")

		}
	}
	return p.Time, nil
}
