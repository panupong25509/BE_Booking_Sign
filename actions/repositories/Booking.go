package repositories

import (
	"time"

	"github.com/JewlyTwin/be_booking_sign/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	// "log"
)

func AddBooking(c buffalo.Context) interface{} {
	db := ConnectDB(c).(*pop.Connection)
	data := DynamicPostForm(c)
	code := data["signname"].(string) + "CODE" + data["firstdate"].(string) + data["lastdate"].(string)
	signID := GetSignByName(c).(*models.Sign)
	firstdate, _ := time.Parse("2006-01-02", data["firstdate"].(string))
	lastdate, _ := time.Parse("2006-01-02", data["lastdate"].(string))
	check := CheckBookingTime(firstdate, lastdate, db)
	if check != true {
		return map[string]interface{}{"error": "The date of booking is not available."}
	}
	newBooking := models.Booking{Code: code, Applicant: data["applicant"].(string), Organization: data["organization"].(string), FirstDate: firstdate, LastDate: lastdate, SignID: signID.ID}
	db.Create(&newBooking)
	return &newBooking
}

func CheckBookingTime(f time.Time, l time.Time, db *pop.Connection) bool {
	bookings := []models.Booking{}
	_ = db.Q().Where("last_date >= (?) and first_date <= (?)", f, l).All(&bookings)
	if len(bookings) != 0 {
		return false
	}
	return true
}

func GetAllBooking(c buffalo.Context) interface{} {
	db := ConnectDB(c).(*pop.Connection)
	allBooking := []models.Booking{}

	db.All(&allBooking)
	return &allBooking
}
