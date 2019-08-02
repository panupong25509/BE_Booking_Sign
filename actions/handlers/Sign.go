package handlers

import (
	"github.com/JewlyTwin/be_booking_sign/actions/repositories"
	"github.com/JewlyTwin/be_booking_sign/models"
	"github.com/gobuffalo/buffalo"
)

func GetAllSign(c buffalo.Context) error {
	allSign, err := repositories.GetAllSign(c)
	if err != nil {
		status := err.(models.Error)
		return c.Render(status.Code, r.JSON(status))
	}
	return c.Render(200, r.JSON(allSign))
}
func AddSign(c buffalo.Context) error {
	data := DynamicPostForm(c)
	success, err := repositories.AddSign(c, data)
	if err != nil {
		status := err.(models.Error)
		return c.Render(status.Code, r.JSON(status))
	}
	return c.Render(200, r.JSON(success))
}

func DeleteSign(c buffalo.Context) error {
	data := DynamicPostForm(c)
	destroy, err := repositories.DeleteSign(c, data)
	if err != nil {
		status := err.(models.Error)
		return c.Render(status.Code, r.JSON(status))
	}
	return c.Render(200, r.JSON(destroy))
}

func UpdateSign(c buffalo.Context) error {
	data := DynamicPostForm(c)
	res, err := repositories.UpdateSign(c, data)
	if err != nil {
		status := err.(models.Error)
		return c.Render(status.Code, r.JSON(status))
	}
	return c.Render(200, r.JSON(res))
}

func GetSignById(c buffalo.Context) error {
	res, err := repositories.GetSignById(c, 0)
	if err != nil {
		status := err.(models.Error)
		return c.Render(status.Code, r.JSON(status))
	}
	return c.Render(200, r.JSON(res))
}
