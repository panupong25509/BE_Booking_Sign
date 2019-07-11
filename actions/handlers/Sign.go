package handlers

import (
	"github.com/JewlyTwin/be_booking_sign/actions/repositories"
	"github.com/gobuffalo/buffalo"
)

func AddSign(c buffalo.Context) error {
	newSign := repositories.AddSign(c)
	return c.Render(200, r.JSON(newSign))
}
func GetAllSign(c buffalo.Context) error {
	allSign := repositories.GetAllSign(c)
	return c.Render(200, r.JSON(allSign))
}
