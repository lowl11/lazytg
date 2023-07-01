package lazybot

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	message2 "github.com/lowl11/lazytg/internal/message"
	"log"
)

func (bot *Bot) ProductionMode() {
	bot.connection.Debug = false
}

func (bot *Bot) SetChatID(chatID int) {
	bot.chatID = int64(chatID)
}

func (bot *Bot) Send(message string) error {
	if bot.chatID == 0 {
		return errors.New("chat ID is empty")
	}

	return bot.sendMessage(message, bot.chatID)
}

func (bot *Bot) SendChat(message string, chatID int) error {
	return bot.sendMessage(message, int64(chatID))
}

func (bot *Bot) RunAnswer(getMessageFunc func(ctx message2.IContext) string, timeoutInSeconds int) {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = timeoutInSeconds

	updates := bot.connection.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message != nil {
			gotMessage := getMessageFunc(message2.NewContext(
				&message2.Message{
					Text:   update.Message.Text,
					ChatID: update.Message.Chat.ID,
				},
				&message2.Author{
					Username:  update.Message.From.UserName,
					LastName:  update.Message.From.LastName,
					FirstName: update.Message.From.FirstName,
				},
			))

			chatID := update.Message.Chat.ID
			answer := tgbotapi.NewMessage(chatID, gotMessage)

			if _, err := bot.connection.Send(answer); err != nil {
				log.Println("Send message error: ", err)
			}
		}
	}
}

func (bot *Bot) RunAnswerAsync(getMessageFunc func(ctx message2.IContext) string, timeoutInSeconds int) {
	go bot.RunAnswer(getMessageFunc, timeoutInSeconds)
}
