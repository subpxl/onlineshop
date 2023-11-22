package controllers

import (
	"log"

	"github.com/go-gomail/gomail"
)

type EmailConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	// From     string
	// To       string
	// Subject  string
	// Body     string
}

func NewEmailConfig(data map[string]interface{}) *EmailConfig {

	return &EmailConfig{
		Host:     data["Host"].(string),     // Assuming 'Host' is a required field of string type
		Port:     data["Port"].(int),        // Assuming 'Port' is a required field of string type
		Username: data["Username"].(string), // Assuming 'Username' is a required field of string type
		Password: data["Password"].(string),
	}
}

func (eh *EmailConfig) SendEmail(from, to, subject, body string) (string, error) {
	// Set up the email message

	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	// Set up the SMTP server configuration for MailHog
	dialer := gomail.NewDialer(eh.Host, eh.Port, eh.Username, eh.Password) // MailHog's default SMTP server address and port

	// Send the email
	if err := dialer.DialAndSend(m); err != nil {
		log.Print("cannot start email server", err)

		return "failed to send email ", err

	}
	return "emain sent successfully", nil

}
