package main

import "onlineshop/pkg/controllers"

func getEmail() {
	emailConfigData := map[string]interface{}{
		"Host":     "localhost",
		"Port":     1025,
		"Username": "",
		"Password": "",
	}

	emailHandler := controllers.NewEmailConfig(emailConfigData)

	from := "your-email@example.com"
	to := "recipient@example.com"
	subject := "Test thing Subject"
	body := "This is the email body from tha th thins."
	emailHandler.SendEmail(from, to, subject, body)
}
