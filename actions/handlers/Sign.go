package handlers

import (
	"github.com/JewlyTwin/be_booking_sign/actions/repositories"
	"github.com/gobuffalo/buffalo"
)

func AddSign(c buffalo.Context) error {
	newSign := repositories.AddSign(c)
	// Json := map[string]interface{}{"error": "test post"}
	return c.Render(200, r.JSON(newSign))
}
func GetAllSign(c buffalo.Context) error {
	allSign := repositories.GetAllSign(c)
	// Json := map[string]interface{}{"error": "test post"}
	return c.Render(200, r.JSON(allSign))
}
func GetSignByName(c buffalo.Context) error {
	sign := repositories.GetSignByName(c)
	// Json := map[string]interface{}{"error": "test post"}
	return c.Render(200, r.JSON(sign))
}
