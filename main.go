package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/owenbebebe/Golang-Linebot/pkg/bot"
	"github.com/owenbebebe/Golang-Linebot/pkg/conf"
	"github.com/owenbebebe/Golang-Linebot/pkg/db"
	"github.com/owenbebebe/Golang-Linebot/pkg/handlr"
	"github.com/owenbebebe/Golang-Linebot/pkg/settings"
)

func main() {
	cfg := conf.NewConfig(
		conf.WithPath("../cinnox-spec1/chat/app"), // ../chat/app
		conf.WithType("yaml"),
	)
	// init line bot
	var botInfo settings.LineBotInfo
	cfg.Unmarshal(&botInfo)
	bot.Init(&botInfo)
	// connect to mongodb
	err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	// set up routes
	r := gin.Default()
	// set up root POST directory to connect to Line Message API
	//give me a POST to root directory with returning 200
	r.POST("/webhook", func(c *gin.Context) {
		handlr.Webhook(c)
	})
	api := r.Group("/api")
	{
		api.POST("/message", func(c *gin.Context) {
			handlr.SendMessages(c, bot.LineBot)
		})
		api.GET("/querymessage", func(c *gin.Context) {
			handlr.QueryMessages(c, db.Collection)
		})

	}
	// running gin on port 8080
	if err := r.Run(":8080"); err != nil {
		log.Println("failed to start gin on port 8080", err)
		panic(err)
	}
}
