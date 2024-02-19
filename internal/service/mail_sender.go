package service

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

const (
	smtpAuthMail   = "smtp.gmail.com"
	smtpAuthSender = "smtp.gmail.com:587"
)

type EmailSender struct {
	name              string
	fromEmailAddress  string
	fromEmailPassword string
}

func NewEmailSender(name, fromEmailAddress, fromEmailPassword string) *EmailSender {
	return &EmailSender{name: name, fromEmailAddress: fromEmailAddress, fromEmailPassword: fromEmailPassword}
}

func (r *EmailSender) Sendmail(subject string, content string, to []string, cc []string, bcc []string, attachFiles []string) error {
	mail := email.NewEmail()
	mail.From = fmt.Sprintf("%s <%s>", r.name, r.fromEmailAddress)
	mail.Subject = subject
	mail.HTML = []byte(content)
	mail.To = to
	mail.Cc = cc
	mail.Bcc = bcc
	for _, file := range attachFiles {
		_, err := mail.AttachFile(file)
		if err != nil {
			return err
		}
	}
	smtpAuth := smtp.PlainAuth("", r.fromEmailAddress, r.fromEmailPassword, smtpAuthMail)
	return mail.Send(smtpAuthSender, smtpAuth)
}
