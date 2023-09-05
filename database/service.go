package database

import (
	"context"
	"errors"
	"fmt"
	"golang-project/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	Client *mongo.Client
}

func (s *Service) TimeAndCost(ctx context.Context, id string, category string) (int64, int64, error) {
	if s.Client == nil {
		return -1, -1, errors.New("nil connection")
	}

	//Query the userproblems database to find cost and time
	problems := []models.UserProblem{}

	results, err := s.Client.Database("userproblemdb").Collection("userproblems").Find(ctx, bson.D{{"uid", id}, {"type", category}})
	if err != nil {
		return -1, -1, err
	}

	if err = results.All(ctx, &problems); err != nil {
		return -1, -1, err
	}

	if len(problems) == 0 {
		return 0, 0,nil
	}

	var cost int64
	var time int64

	for _, problem := range problems {
		c, err := s.Cost(ctx, problem.Problem, category)
		if err != nil {
			return -1, -1, err
		}
		t, err := s.Time(ctx, problem.Problem, category)
		if err != nil {
			return -1, -1, err
		}
		cost += c
		time += t
	}

	return cost, time, nil

}

func (s *Service) Cost(ctx context.Context, problem string, category string) (int64, error) {
	if s.Client == nil {
		return -1, errors.New("nil connection")
	}

	p := new(models.Problem)
	err := s.Client.Database("problemdb").Collection("problems").FindOne(ctx, bson.D{{"name", problem}, {"type", category}}).Decode(p)
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

func (s *Service) Time(ctx context.Context, problem string, category string) (int64, error) {
	if s.Client == nil {
		return -1, errors.New("nil connection")
	}

	p := new(models.Problem)
	err := s.Client.Database("problemdb").Collection("problems").FindOne(ctx, bson.D{{"name", problem}, {"type", category}}).Decode(p)
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