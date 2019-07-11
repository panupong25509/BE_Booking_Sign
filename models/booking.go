package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
)

type Booking struct {
	ID           int       `json:"id" db:"id"`
	Code         string    `json:"booking_code" db:"booking_code"`
	Applicant    string    `json:"applicant" db:"applicant"`
	Organization string    `json:"organization" db:"organization"`
	FirstDate    time.Time `json:"first_date" db:"first_date"`
	LastDate     time.Time `json:"last_date" db:"last_date"`
	SignID       int       `json:"sign_id" db:"sign_id" fk_id:"id"`
	Sign         Sign      `json:"sign" db:"-"`
	CreatedAt    time.Time `json:"-" db:"created_at"`
	UpdatedAt    time.Time `json:"-" db:"updated_at"`
}

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

func (b *Booking) CheckParamPostForm(data map[string]interface{}) bool {
	if data["signname"] == nil {
		return false
	}
	if data["applicant"] == nil {
		return false
	}
	if data["organization"] == nil {
		return false
	}
	if data["firstdate"] == nil {
		return false
	}
	if data["lastdate"] == nil {
		return false
	}
	return true
}
func (b *Booking) CreateBookingModel(data map[string]interface{}, code string, signid int) {
	b.Code = code
	b.Applicant = data["applicant"].(string)
	b.Organization = data["organization"].(string)
	b.FirstDate, _ = time.Parse("2006-01-02", data["firstdate"].(string))
	b.LastDate, _ = time.Parse("2006-01-02", data["lastdate"].(string))
	b.SignID = signid
}
