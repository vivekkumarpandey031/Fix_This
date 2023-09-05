package database

import (
	"context"
	"errors"
	"fmt"
	"golang-project/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mobile struct {
	Client     *mongo.Client
	Dbname     string
	Collection string
}

func (m *Mobile) Insert(ctx context.Context, userproblem *models.UserProblem) (any, error) {
	if m.Client == nil {
		return nil, fmt.Errorf("nil connection")
	}

	//Check if the problem exists in the problem database
	problem := new(models.Problem)
	err := m.Client.Database("problemdb").Collection("problems").FindOne(ctx, bson.D{{"type", "mobile"}, {"name", userproblem.Problem}}).Decode(problem)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return nil, err
		} else {
			return nil, fmt.Errorf("Internal server error")
		}
	}

	//Insert into the problems database
	result, err := m.Client.Database("userproblemdb").Collection("userproblems").InsertOne(ctx, userproblem)
	return result.InsertedID, err
}

func (m *Mobile) GetAll(ctx context.Context) ([]models.Mobile, error) {
	if m.Client == nil {
		return nil, fmt.Errorf("nil connection")
	}
	//Get All Mobiles  into the database
	results, err := m.Client.Database(m.Dbname).Collection(m.Collection).Find(ctx, bson.D{})
	documents := []models.Mobile{}
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


func (m *Mobile) Find(ctx context.Context,brand string,model string) (error) {
	if m.Client == nil {
		return fmt.Errorf("nil connection")
	}
	//Find Mobile  into the database
	
	var document models.Mobile
	err := m.Client.Database(m.Dbname).Collection(m.Collection).FindOne(ctx, bson.D{{"model",model},{"brand",brand}}).Decode(&document)
	if err != nil {
		return err
	}
	
	return nil
}