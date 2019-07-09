package grifts

import (
	"github.com/JewlyTwin/be_booking_sign/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
