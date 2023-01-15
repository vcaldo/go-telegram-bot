package main

import (
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	qb "github.com/vcaldo/go-telegram-bot/qbitorrent"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "help":
			msg.Text = "I understand /sayhi, /status and /torrent"
		case "sayhi":
			msg.Text = "Hi :)"
		case "status":
			msg.Text = "I'm ok. " + strings.Split(update.Message.Text, " ")[1]
		case "torrent":
			a := qb.GetTorrents()
			// a := qb.Auth()
			msg.Text = string(a)
		case "add":
			torrentUrl := strings.Split(update.Message.Text, " ")[1]
			qb.AddTorrent(torrentUrl)

		default:
			msg.Text = "I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
