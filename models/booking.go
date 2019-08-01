package models

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
)

type Booking struct {
	ID          int       `json:"id" db:"id"`
	Code        string    `json:"booking_code" db:"booking_code"`
	ApplicantID uuid.UUID `json:"applicant_id" db:"applicant_id" fk_id:"id"`
	SignID      int       `json:"sign_id" db:"sign_id" fk_id:"id"`
	Description string    `json:"description" db:"description"`
	FirstDate   time.Time `json:"first_date" db:"first_date"`
	LastDate    time.Time `json:"last_date" db:"last_date"`
	Status      string    `json:"status" db:"status"`
	Comment     string    `json:"comment" db:"comment"`
	Applicant   User      `json:"applicant" db:"-"`
	Sign        Sign      `json:"sign" db:"-"`
	CreatedAt   time.Time `json:"-" db:"created_at"`
	UpdatedAt   time.Time `json:"-" db:"updated_at"`
}

type BookingDay struct {
	Firstdate time.Time `json:"firstdate"`
	Lastdate  time.Time `json:"lastdate"`
}

type BookingDays []BookingDay

// String is not required by pop and may be deleted
func (b Booking) String() string {
	jb, _ := json.Marshal(b)
	return string(jb)
}

// Bookings is not required by pop and may be deleted
type Bookings []Booking

// String is not required by pop and may be deleted
func (b Bookings) String() string {
	jb, _ := json.Marshal(b)
	return string(jb)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (b *Booking) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (b *Booking) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (b *Booking) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (b *Booking) CreateModel(data map[string]interface{}, code string) bool {
	if data["applicant_id"] == nil {
		return false
	}
	if data["sign_id"] == nil {
		return false
	}
	if data["description"] == nil {
		return false
	}
	if data["first_date"] == nil {
		return false
	}
	if data["last_date"] == nil {
		return false
	}
	b.Code = code
	b.ApplicantID, _ = uuid.FromString(data["applicant_id"].(string))
	b.SignID, _ = strconv.Atoi(data["sign_id"].(string))
	b.Description = data["description"].(string)
	b.FirstDate, _ = time.Parse("2006-01-02", data["first_date"].(string))
	b.LastDate, _ = time.Parse("2006-01-02", data["last_date"].(string))
	b.Status = "pending"
	b.Comment = ""
	return true
}

// func (b *Booking) CreateBookingModel(data map[string]interface{}, code string, sign Sign) {
// 	b.Code = code
// 	b.Applicant = data["applicant"].(string)
// 	b.Organization = data["organization"].(string)
// 	b.FirstDate, _ = time.Parse("2006-01-02", data["firstdate"].(string))
// 	b.LastDate, _ = time.Parse("2006-01-02", data["lastdate"].(string))
// 	b.SignID = sign.ID
// 	b.Sign = Sign{Name: sign.Name, Location: sign.Location, Limitdate: sign.Limitdate, Beforebooking: sign.Beforebooking}
// }

// func (b *Bookings) CreateBookingDays() interface{} {
// 	days := BookingDays{}
// 	newBookings := b
// 	for value, _ := range  {
// 		log.Print(value)
// 		// days = append(days, BookingDay{})
// 	}
// 	return days
// }

func (b *Booking) ReturnJsonID() IDbooking {
	idbook := IDbooking{b.ID}
	return idbook
}

type IDbooking struct {
	ID int `json:"id"`
}

type Allbooking struct {
	Booking Bookings `json:"bookings"`
}
