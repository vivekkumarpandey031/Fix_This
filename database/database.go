package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



func GetConnection(dsn string) (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1) //which server ur using
	opts := options.Client().ApplyURI(dsn).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	return mongo.Connect(context.TODO(), opts) //mongo.Connect to connect to the database
}

func Ping(client *mongo.Client,dbname string) error {
	// Send a ping to confirm a successful connection
	if err := client.Database(dbname).RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		return err
	}
	return nil
}