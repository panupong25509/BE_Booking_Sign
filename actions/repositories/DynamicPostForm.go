package repositories

import (
	"github.com/gobuffalo/buffalo"
)

func DynamicPostForm(c buffalo.Context) map[string]interface{} {
	c.Request().ParseForm()
	param := c.Request().PostForm
	v := make(map[string]interface{})
	for key, value := range param {
		v[key] = value[0]
	}
	return v
}
