package bot

import (
	"fmt"
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/v8/linebot"
	"github.com/owenbebebe/Goland-Linebot/pkg/setting"
)

func Init(s *setting.LineBotInfo) {
	bot, err := linebot.New(
		s.ChannelSecret,
		s.ChannelAccessToken,
	)
	if err != nil {
		log.Fatal(err)
	}
	// sending api request to line server to check for response
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:

					// if linebot receives text from user, echo back the same message
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	})
	// line bot operating on localhose:8080
	fmt.Println("LINE BOT operating ...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
