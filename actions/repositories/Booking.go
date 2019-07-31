package repositories

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/JewlyTwin/be_booking_sign/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

func AddBooking(c buffalo.Context) (interface{}, interface{}) {
	db, err := ConnectDB(c)
	if err != nil {
		return nil, err
	}
	data := DynamicPostForm(c)
	signID, _ := strconv.Atoi(data["sign_id"].(string))
	sign, err := GetSignById(c, signID)
	if err != nil {
		return nil, err
	}
	code := GenCodeBooking(data, sign.(models.Sign))
	newBooking := models.Booking{}
	if !newBooking.CreateModel(data, code) {
		return nil, models.Error{400, "กรอกข้อมูลไม่ครบ"}
	}
	validate, err := ValidateBookingTime(newBooking, db, sign.(models.Sign))
	if err != nil {
		return nil, err
	}
	if !validate {
		return nil, models.Error{400, "วันที่เช่าไม่ว่าง"}
	}
	err = db.Create(&newBooking)
	if err != nil {
		return nil, models.Error{500, "Can't Create to Database"}
	}
	return newBooking, nil
}

func GenCodeBooking(data map[string]interface{}, sign models.Sign) string {
	code := sign.Name + "CODE" + data["first_date"].(string) + data["last_date"].(string)
	return code
}

func ValidateBookingTime(newBooking models.Booking, db *pop.Connection, sign models.Sign) (bool, interface{}) {
	bookings := models.Bookings{}
	_ = db.Q().Where("last_date >= (?) and first_date <= (?) and sign_id = (?)", newBooking.FirstDate, newBooking.LastDate, newBooking.SignID).All(&bookings)
	if len(bookings) != 0 {
		return false, models.Error{500, "วันที่เช่าไม่ว่าง"}
	}
	if CheckDate(newBooking.FirstDate, newBooking.LastDate) > sign.Limitdate {
		return false, models.Error{500, "กรุณาจองในระยะเวลา " + strconv.Itoa(sign.Limitdate) + " วัน"}
	}
	if CheckDate(time.Now(), newBooking.FirstDate) < sign.Beforebooking {
		return false, models.Error{500, "กรุณาจองก่อน " + strconv.Itoa(sign.Beforebooking) + " วัน"}
	}
	return true, nil
}

func GetBookingByUser(c buffalo.Context) (interface{}, interface{}) {
	jwtReq := c.Request().Header.Get("Authorization")
	log.Print(jwtReq)
	jwtStrings := strings.Split(jwtReq, "Bearer ")
	log.Print(jwtStrings[1])
	token, _ := jwt.Parse(jwtStrings[1], func(token *jwt.Token) (interface{}, error) {
		return []byte("bookingsign"), nil
	})
	tokens := token.Claims.(jwt.MapClaims)
	log.Print(tokens["UserID"])
	db, err := ConnectDB(c)
	if err != nil {
		return nil, err
	}
	allBooking := models.Allbooking{}
	err = db.Q().Where("applicant_id = (?)", tokens["UserID"]).All(&allBooking.Booking)
	if err != nil {
		return nil, models.Error{500, "Can't Select data form Database"}
	}
	bookings := models.Allbooking{}
	for _, value := range allBooking.Booking {
		user, err := GetUserByIduuid(c, value.ApplicantID)
		log.Print(user)
		if err != nil {
			return nil, err
		}
		value.Applicant = user.(models.User)
		sign, err := GetSignById(c, value.SignID)
		if err != nil {
			return nil, err
		}
		value.Sign = sign.(models.Sign)
		bookings.Booking = append(bookings.Booking, value)
	}
	log.Print(bookings)
	return &bookings, nil
}

func CheckDate(D1 time.Time, D2 time.Time) int {
	diff := D2.Sub(D1)
	allDay := int(diff.Hours()/24) + 1 //first-last
	day := D1
	sunday := 0
	for day.Before(D2) {
		if int(day.Weekday()) == 0 {
			sunday = sunday + 1
			day = day.AddDate(0, 0, 7)
		} else {
			day = day.AddDate(0, 0, 1)
		}
	}
	weekday := sunday * 2 // weekday in firstdate - lastdate
	return allDay - weekday
}

func DeleteBooking(c buffalo.Context) (interface{}, interface{}) {
	db, err := ConnectDB(c)
	if err != nil {
		return nil, models.Error{500, "Can't connect Database"}
	}
	data := DynamicPostForm(c)
	booking := models.Booking{}
	id, _ := strconv.Atoi(data["id"].(string))
	err = db.Find(&booking, id)
	if err != nil {
		return nil, models.Error{500, "Data มีปัญหาไม่สามารถยกเลิกได้"}
	}
	_ = db.Destroy(&booking)
	return models.Error{200, "ยกเลิกสำเร็จ"}, nil
}

func GetBookingDaysBySign(c buffalo.Context) (interface{}, interface{}) {
	db, err := ConnectDB(c)
	if err != nil {
		return nil, models.Error{500, "Can't connect Database"}
	}
	bookings := models.Bookings{}
	bookingdate := time.Now().Format("2006-01-02")
	signid, _ := strconv.Atoi(c.Param("id"))
	err = db.Q().Where("( last_date >= (?) or first_date >= (?) ) and sign_id = (?)", bookingdate, bookingdate, signid).All(&bookings)
	if err != nil {
		return nil, models.Error{400, "DB"}
	}
	days := models.BookingDays{}
	for _, value := range bookings {
		days = append(days, models.BookingDay{value.FirstDate, value.LastDate})
	}
	return days, nil
}
