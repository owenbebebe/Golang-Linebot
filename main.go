package test

import (
	"context"
	"fmt"
	"log"

	"github.com/your-username/linebot/linebot"
	"github.com/your-username/mongodb/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Initialize linebot package
	bot, err := linebot.New("<YOUR_CHANNEL_SECRET>", "<YOUR_CHANNEL_ACCESS_TOKEN>")
	if err != nil {
		log.Fatal(err)
	}

	// Initialize mongodb package
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Use the bot and client objects for further operations

	// Example: Print the bot and client objects
	fmt.Println(bot)
	fmt.Println(client)
}
