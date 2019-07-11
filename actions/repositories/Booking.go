package repositories

import (
	"time"

	"github.com/JewlyTwin/be_booking_sign/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

func AddBooking(c buffalo.Context) (*models.Booking, interface{}) {
	db, err := ConnectDB(c)
	if err != nil {
		return nil, err
	}
	data := DynamicPostForm(c)
	code := data["signname"].(string) + "CODE" + data["firstdate"].(string) + data["lastdate"].(string)
	signInterface, err := GetSignByName(c)
	if err != nil {
		return nil, err
	}
	sign := signInterface.(*models.Sign)
	firstdate, _ := time.Parse("2006-01-02", data["firstdate"].(string))
	lastdate, _ := time.Parse("2006-01-02", data["lastdate"].(string))
	check := CheckBookingTime(firstdate, lastdate, sign.ID, db)
	if check != true {
		return nil, models.Error{400, "The date of booking is not available"}
	}
	newBooking := models.Booking{Code: code, Applicant: data["applicant"].(string), Organization: data["organization"].(string), FirstDate: firstdate, LastDate: lastdate, SignID: sign.ID}
	db.Create(&newBooking)
	return &newBooking, nil
}

func CheckBookingTime(f time.Time, l time.Time, signid int, db *pop.Connection) bool {
	bookings := []models.Booking{}
	_ = db.Q().Where("last_date >= (?) and first_date <= (?) and sign_id = (?)", f, l, signid).All(&bookings)
	if len(bookings) != 0 {
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
	db.All(&allBooking)
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
