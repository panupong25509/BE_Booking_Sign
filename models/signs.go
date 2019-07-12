package models

import (
	"encoding/json"
	"time"
	"strconv"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
)

type Sign struct {
	ID        			int       `json:"id" db:"id"`
	Name      			string    `json:"name" db:"sign_name"`
	Location  			string    `json:"location" db:"location"`
	Limitdate  			int    		`json:"limitdate" db:"limitdate"`
	Beforebooking	  int 		  `json:"beforebooking" db:"beforebooking"`
	Booking   []Booking `json:"booking" db:"-"  has_many:"bookings"`
	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (s Sign) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Signs is not required by pop and may be deleted
type Signs []Sign

// String is not required by pop and may be deleted
func (s Signs) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (s *Sign) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (s *Sign) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (s *Sign) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}


func (s *Sign) AfterFind()  error {
	return nil
}


func (s *Sign) CheckParamPostForm(data map[string]interface{})  bool {
	if data["name"] == nil {
		return false
	}
	if data["location"] == nil {
		return false
	}
	return true
}

func (s *Sign) CreateBookingModel(data map[string]interface{}) {
	s.Name = data["name"].(string)
	s.Location = data["location"].(string)
	s.Limitdate, _ = strconv.Atoi(data["limitdate"].(string))
	s.Beforebooking, _ = strconv.Atoi(data["beforebooking"].(string))
}