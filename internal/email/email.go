package email

import "github.com/go-gomail/gomail"

type Emailer struct {
	Sender   string
	Passcode string
}

func (e *Emailer) SendEmail(to, subject, body, attachmentPath string) error {
	m := gomail.NewMessage()

	// Set sender and recipient
	m.SetHeader("From", e.Sender)
	m.SetHeader("To", to)

	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	// Attaching the PDF file to the email
	m.Attach(attachmentPath)

	d := gomail.NewDialer("smtp.gmail.com", 587, e.Sender, e.Passcode)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
