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

func AddSign(c buffalo.Context) (*models.Sign, interface{}) {
	db, err := ConnectDB(c)
	if err != nil {
		return nil, err
	}
	data := DynamicPostForm(c)
	sign := models.Sign{}
	if sign.CheckParamPostForm(data) {
		err, _ := GetSignByName(c)
		if err != nil {
			return nil, models.Error{500, "ชื่อนี้เคยสร้างไปแล้ว"}
		}
		file, err := UploadImg(c, data)
		if err != nil {
			return nil, err
		}
		sign.CreateSignModel(data, file.(string))
		db.Create(&sign)
		return &sign, nil
	}
	return nil, models.Error{400, "กรอกข้อมูลไม่ครบ"}
}

func UploadImg(c buffalo.Context, sign *models.Sign) (interface{}, interface{}) {
	f, err := c.File("file")
	tempFile, err := ioutil.TempFile(os.TempDir(), sign.Name+`-*.jpg`)
	if err != nil {
		log.Print(err)
		return nil, models.Error{500, "can't add sign"}
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(f)
	if err != nil {
		log.Print(err)
		return nil, models.Error{500, "can't add sign"}
	}

	tempFile.Write(fileBytes)
	in, err := os.Open(tempFile.Name())
	if err != nil {
		log.Print(err)
		return nil, models.Error{500, "can't add sign"}
	}
	defer in.Close()
	_, file := filepath.Split(tempFile.Name())
	out, err := os.Create(`D:\fe_booking_sign\public\img\` + file)
	if err != nil {
		log.Print(err)
		return nil, models.Error{500, "can't add sign"}
	}
	if _, err = io.Copy(out, in); err != nil {
		log.Print(err)
		return nil, models.Error{500, "can't add sign"}
	}

	return file, nil

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
	id, _ := strconv.Atoi(data["id"].(string))
	_ = db.Find(&sign, id)
	_ = db.Destroy(&sign)
	return &sign, nil
}
