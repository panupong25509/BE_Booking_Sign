package handlers

import (
	// "github.com/JewlyTwin/practice/actions/repositories"
	"github.com/gobuffalo/buffalo"
)

func Hello(c buffalo.Context) error {
	Json := map[string]interface{}{"error": "can't connect DB"}
	return c.Render(200, r.JSON(Json))
}
