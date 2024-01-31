package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// declare message struct
type UserMessage struct {
	UserId  string `bson:"user_id"`
	Message string `bson:"message"`
}

var client *mongo.Client

// connect to mongodb with localhost set on Docker

func Connect() error {
	var err error
	client, err = mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return err
	}
	// ping the server if properly connected
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}
	log.Println("Ping succeeded")
	return nil
}

// save user message to mongodb
func SaveUserMessage(msg UserMessage) error {
	collection := client.Database("test").Collection("user_messages")
	_, err := collection.InsertOne(context.Background(), msg)
	return err
}

// query user message from mongodb that match with specific user id
func QueryUserMessages(userID string) ([]UserMessage, error) {
	collection := client.Database("test").Collection("user_messages")
	filter := bson.M{"userid": userID}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	var messages []UserMessage
	if err = cursor.All(context.Background(), &messages); err != nil {
		return nil, err
	}
	return messages, nil
}
