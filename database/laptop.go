package database

import (
	"context"
	"errors"
	"fmt"
	"golang-project/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Laptop struct {
	Client     *mongo.Client
	Dbname     string
	Collection string
}

func (l *Laptop) GetAll(ctx context.Context) ([]models.Laptop, error) {
	if l.Client == nil {
		return nil, fmt.Errorf("nil connection")
	}
	//Get All Laptops  into the database
	results, err := l.Client.Database(l.Dbname).Collection(l.Collection).Find(ctx, bson.D{})
	documents := []models.Laptop{}
	if err != nil {
		return documents, err
	}

	if err = results.All(ctx, &documents); err != nil {
		return documents, err
	}

	if len(documents) == 0 {
		return documents, errors.New("No documents found")
	}

	return documents, nil
}

func (l *Laptop) Insert(ctx context.Context, userproblem *models.UserProblem) (any, error) {
	if l.Client == nil {
		return nil, fmt.Errorf("nil connection")
	}
	//Check if the problem exists in the problem database
	problem := new(models.Problem)
	err := l.Client.Database("problemdb").Collection("problems").FindOne(ctx, bson.D{{"type", "mobile"}, {"name", userproblem.Problem}}).Decode(problem)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return nil, err
		} else {
			return nil, fmt.Errorf("Internal server error")
		}
	}
	//Insert into the problems database
	result, err := l.Client.Database("userproblemdb").Collection("userproblems").InsertOne(ctx, userproblem)
	return result.InsertedID, err
}
