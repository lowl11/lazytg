package message

type IContext interface {
	Message() *Message
	Author() *Author
}

type TelegramContext struct {
	message *Message
	author  *Author
}

func CreateContext(message *Message, author *Author) IContext {
	return &TelegramContext{
		message: message,
		author:  author,
	}
}

func (ctx *TelegramContext) Message() *Message {
	return ctx.message
}

func (ctx *TelegramContext) Author() *Author {
	return ctx.author
}
