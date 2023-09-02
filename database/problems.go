package database

import "go.mongodb.org/mongo-driver/mongo"

type Problem struct {
	Client     *mongo.Client
	Dbname     string
	Collection string
}
