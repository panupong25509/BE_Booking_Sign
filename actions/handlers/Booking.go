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

// func DeleteBooking(c buffalo.Context) error {
// 	deleteBooking, err := repositories.DeleteBooking(c)
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

//admin
func GetBookingForAdmin(c buffalo.Context) error {
	allBooking, err := repositories.GetBookingForAdmin(c)
	if err != nil {
		status := err.(models.Error)
		return c.Render(status.Code, r.JSON(status))
	}
	return c.Render(200, r.JSON(allBooking))
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
