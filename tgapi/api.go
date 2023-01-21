package tgapi

import "github.com/lowl11/lazytg/lazybot"

func NewBot(token string) (*lazybot.Bot, error) {
	return lazybot.Create(token)
}
