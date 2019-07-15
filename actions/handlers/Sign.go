package handlers

import (
	"github.com/JewlyTwin/be_booking_sign/actions/repositories"
	"github.com/JewlyTwin/be_booking_sign/models"
	"github.com/gobuffalo/buffalo"
)

func AddSign(c buffalo.Context) error {
	success, err := repositories.AddSign(c)
	if err != nil {
		status := err.(models.Error)
		return c.Render(status.Code, r.JSON(status.Message))
	}
	return c.Render(200, r.JSON(success))
}

func GetAllSign(c buffalo.Context) error {
	allSign, err := repositories.GetAllSign(c)
	if err != nil {
		status := err.(models.Error)
		return c.Render(status.Code, r.JSON(status.Message))
	}
	return c.Render(200, r.JSON(allSign))
}

func DeleteSign(c buffalo.Context) error {
	destroy, err := repositories.DeleteSignByID(c)
	if err != nil {
		status := err.(models.Error)
		return c.Render(status.Code, r.JSON(status.Message))
	}
	return c.Render(200, r.JSON(destroy))
}
