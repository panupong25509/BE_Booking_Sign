package handlers

import (
	"github.com/JewlyTwin/be_booking_sign/actions/repositories"
	"github.com/JewlyTwin/be_booking_sign/models"
	"github.com/gobuffalo/buffalo"
)

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

// func DeleteBooking(c buffalo.Context) error {
// 	data := DynamicPostForm(c)

// 	deleteBooking, err := repositories.DeleteBooking(c, data)
// 	if err != nil {
// 		status := err.(models.Error)
// 		return c.Render(status.Code, r.JSON(status))
// 	}
// 	return c.Render(200, r.JSON(deleteBooking))
// }

func GetBookingDayBySign(c buffalo.Context) error {
	days, err := repositories.GetBookingDaysBySign(c)
	if err != nil {
		status := err.(models.Error)
		return c.Render(status.Code, r.JSON(status))
	}
	return c.Render(200, r.JSON(days))
}
func ApproveBooking(c buffalo.Context) error {
	data := DynamicPostForm(c)
	message, err := repositories.ApproveBooking(c, data)
	if err != nil {
		status := err.(models.Error)
		return c.Render(status.Code, r.JSON(status))
	}
	return c.Render(200, r.JSON(message))
}

func RejectBooking(c buffalo.Context) error {
	data := DynamicPostForm(c)
	message, err := repositories.RejectBooking(c, data)
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
