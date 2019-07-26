package repositories

import (
	"log"

	"golang.org/x/crypto/bcrypt"

	"github.com/JewlyTwin/be_booking_sign/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gofrs/uuid"
)

func Register(c buffalo.Context) (interface{}, interface{}) {
	db, err := ConnectDB(c)
	if err != nil {
		return nil, err
	}
	data := DynamicPostForm(c)
	user := models.User{}
	if !user.CheckParams(data) {
		return nil, models.Error{400, "กรอกข้อมูลไม่ครบ"}
	}
	_, err = GetUserByUsername(c)
	if err == nil {
		log.Print(err)
		return nil, models.Error{500, "Username นี้มีผู้ใช้แล้ว"}
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(data["password"].(string)), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	_ = user.CreateModel(data, string(hash))
	err = db.Create(&user)
	if err != nil {
		return nil, err
	}
	return resSuccess(nil), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
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

func GetUserByIduuid(c buffalo.Context, id uuid.UUID) (interface{}, interface{}) {
	db, err := ConnectDB(c)
	if err != nil {
		return nil, err
	}
	user := models.User{}
	err = db.Find(&user, id)
	// log.Print(user)
	if err != nil {
		return nil, models.Error{400, "ไม่มีผู้ใช้นี้ใน"}
	}
	return user, nil
}

func GetUserByUsername(c buffalo.Context) (interface{}, interface{}) {
	db, err := ConnectDB(c)
	if err != nil {
		return nil, err
	}
	data := DynamicPostForm(c)
	if data["username"] == nil {
		return nil, models.Error{400, "ไม่มี username"}
	}
	username := data["username"].(string)
	user := models.Users{}
	_ = db.Q().Where("username >= (?)", username).All(&user)
	if len(user) == 0 {
		return nil, models.Error{400, "ไม่มี username"}
	}
	return user[0], nil
}

func CheckUsernamePassword(c buffalo.Context) (interface{}, interface{}) {
	db, err := ConnectDB(c)
	if err != nil {
		return nil, err
	}
	data := DynamicPostForm(c)
	username := data["username"].(string)
	password := data["password"].(string)
	user := models.Users{}
	if username == "" {
		return nil, models.Error{400, "ไม่มี username"}
	}
	if password == "" {
		return nil, models.Error{400, "ไม่มี password"}
	}
	_ = db.Q().Where("username = (?)", username).All(&user)
	if CheckPasswordHash(password, user[0].Password) {
		return &user[0], nil
	}
	return nil, models.Error{400, "ผิดดดดดด"}
}
