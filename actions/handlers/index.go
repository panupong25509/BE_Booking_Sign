package handlers

import (
	"github.com/gobuffalo/buffalo"
)

func Hello(c buffalo.Context) error {
	Json := map[string]interface{}{"status": "Connected"}
	return c.Render(200, r.JSON(Json))
}
