package main

import (
	"log"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := telegram.NewBotAPI("957372446:AAFj_uWJwHgc_lMnX8ZBIfFqE5nccgrx1oQ")
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s\n", bot.Self.UserName)

	u := telegram.NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := telegram.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
	}
}
