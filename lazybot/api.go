package lazybot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/lowl11/lazytg/message"
	"log"
)

func (bot *Bot) ProductionMode() {
	bot.connection.Debug = false
}

func (bot *Bot) SetChatID(chatID int) {
	bot.chatID = int64(chatID)
}

func (bot *Bot) Send(message string) error {
	return bot.sendMessage(message, bot.chatID)
}

func (bot *Bot) SendChat(message string, chatID int) error {
	return bot.sendMessage(message, int64(chatID))
}

func (bot *Bot) RunAnswer(getMessageFunc func(ctx message.IContext) string, timeoutInSeconds int) {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = timeoutInSeconds

	updates := bot.connection.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message != nil {
			gotMessage := getMessageFunc(message.CreateContext(
				&message.Message{
					Text:   update.Message.Text,
					ChatID: update.Message.Chat.ID,
				},
				&message.Author{
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
