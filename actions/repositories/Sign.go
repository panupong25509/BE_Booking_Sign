package repositories

import (
	"github.com/gobuffalo/buffalo"

	"github.com/JewlyTwin/be_booking_sign/models"
)

func AddSign(c buffalo.Context) (*models.Sign, interface{}) {
	db, err := ConnectDB(c)
	if err != nil {
		return nil, err
	}
	data := DynamicPostForm(c)
	newSign := models.Sign{Name: data["name"].(string), Location: data["location"].(string)}
	db.Create(&newSign)
	return &newSign, nil
}

func GetAllSign(c buffalo.Context) (interface{}, interface{}) {
	db, errr := ConnectDB(c)
	if errr != nil {
		return nil, errr
	}
	allSign := []models.Sign{}
	err := db.Eager().All(&allSign)
	if err != nil {
		return nil, nil
	}
	return &allSign, nil
}

func GetSignByName(c buffalo.Context) (interface{}, interface{}) {
	db, err := ConnectDB(c)
	if err != nil {
		return nil, err
	}
	data := DynamicPostForm(c)
	sign := []models.Sign{}
	_ = db.Where("sign_name in (?)", data["signname"].(string)).All(&sign)
	return &sign[0], nil
}

func GetSignById(c buffalo.Context, id int) (interface{}, interface{}) {
	db, err := ConnectDB(c)
	if err != nil {
		return nil, err
	}
	sign := models.Sign{}
	_ = db.Find(&sign, id)
	return sign, nil
}
