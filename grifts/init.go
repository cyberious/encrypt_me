package grifts

import (
	"github.com/cyberious/encyrpt_me/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
