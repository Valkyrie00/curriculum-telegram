package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

var (
	bot *tgbotapi.BotAPI
)

func pingServer() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8011"
		Info("$PORT must be set")
	}

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello :)")
	})

	http.ListenAndServe(":"+port, nil)
}

func updatesHandler() {
	var telegramApikey string

	errDotEnv := godotenv.Load()
	if errDotEnv != nil {
		log.Fatal("Error loading .env file")
	}
	telegramApikey = os.Getenv("TELEGRAM_APIKEY")

	var err error
	bot, err = tgbotapi.NewBotAPI(telegramApikey)
	bot.Debug = true

	if err != nil {
		Panic(err)
	}

	logMessage := fmt.Sprintf("Bot connesso correttamente %s", bot.Self.UserName)
	Info(logMessage)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, chanErr := bot.GetUpdatesChan(u)
	if chanErr != nil {
		Panic(chanErr)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		commandsHandler(update)
	}
}

func commandsHandler(update tgbotapi.Update) {
	command, _, ok := breakCommand(update.Message.Text)
	if ok {
		switch command {
		case "Torna":
			TornaCommand(update)
		case "Start":
			StartCommand(update)
		case "/Start":
			StartCommand(update)
		case "Raccontami":
			StoryCommand(update)
		case "Percorso":
			JobsCommand(update)
		case "Tecnologie":
			TechCommand(update)
		case "Contatti":
			ContactsCommand(update)
		}
	}
}

// SendMsg - Send telegram message
func SendMsg(response tgbotapi.Chattable) {
	if _, err := bot.Send(response); err != nil {
		Panic(err)
	}
}

func breakCommand(message string) (string, []string, bool) {
	var command []string
	var arguments []string
	if message == "" {
		return "", arguments, false
	}

	command = strings.Split(message, " ")
	if len(command) >= 2 {
		arguments = strings.Split(command[1], ",")
	}

	return command[0], arguments, true
}

func main() {
	// PingServer
	Info("Start PingServer")
	go pingServer()

	// Bot
	Info("Start Bot")
	updatesHandler()
}
