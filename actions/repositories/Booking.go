package repositories

import (
	"time"

	"github.com/JewlyTwin/be_booking_sign/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

func AddBooking(c buffalo.Context) (*models.Booking, interface{}) {
	db, errdb := ConnectDB(c)
	if errdb != nil {
		return nil, errdb
	}
	data := DynamicPostForm(c)
	newBooking := models.Booking{}
	if !newBooking.CheckParamPostForm(data) {
		return nil, models.Error{400, "param not complete"}
	}
	code := GenCodeBooking(data)
	signInterface, errsignInterface := GetSignByName(c)
	if errsignInterface != nil {
		return nil, errsignInterface
	}
	sign := signInterface.(*models.Sign)
	newBooking.CreateBookingModel(data, code, *sign)
	if !ValidateBookingTime(&newBooking, db, *sign) {
		return nil, models.Error{400, "The date of booking is not available"}
	}
	errdbcreate := db.Create(&newBooking)
	if errdbcreate != nil {
		return nil, models.Error{500, "Can't Create to Database"}
	}
	return &newBooking, nil
}

func GenCodeBooking(data map[string]interface{}) string {
	code := data["signname"].(string) + "CODE" + data["firstdate"].(string) + data["lastdate"].(string)
	return code
}

func ValidateBookingTime(newBooking *models.Booking, db *pop.Connection, sign models.Sign) bool {
	bookings := models.Bookings{}
	_ = db.Q().Where("last_date >= (?) and first_date <= (?) and sign_id = (?)", newBooking.FirstDate, newBooking.LastDate, newBooking.SignID).All(&bookings)
	if len(bookings) != 0 {
		return false
	}
	if CheckDate(newBooking.FirstDate, newBooking.LastDate) > sign.Limitdate {
		return false
	}
	if CheckDate(time.Now(), newBooking.FirstDate) < sign.Beforebooking {
		return false
	}
	return true
}

func GetAllBooking(c buffalo.Context) (*models.Bookings, interface{}) {
	db, err := ConnectDB(c)
	if err != nil {
		return nil, err
	}
	allBooking := models.Bookings{}
	errselectall := db.All(&allBooking)
	if errselectall != nil {
		return nil, models.Error{500, "Can't Select data form Database"}
	}
	bookings := models.Bookings{}
	for _, value := range allBooking {
		sign, err := GetSignById(c, value.SignID)
		if err != nil {
			return nil, err
		}
		value.Sign = sign.(models.Sign)
		bookings = append(bookings, value)
	}
	return &bookings, nil
}

func CheckDate(D1 time.Time, D2 time.Time) int {
	diff := D2.Sub(D1)
	return int(diff.Hours()/24) + 1
}
