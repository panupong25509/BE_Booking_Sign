package handlers

import (
	// "github.com/JewlyTwin/practice/actions/repositories"
	"github.com/JewlyTwin/be_booking_sign/models"
	"github.com/gobuffalo/buffalo"
)

func Hello(c buffalo.Context) error {
	Json := map[string]interface{}{"error": "can't connect DB"}
	return c.Render(200, r.JSON(Json))
}

func HelloPost(c buffalo.Context) error {
	test := models.Sign{Name: "Sa1", Location: "หน้ามอ"}
	// Json := map[string]interface{}{"error": "test post"}
	return c.Render(200, r.JSON(test))
}
