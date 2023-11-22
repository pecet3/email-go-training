package main

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func sendMail() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error dotenv")
	}

	googlePassword := os.Getenv("GOOGLE_PASSWORD")

	auth := smtp.PlainAuth(
		"",
		"poznajprawdejp2@gmail.com",
		googlePassword,
		"smtp.gmail.com",
	)

	msg := "Subject: hello from go lang! \r\n\r\n hello my dad, sending from gmail with go lang works! have a good day and keep the code. please do not forget our blog project ;)"

	// Connect to the server, enable TLS, and send the mail
	err = smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"poznajprawdejp2@gmail.com",
		[]string{"pecet3107@gmail.com", "g.pacewicz@gmail.com"},
		[]byte(msg),
	)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	sendMail()
}
