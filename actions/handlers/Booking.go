package handlers

import (
	"github.com/JewlyTwin/be_booking_sign/actions/repositories"
	"github.com/JewlyTwin/be_booking_sign/models"
	"github.com/gobuffalo/buffalo"
)

func GetBookingByUser(c buffalo.Context) error {
	allBooking, err := repositories.GetBookingByUser(c)
	if err != nil {
		status := err.(models.Error)
		return c.Render(status.Code, r.JSON(status))
	}
	return c.Render(200, r.JSON(allBooking))
}
func AddBooking(c buffalo.Context) error {
	data := DynamicPostForm(c)
	newBooking, err := repositories.AddBooking(c, data)
	if err != nil {
		status := err.(models.Error)
		return c.Render(status.Code, r.JSON(status))
	}
	booking := newBooking.(models.Booking)
	return c.Render(200, r.JSON(booking.ReturnJsonID()))
}

func DeleteBooking(c buffalo.Context) error {
	deleteBooking, err := repositories.DeleteBooking(c)
	if err != nil {
		status := err.(models.Error)
		return c.Render(status.Code, r.JSON(status))
	}
	return c.Render(200, r.JSON(deleteBooking))
}

func GetBookingDayBySign(c buffalo.Context) error {
	days, err := repositories.GetBookingDaysBySign(c)
	if err != nil {
		status := err.(models.Error)
		return c.Render(status.Code, r.JSON(status))
	}
	return c.Render(200, r.JSON(days))
}

func RejectBooking(c buffalo.Context) error {
	message, err := repositories.RejectBooking(c)
	if err != nil {
		status := err.(models.Error)
		return c.Render(status.Code, r.JSON(status))
	}
	return c.Render(200, r.JSON(message))
}

func GetPaginateAdmin(c buffalo.Context) error {
	page := c.Param("page")
	booking, err := repositories.GetPaginateAdmin(page, c)
	if err != nil {
		status := err.(models.Error)
		return c.Render(status.Code, r.JSON(status))
	}
	return c.Render(200, r.JSON(booking))
}

func GetPaginateUser(c buffalo.Context) error {
	page := c.Param("page")
	allBooking, err := repositories.GetPaginateUser(page, c)
	if err != nil {
		status := err.(models.Error)
		return c.Render(status.Code, r.JSON(status))
	}
	return c.Render(200, r.JSON(allBooking))
}
