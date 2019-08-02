package mailers

import (
	"log"

	"github.com/gobuffalo/buffalo/mail"
	"github.com/gobuffalo/buffalo/render"
	"github.com/pkg/errors"
)

// func TextEmail
func SendWelcomeEmails(Subject string, email string) error {
	m := mail.NewMessage()

	// fill in with your stuff:
	m.Subject = Subject
	m.From = "panupong.jkn@gmail.com"
	log.Print("email:", email)
	m.To = []string{email}
	err := m.AddBody(r.HTML("welcome_email.html"), render.Data{})
	if err != nil {
		return err
	}
	err = smtp.Send(m)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
