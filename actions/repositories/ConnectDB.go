package repositories

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

func ConnectDB(c buffalo.Context) interface{} {
	db, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return map[string]interface{}{"error": "can't connect DB"}
	}
	return db
}
