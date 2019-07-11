package repositories

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"

	"github.com/JewlyTwin/be_booking_sign/models"
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
	allSign := []models.Sign{}

	err := db.Eager().All(&allSign)
	if err != nil {
		return &allSign
	}
	return nil
}

func GetSignByName(c buffalo.Context) interface{} {
	db := ConnectDB(c).(*pop.Connection)
	data := DynamicPostForm(c)
	sign := []models.Sign{}

	_ = db.Where("sign_name in (?)", data["signname"].(string)).All(&sign)
	return &sign[0]
}

func GetSignById(c buffalo.Context, id int) models.Sign {
	db := ConnectDB(c).(*pop.Connection)
	sign := models.Sign{}

	_ = db.Find(&sign, id)
	return sign
}
