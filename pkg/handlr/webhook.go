package handlr

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v8/linebot"
	"github.com/owenbebebe/Golang-Linebot/pkg/bot"
	"github.com/owenbebebe/Golang-Linebot/pkg/gpt"
	"github.com/owenbebebe/Golang-Linebot/pkg/model"
)

// handle webhook
func Webhook(c *gin.Context) {
	fmt.Println("Successfully called webhook function")
	events, err := bot.LineBot.ParseRequest(c.Request)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			fmt.Println("Invalid Signature", err)
			c.AbortWithStatus(400)
		} else {
			fmt.Println("Error", err)
			c.AbortWithStatus(500)
		}
		return
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			// check for text messages only
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				userId := event.Source.UserID
				profile, err := bot.LineBot.GetProfile(userId).Do()
				// Check for error on gettting progile error
				if err != nil {
					fmt.Println("Get profile err:", err)
				}
				// Create message object
				userMessage := &model.Message{
					MessageId: message.ID,
					Message:   message.Text,
					UserId:    profile.UserID,
					UserName:  profile.DisplayName,
				}
				// Create message in database
				model.CreateMessage(userMessage)
				// Send message to GPT-3
				response, err := gpt.GetGPT3Response(message.Text)
				if err != nil {
					fmt.Println("Error:", err)
				} else {
					fmt.Println("Response:", response)
				}
				// Send message back to user
				if _, err := bot.LineBot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(response)).Do(); err != nil {
					fmt.Println("Error:", err)
				}

			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})

}
