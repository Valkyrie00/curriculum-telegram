package types

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"time"
)

// Message
type Message struct {
	ChatID   int64
	MsgType  string
	Duration time.Duration
	Content  string
	Keyboard *tgbotapi.ReplyKeyboardMarkup
}

