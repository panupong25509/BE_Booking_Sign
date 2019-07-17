package repositories

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/JewlyTwin/be_booking_sign/models"
	"github.com/gobuffalo/buffalo"
)

func AddSign(c buffalo.Context) (interface{}, interface{}) {
	db, err := ConnectDB(c)
	if err != nil {
		return nil, err
	}
	data := DynamicPostForm(c)
	sign := models.Sign{}
	if !sign.CheckParamPostForm(data) {
		return nil, models.Error{400, "กรอกข้อมูลไม่ครบ"}
	}
	samename, err := GetSignByName(c)
	if err != nil {
		return nil, err
	}
	if samename != nil {
		return nil, models.Error{400, "ชื่อป้ายนี้มีอยู่ในระบบแล้ว"}
	}
	file, err := UploadImg(c, data["signname"].(string))
	if err != nil {
		return nil, err
	}
	sign.CreateSignModel(data, file.(string))
	db.Create(&sign)
	return resSuccess(nil), nil
}

func UploadImg(c buffalo.Context, sign string) (interface{}, interface{}) {
	f, err := c.File("file")
	tempFile, err := ioutil.TempFile(os.TempDir(), sign+`-*.jpg`)
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(f)
	tempFile.Write(fileBytes)
	in, err := os.Open(tempFile.Name())
	defer in.Close()
	_, file := filepath.Split(tempFile.Name())
	out, err := os.Create(`D:\fe_booking_sign\public\img\` + file)
	if _, err = io.Copy(out, in); err != nil {
		log.Print(err)
		return nil, models.Error{500, "can't add sign"}
	}
	return file, nil

}

func GetAllSign(c buffalo.Context) (interface{}, interface{}) {
	db, err := ConnectDB(c)
	if err != nil {
		return nil, err
	}
	allSign := models.AllSign{}
	err = db.All(&allSign.Signs)
	if err != nil {
		return nil, nil
	}
	if len(allSign.Signs) == 0 {
		return nil, models.Error{400, "Not have sign"}
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
	log.Print(data["signname"].(string))
	_ = db.Where("sign_name in (?)", data["signname"].(string)).All(&sign)
	if len(sign) != 0 {
		return &sign[0], nil
	}
	return nil, nil
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

func DeleteSign(c buffalo.Context) (interface{}, interface{}) {
	db, err := ConnectDB(c)
	if err != nil {
		return nil, err
	}
	data := DynamicPostForm(c)
	sign := models.Sign{}
	id, _ := strconv.Atoi(data["id"].(string))
	err = db.Find(&sign, id)
	if err != nil {
		return nil, models.Error{500, "ไม่มี id นี้ใน Database"}
	}
	os.Remove(`D:\fe_booking_sign\public\img\` + sign.Picture)
	_ = db.Destroy(&sign)
	return resSuccess(nil), nil
}

func UpdateSign(c buffalo.Context) (interface{}, interface{}) {
	db, err := ConnectDB(c)
	if err != nil {
		return nil, err
	}
	data := DynamicPostForm(c)
	sign := models.Sign{}
	if !sign.CheckParamPostForm(data) {
		return nil, models.Error{400, "กรอกข้อมูลไม่ครบ"}
	}
	file, err := UploadImg(c, data["signname"].(string))
	if err != nil {
		return nil, err
	}
	sign.CreateSignModel(data, file.(string))
	oldSign, err := GetSignById(c, sign.ID)
	if err != nil {
		return nil, err
	}
	oldPicture := oldSign.(models.Sign).Picture
	os.Remove(`D:\fe_booking_sign\public\img\` + oldPicture)
	db.Update(&sign)
	return resSuccess(nil), nil
}
