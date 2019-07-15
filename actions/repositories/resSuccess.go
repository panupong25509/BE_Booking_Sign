package repositories

import "github.com/JewlyTwin/be_booking_sign/models"

func resSuccess(m interface{}) models.Success {
	if m != nil {
		return models.Success{m.(string)}
	}
	return models.Success{"success"}
}
