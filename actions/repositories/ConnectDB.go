package repositories

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

func ConnectDB(c buffalo.Context) (*pop.Connection, interface{}) {
	db, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return nil, 500
	}
	return db, nil
}
