package bot

import (
	"log"

	"github.com/line/line-bot-sdk-go/v8/linebot"
	"github.com/owenbebebe/Golang-Linebot/pkg/settings"
)

var LineBot *linebot.Client

func Init(s *settings.LineBotInfo) {
	var err error
	// create new line bot that reads Channel Secret and ChannelAccessToken from the file
	LineBot, err = linebot.New(
		s.ChannelSecret,
		s.ChannelAccessToken,
	)
	if err != nil {
		log.Fatal(err)
		panic("line bot init error")
	}
	log.Println("Line bot init success")
}
