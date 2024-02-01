package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var Collection *mongo.Collection

var dbName = "linebot"
var collectionName = "messages"

// connect to mongodb with localhost set on Docker

func Connect() error {
	fmt.Println("Connecting to mongodb")
	var err error
	client, err = mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return err
	}
	Collection = client.Database(dbName).Collection(collectionName)
	// ping the server if properly connected
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}
	log.Println("Ping succeeded")
	return nil
}
