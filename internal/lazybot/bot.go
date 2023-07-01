package lazybot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	chatID int64

	connection *tgbotapi.BotAPI

	username string
	name     string
}

func New(token string) (*Bot, error) {
	connection, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	connection.Debug = true
	username := connection.Self.UserName
	name := connection.Self.FirstName + " " + connection.Self.LastName

	return &Bot{
		connection: connection,

		username: username,
		name:     name,
	}, err
}
