package repositories

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"

	"github.com/JewlyTwin/be_booking_sign/models"
	// "log"
)

func AddSign(c buffalo.Context) interface{} {
	db := ConnectDB(c).(*pop.Connection)
	data := DynamicPostForm(c)
	newSign := models.Sign{Name: data["name"].(string), Location: data["location"].(string)}

	db.Create(&newSign)
	return &newSign
}
func GetAllSign(c buffalo.Context) interface{} {
	db := ConnectDB(c).(*pop.Connection)
	// data := DynamicPostForm(c)
	allSign := []models.Sign{}

	db.All(&allSign)
	return &allSign
}
func GetSignByName(c buffalo.Context) interface{} {
	db := ConnectDB(c).(*pop.Connection)
	data := DynamicPostForm(c)
	sign := []models.Sign{}

	_ = db.Where("sign_name in (?)", data["signname"].(string)).All(&sign)
	return &sign[0]
}
