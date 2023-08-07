package lazytg

import "github.com/lowl11/lazytg/internal/lazybot"

func NewBot(token string) (Bot, error) {
	return lazybot.New(token)
}
