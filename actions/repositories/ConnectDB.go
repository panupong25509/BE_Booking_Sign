package repositories

import (
	"github.com/JewlyTwin/be_booking_sign/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

func ConnectDB(c buffalo.Context) (*pop.Connection, interface{}) {
	db, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return nil, models.Error{500, "can't connect db"}
	}
	return db, nil
}
