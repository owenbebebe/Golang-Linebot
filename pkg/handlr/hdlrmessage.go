package handlr

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v8/linebot"
)

// SendMessages API function that sends message to specific user ID
func SendMessages(c *gin.Context, bot *linebot.Client) {
	userId := "<specific userID>"
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
