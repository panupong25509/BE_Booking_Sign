package handlers

import (
	"github.com/JewlyTwin/be_booking_sign/actions/repositories"
	"github.com/JewlyTwin/be_booking_sign/models"
	"github.com/gobuffalo/buffalo"
)

func AddBooking(c buffalo.Context) error {
	newBooking, err := repositories.AddBooking(c)
	if err != nil {
		status := err.(models.Error)
		return c.Render(status.Code, r.JSON(status.Message))
	}
	return c.Render(200, r.JSON(newBooking))
}

func GetAllBooking(c buffalo.Context) error {
	allBooking, err := repositories.GetAllBooking(c)
	if err != nil {
		status := err.(models.Error)
		return c.Render(status.Code, r.JSON(status.Message))
	}
	return c.Render(200, r.JSON(allBooking))
}
