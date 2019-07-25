package repositories

import (
	"github.com/JewlyTwin/be_booking_sign/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gofrs/uuid"
)

func AddUser(c buffalo.Context) (interface{}, interface{}) {
	db, err := ConnectDB(c)
	if err != nil {
		return nil, err
	}
	data := DynamicPostForm(c)
	user := models.User{}
	if !user.CreateModel(data) {
		return nil, models.Error{400, "กรอกข้อมูลไม่ครบ"}
	}
	err = db.Create(&user)
	if err != nil {
		return nil, err
	}
	return resSuccess(nil), nil
}

func GetUserById(c buffalo.Context) (interface{}, interface{}) {
	db, err := ConnectDB(c)
	if err != nil {
		return nil, err
	}
	data := DynamicPostForm(c)
	id, err := uuid.FromString(data["id"].(string))
	if err != nil {
		return nil, err
	}
	user := models.User{}
	err = db.Find(&user, id)
	if err != nil {
		return nil, models.Error{400, "ไม่มีผู้ใช้นี้ใน database"}
	}
	return user, nil
}
