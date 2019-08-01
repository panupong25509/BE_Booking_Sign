package mailers

import (
	"github.com/gobuffalo/buffalo/mail"
	"github.com/gobuffalo/buffalo/render"
	"github.com/pkg/errors"
)

func SendWelcomeEmails() error {
	m := mail.NewMessage()

	// fill in with your stuff:
	m.Subject = "Welcome Email"
	m.From = "panupong.jkn@gmail.com"
	m.To = []string{"l2jew123@gmail.com"}
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
