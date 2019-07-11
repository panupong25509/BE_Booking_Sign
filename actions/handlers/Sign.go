package handlers

import (
	"github.com/JewlyTwin/be_booking_sign/actions/repositories"
	"github.com/gobuffalo/buffalo"
)

func AddSign(c buffalo.Context) error {
	newSign, err := repositories.AddSign(c)
	return c.Render(200, r.JSON(newSign))
}
func GetAllSign(c buffalo.Context) error {
	allSign, err := repositories.GetAllSign(c)
	return c.Render(200, r.JSON(allSign))
}
