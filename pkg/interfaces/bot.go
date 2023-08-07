package interfaces

import "github.com/lowl11/lazytg/internal/message"

type Bot interface {
	ProductionMode() Bot
	ThreadSafe() Bot
	SetChatID(chatID int) Bot
	Send(message string) error
	SendChat(message string, chatID int) error
	RunAnswer(getMessageFunc func(ctx message.IContext) string, timeoutInSeconds int) Bot
	RunAnswerAsync(getMessageFunc func(ctx message.IContext) string, timeoutInSeconds int) Bot
}
