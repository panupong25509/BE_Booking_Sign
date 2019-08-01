package repositories

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"gopkg.in/gomail.v2"

	"github.com/JewlyTwin/be_booking_sign/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

func AddBooking(c buffalo.Context, data map[string]interface{}) (interface{}, interface{}) {
	db, err := ConnectDB(c)
	if err != nil {
		return nil, err
	}
	signID, _ := strconv.Atoi(data["sign_id"].(string))
	sign, err := GetSignById(c, signID)
	if err != nil {
		return nil, err
	}
	code := GenCodeBooking(data, sign.(models.Sign))
	newBooking := models.Booking{}
	if !newBooking.CreateModel(data, code) {
		return nil, models.Error{400, "Please complete all fields"}
	}
	validate, err := ValidateBookingTime(newBooking, db, sign.(models.Sign))
	if err != nil {
		return nil, err
	}
	if !validate {
		return nil, models.Error{400, "Busy date"}
	}
	err = db.Create(&newBooking)
	if err != nil {
		return nil, models.Error{500, "Can't Create to Database"}
	}
	send2()
	return newBooking, nil
}

// func send(body string) {
// 	log.Print("_________________________________________________________________________________________________________-_________________________________________________________________________________________________________-_________________________________________________________________________________________________________-_________________________________________________________________________________________________________-_________________________________________________________________________________________________________-_________________________________________________________________________________________________________-_________________________________________________________________________________________________________-")
// 	emailsender := "l2jew123@gmail.com"
// 	pass := "JewlyTwin123"
// 	emailreceiver := "panupong.jkn@gmail.com"

// 	msg := "From: " + emailsender + "\n" +
// 		"To: " + emailreceiver + "\n" +
// 		"Subject: Hello there\n\n" +
// 		body

// 	emailAuth := smtp.PlainAuth("", emailsender, pass, "smtp.gmail.com")

// 	err := smtp.SendMail("smtp.gmail.com:587",
// 		emailAuth,
// 		emailsender, []string{emailreceiver}, []byte(msg))

// 	if err != nil {
// 		log.Printf("smtp error: %s", err)
// 		return
// 	}

// 	log.Print("sent, Success")
// }

func send2() {
	mail := gomail.NewMessage()
	mail.SetHeader("From", "l2jew123@gmail.com")
	mail.SetHeader("To", "panupong.jkn@gmail.com")
	mail.SetHeader("Subject", "SADDDDDDDDDDDDDD")
	mail.SetBody("test/plain", "How are you")
	dialer := gomail.NewPlainDialer("smtp.gmail.com", 587, "l2jew123@gmail.com", "JewlyTwin123")
	if err := dialer.DialAndSend(mail); err != nil {
		panic(err)
	}
	fmt.Print("Email Sent")
}

func GenCodeBooking(data map[string]interface{}, sign models.Sign) string {
	code := sign.Name + "CODE" + data["first_date"].(string) + data["last_date"].(string)
	return code
}

func ValidateBookingTime(newBooking models.Booking, db *pop.Connection, sign models.Sign) (bool, interface{}) {
	bookings := models.Bookings{}
	_ = db.Q().Where("last_date >= (?) and first_date <= (?) and sign_id = (?)", newBooking.FirstDate, newBooking.LastDate, newBooking.SignID).All(&bookings)
	if len(bookings) != 0 {
		return false, models.Error{500, "Busy date"}
	}
	if CheckDate(newBooking.FirstDate, newBooking.LastDate) > sign.Limitdate {
		return false, models.Error{500, "Please book within " + strconv.Itoa(sign.Limitdate) + " days"}
	}
	if CheckDate(time.Now(), newBooking.FirstDate) < sign.Beforebooking {
		return false, models.Error{500, "Please book before " + strconv.Itoa(sign.Beforebooking) + " days"}
	}
	return true, nil
}

func GetBookingByUser(c buffalo.Context) (interface{}, interface{}) {
	jwtReq := c.Request().Header.Get("Authorization")
	tokens, err := DecodeJWT(jwtReq, "bookingsign")
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
	return &bookings, nil
}
func GetBookingForAdmin(c buffalo.Context) (interface{}, interface{}) {
	jwtReq, err := GetJWT(c)
	if err != nil {
		return nil, err
	}
	tokens, err := DecodeJWT(jwtReq.(string), "bookingsign")
	if err != nil {
		return nil, err
	}

	db, err := ConnectDB(c)
	if err != nil {
		return nil, err
	}
	allBooking := models.Allbooking{}
	if tokens["Role"] != "admin" {
		return nil, models.Error{500, "You not Admin"}
	}
	err = db.Q().Where("status = (?)", "pending").All(&allBooking.Booking)
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
	return bookings, nil
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
	return models.Error{200, "Delete success"}, nil
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
