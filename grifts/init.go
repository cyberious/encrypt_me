package grifts

import (
	"github.com/cyberious/encrypt_me/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
