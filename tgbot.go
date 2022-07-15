package main

import (
	"telegram-bot-api"
	//"github.com/lib/pq"
	"log"
)

const TGBotTokken = "5537584496:AAER6fpsUAOvMfQGWBmo6WLjs1HS5y9MFMg"

func main() {
	bot, err := tgbotapi.NewBotAPI(TGBotTokken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
}
