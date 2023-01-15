package main

import (
	"errors"
	"log"
	"os"
	"time"

	qbitorrent "bot/commands/qbitorrent"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// NOTE: Uncomment after configuring `NEW_RELIC_APP_NAME` and `NEW_RELIC_LICENSE_KEY`
	newrelicApp, err := StartNewRelicAgent()
	if err != nil {
		log.Fatalf("Failed to start NewRelic Agent: %v", err)
	}

	if err := newrelicApp.WaitForConnection(5 * time.Second); err != nil {
		log.Fatalf("Failed to connect to New Relic: %v", err)
	}

	// NOTE: Replace `nil` with `newrelicApp` once agent-related vars are implemented
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
			msg.Text = "I'm ok."
		case "torrent":
			msg.Text = "get torrents"
			qbitorrent.auth()
		default:
			msg.Text = "I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
func StartNewRelicAgent() (newrelic.Application, error) {
	appName := os.Getenv("NEW_RELIC_APP_NAME")
	if appName == "" {
		return nil, errors.New("NEW_RELIC_APP_NAME not set")
	}

	licenseKey := os.Getenv("NEW_RELIC_LICENSE_KEY")
	if licenseKey == "" {
		return nil, errors.New("NEW_RELIC_LICENSE_KEY not set")
	}

	agentConfig := newrelic.NewConfig(appName, licenseKey)
	agentConfig.Logger = nrlogrus.StandardLogger()

	newrelicApp, startErr := newrelic.NewApplication(agentConfig)
	if nil != startErr {
		return nil, startErr
	}

	return newrelicApp, nil
}
