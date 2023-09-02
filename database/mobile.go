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

func (m *Mobile) Insert(ctx context.Context, mobile *models.Mobile) (any, error) {
	if m.Client == nil {
		return nil, fmt.Errorf("nil connection")
	}
	//Insert into the database
	result, err := m.Client.Database(m.Dbname).Collection(m.Collection).InsertOne(ctx, mobile)
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
