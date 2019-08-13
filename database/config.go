package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongo() (database *mongo.Database, err error) {

	// Create Connection with MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://admin:admin123@ds261817.mlab.com:61817/rlist"))
	if err != nil {
		log.Fatalf("Error to Connect Database %s", err.Error())
		return nil, err
	}

	// if operation connect to client more than 2 sec, this operation will be stopped
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Error to Connect Client %+s", err.Error())
		return nil, err
	}

	database = client.Database("rlist")
	return database, nil
}
