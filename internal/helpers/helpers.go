package helpers

import (
	"github.com/Valkyrie00/curriculum-telegram/internal/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"time"
)

// ConsumeChainMessage
func ConsumeChainMessage(bot *tgbotapi.BotAPI, structure types.Message) {
	switch structure.MsgType {
	case "Message":
		var response = tgbotapi.NewMessage(structure.ChatID, structure.Content)
		response.ParseMode = "Markdown"
		if structure.Keyboard != nil {
			response.ReplyMarkup = structure.Keyboard
		}

		SendMsg(bot, response)

	case "NewDocumentUpload":
		var response = tgbotapi.NewDocumentUpload(structure.ChatID, structure.Content)
		if structure.Keyboard != nil {
			response.ReplyMarkup = structure.Keyboard
		}

		SendMsg(bot, response)

	case "NewPhotoUpload":
		var response = tgbotapi.NewPhotoUpload(structure.ChatID, structure.Content)
		if structure.Keyboard != nil {
			response.ReplyMarkup = structure.Keyboard
		}

		SendMsg(bot, response)
	}

	time.Sleep(structure.Duration * time.Second)

}

// SendMsg - Send telegram message
func SendMsg(bot *tgbotapi.BotAPI, response tgbotapi.Chattable) {
	if _, err := bot.Send(response); err != nil {
		log.Panicln(err)
	}
}