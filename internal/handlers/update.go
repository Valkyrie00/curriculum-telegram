package handlers

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	bot *tgbotapi.BotAPI
)

// UpdatesHandler
func UpdatesHandler() {
	var err error

	debug := os.Getenv("DEBUG")
	telegramApikey := os.Getenv("TELEGRAM_APIKEY")

	if bot, err = tgbotapi.NewBotAPI(telegramApikey); err != nil {
		log.Fatalf("error bot api: %s", err)
	}

	// Check debug status
	if debug == "true" {
		bot.Debug = true
	}

	fmt.Printf("Bot connected %s", bot.Self.UserName)

	// Get telegram updates
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	var updates tgbotapi.UpdatesChannel
	if updates, err = bot.GetUpdatesChan(u); err  != nil {
		log.Fatalf("error gettin updates chan: %s", err)
	}

	// Each upodates
	for update := range updates {
		if update.Message == nil {
			continue
		}

		go commandsHandler(update)
	}
}
