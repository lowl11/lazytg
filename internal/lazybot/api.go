package lazybot

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/lowl11/lazytg/internal/message"
	"github.com/lowl11/lazytg/pkg/interfaces"
	"log"
)

func (bot *Bot) ProductionMode() interfaces.Bot {
	bot.connection.Debug = false
	return bot
}

func (bot *Bot) ThreadSafe() interfaces.Bot {
	bot.threadSafe = true
	return bot
}

func (bot *Bot) SetChatID(chatID int) interfaces.Bot {
	bot.chatID = int64(chatID)
	return bot
}

func (bot *Bot) Send(message string) error {
	if bot.chatID == 0 {
		return errors.New("chat ID is empty")
	}

	bot.lock()
	defer bot.unlock()

	return bot.sendMessage(message, bot.chatID)
}

func (bot *Bot) SendChat(message string, chatID int) error {
	bot.lock()
	defer bot.unlock()
	return bot.sendMessage(message, int64(chatID))
}

func (bot *Bot) RunAnswer(getMessageFunc func(ctx message.IContext) string, timeoutInSeconds int) interfaces.Bot {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = timeoutInSeconds

	updates := bot.connection.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message != nil {
			gotMessage := getMessageFunc(message.NewContext(
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

	return bot
}

func (bot *Bot) RunAnswerAsync(getMessageFunc func(ctx message.IContext) string, timeoutInSeconds int) interfaces.Bot {
	go bot.RunAnswer(getMessageFunc, timeoutInSeconds)
	return bot
}
