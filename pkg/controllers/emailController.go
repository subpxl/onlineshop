package controllers

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/smtp"
	"time"

	"github.com/go-chi/chi"
)

var tokenMap = make(map[string]bool)

func generateToken() string {
	rand.Seed(time.Now().UnixNano())
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	tokenLength := 32
	token := make([]byte, tokenLength)
	for i := range token {
		token[i] = charset[rand.Intn(len(charset))]
	}
	return string(token)
}

func sendVerificationEmail(userEmail string, token string) error {

	from := "your-email@gmail.com"
	password := ""
	smtpServer := "localhost"

	auth := smtp.PlainAuth("", from, password, smtpServer)

	to := []string{userEmail}
	subject := "Email Verification"
	body := fmt.Sprintf("Click the following link to verify your email: http://localhost:8000/verify/%s", token)

	msg := []byte("To: " + to[0] + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	err := smtp.SendMail(smtpServer+":1025", auth, from, to, msg)
	if err != nil {
		log.Printf("Email sending error: %v", err)
		return err
	}

	return nil
}

func (uh *UserHanlder) VerifyEmail(w http.ResponseWriter, r *http.Request) {
	// token := mux.Vars(r)["token"]
	token := chi.URLParam(r, "token")
	fmt.Println("token is ", token)

	// Check if the token is valid and hasn't been used before.
	if emailVerified := tokenMap[token]; emailVerified {
		fmt.Fprintf(w, "Email has already been verified with token: %s", token)
	} else {
		// Mark the email as verified in your database (for example, by setting a flag to true).
		tokenMap[token] = true
		fmt.Fprintf(w, "Email verified with token: %s", token)
	}

}
func (uh *UserHanlder) SendVerificationEmailHandler(w http.ResponseWriter, r *http.Request) {
	userEmail := "user@example.com" // Replace with the user's email
	token := generateToken()
	tokenMap[token] = false

	// Send the verification email with the generated token.
	err := sendVerificationEmail(userEmail, token)
	if err != nil {
		http.Error(w, "Failed to send verification email", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Verification email sent. Please check your inbox.")
}
