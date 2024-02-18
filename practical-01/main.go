package main

import (
	"errors"
	"fmt"
	"regexp"
)

type EmailService struct {
	DbConnectionString string
}

type Email struct {
	Sender    string
	Recipient string
	Message   string
}

type EmailSender interface {
	Send(email Email, protocol string) error
	Validate(email Email) error
	Store(email Email, sender, recipient string)
}

func (emailService *EmailService) Send(email Email, protocol string) error {
	if err := emailService.Validate(email); err != nil {
		fmt.Printf("Error validating email: %v\n", err)
		return err
	}

	switch protocol {
	case "SMTP":
		 emailService.sendSMTP(email)
	case "IMAP":
		 emailService.sendIMAP(email)
	case "POP3":
		 emailService.sendPOP3(email)
	default:
		 emailService.sendSMTP(email)
	}

	return nil
}

func (es *EmailService) Validate(email Email) error {
	if email.Sender == "" || email.Recipient == "" {
		return errors.New("sender or recipient cannot be empty")
	}

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	isMatching := emailRegex.MatchString(email.Sender) && emailRegex.MatchString(email.Recipient)

	if !isMatching {
		return errors.New("the email provided is not correct")
	}

	return nil
}

func (es *EmailService) Store(email Email, sender, recipient string) {
	fmt.Printf("Storing email from %s to %s\n", sender, recipient)
}

func (es *EmailService) sendSMTP(email Email) {
	fmt.Printf("Sending email via SMTP: %+v\n", email)
}

func (es *EmailService) sendIMAP(email Email) {
	fmt.Printf("Sending email via IMAP: %+v\n", email)
}

func (es *EmailService) sendPOP3(email Email) {
	fmt.Printf("Sending email via POP3: %+v\n", email)
}

func main() {
	emailSender := &EmailService{DbConnectionString: "some_connection_string"}

	email := Email{
		Sender:    "lets@go.com",
		Recipient: "davd00@vse.cz",
		Message:   "Tak tohle je email!",
	}

	protocol := "SMTP"

	if err := emailSender.Send(email, protocol); err != nil {
		fmt.Printf("Error sending email: %v\n", err)
		return
	}

	emailSender.Store(email, email.Sender, email.Recipient)
}