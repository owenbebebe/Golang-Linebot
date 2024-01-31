package main

import (
	"github.com/owenbebebe/Golang-Linebot/pkg/bot"
	"github.com/owenbebebe/Golang-Linebot/pkg/conf"
	"github.com/owenbebebe/Golang-Linebot/pkg/settings"
)

func main() {
	cfg := conf.NewConfig(
		conf.WithPath("../cinnox-spec1/chat/app"), // ../chat/app
		conf.WithType("yaml"),
	)
	var botInfo settings.LineBotInfo
	cfg.Unmarshal(&botInfo)
	bot.Init(&botInfo)
}
