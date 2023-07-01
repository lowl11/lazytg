package lazybot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (bot *Bot) sendMessage(message string, chatID int64) error {
	messageObject := tgbotapi.NewMessage(chatID, message)
	messageObject.ParseMode = tgbotapi.ModeMarkdownV2

	if _, err := bot.connection.Send(messageObject); err != nil {
		return err
	}

	return nil
}
