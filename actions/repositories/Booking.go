package repositories

import (
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
	newBooking := models.Booking{}
	if !newBooking.CheckParamPostForm(data) {
		return nil, models.Error{400, "param not complete"}
	}
	code := GenCodeBooking(data)
	signInterface, err := GetSignByName(c)
	if err != nil {
		return nil, err
	}
	sign := signInterface.(*models.Sign)
	newBooking.CreateBookingModel(data, code, sign.ID)
	if !ValidateBookingTime(&newBooking, db) {
		return nil, models.Error{400, "The date of booking is not available"}
	}
	db.Create(&newBooking)
	return &newBooking, nil
}

func GenCodeBooking(data map[string]interface{}) string {
	code := data["signname"].(string) + "CODE" + data["firstdate"].(string) + data["lastdate"].(string)
	return code
}

func ValidateBookingTime(newBooking *models.Booking, db *pop.Connection) bool {
	bookings := models.Bookings{}
	_ = db.Q().Where("last_date >= (?) and first_date <= (?) and sign_id = (?)", newBooking.FirstDate, newBooking.LastDate, newBooking.SignID).All(&bookings)
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
