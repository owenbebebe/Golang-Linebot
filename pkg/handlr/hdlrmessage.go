package handlr

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v8/linebot"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SendMessages API function that sends message to specific user ID
func SendMessages(c *gin.Context, bot *linebot.Client) {
	userId := "<specify user ID here>"
	pushMessage := c.DefaultQuery("message", "send empty message")
	message := linebot.NewTextMessage(pushMessage)

	//error handling
	_, err := bot.PushMessage(userId, message).Do()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		log.Print(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "send message successfully!"})
}

// get all of the messages sent by specific user ID
func QueryMessages(c *gin.Context, collection *mongo.Collection) {
	userId := "<specify user ID here>"

	// Query messages from MongoDB
	filter := bson.M{"userId": userId}
	findOptions := options.Find()
	cursor, err := collection.Find(c.Request.Context(), filter, findOptions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		log.Print(err)
		return
	}
	defer cursor.Close(c.Request.Context())

	// Iterate over the cursor and retrieve messages
	var messages []bson.M
	if err := cursor.All(c.Request.Context(), &messages); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		log.Print(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"messages": messages})
}
