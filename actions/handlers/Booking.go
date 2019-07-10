package handlers

import (
	"github.com/JewlyTwin/be_booking_sign/actions/repositories"
	"github.com/gobuffalo/buffalo"
)

func AddBooking(c buffalo.Context) error {
	newBooking := repositories.AddBooking(c)
	return c.Render(200, r.JSON(newBooking))
}

func GetAllBooking(c buffalo.Context) error {
	allBooking := repositories.GetAllBooking(c)
	return c.Render(200, r.JSON(allBooking))
}
