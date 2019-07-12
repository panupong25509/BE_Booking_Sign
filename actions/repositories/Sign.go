package repositories

import (
	"github.com/JewlyTwin/be_booking_sign/models"
	"github.com/gobuffalo/buffalo"
	"strconv"
)

func AddSign(c buffalo.Context) (*models.Sign, interface{}) {
	db, err := ConnectDB(c)
	if err != nil {
		return nil, err
	}
	data := DynamicPostForm(c)
	sign := models.Sign{}
	if sign.CheckParamPostForm(data) {
		sign.CreateBookingModel(data)
		samename := db.Create(&sign)
		if samename != nil {
			return nil, models.Error{500, "ชื่อนี้เคยสร้างไปแล้ว"}
		}
		return &sign, nil
	}
	return nil, models.Error{400, "กรอกข้อมูลไม่ครบ"}
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

func DeleteSignByID(c buffalo.Context) (interface{}, interface{}) {
	db, err := ConnectDB(c)
	if err != nil {
		return nil, err
	}
	data := DynamicPostForm(c)
	sign := models.Sign{}
	id ,_ := strconv.Atoi(data["id"].(string))
	_ = db.Find(&sign, id)
	_ = db.Destroy(&sign)
	return &sign, nil
}
