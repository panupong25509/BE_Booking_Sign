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

// func (b *Booking) GetCSign(c *buffalo.Context) error {
// 	db, _ := c.Value("tx").(*pop.Connection)
// 	sign := Sign{}
// 	if len(s.Clubs) > 0 {
// 		for _, value := range s.Clubs {
// 			// log.Println(value.Club_id)
// 			// id, _ := uuid.FromString(value.Club_id)
// 			db.Find(&club ,value.Club_id)
// 			// log.Println(&club)
// 			s.Club = append(s.Club, ClubForShow{club.ID, club.Name})
// 		}
// 	}
// 	log.Println(s.Club)

// 	return nil
// }
