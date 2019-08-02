package handlers

import (
	"github.com/JewlyTwin/be_booking_sign/actions/repositories"
	"github.com/JewlyTwin/be_booking_sign/models"
	"github.com/gobuffalo/buffalo"
)

func Login(c buffalo.Context) error {
	data := DynamicPostForm(c)
	success, err := repositories.Login(c, data)
	if err != nil {
		status := err.(models.Error)
		return c.Render(status.Code, r.JSON(status))
	}
	return c.Render(200, r.JSON(success))
}
func Register(c buffalo.Context) error {
	data := DynamicPostForm(c)
	success, err := repositories.Register(c, data)
	if err != nil {
		status := err.(models.Error)
		return c.Render(status.Code, r.JSON(status))
	}
	return c.Render(200, r.JSON(success))
}
func GetUserByUsername(c buffalo.Context) error {
	data := DynamicPostForm(c)
	success, err := repositories.GetUserByUsername(c, data)
	if err != nil {
		status := err.(models.Error)
		return c.Render(status.Code, r.JSON(status))
	}
	return c.Render(200, r.JSON(success))
}

func GetUserById(c buffalo.Context) error {
	success, err := repositories.GetUserById(c)
	if err != nil {
		status := err.(models.Error)
		return c.Render(status.Code, r.JSON(status))
	}
	return c.Render(200, r.JSON(success))
}
