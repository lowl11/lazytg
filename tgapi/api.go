package tgapi

import (
	"github.com/lowl11/lazytg/internal/lazybot"
)

func NewBot(token string) (*lazybot.Bot, error) {
	return lazybot.New(token)
}
