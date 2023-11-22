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

	msg := "Subject: test \n hello world"

	err = smtp.SendMail(
		"smtp.google.com",
		auth,
		"poznajprawdejp2@gmail.com",
		[]string{"pecet3107@gmail.com"},
		[]byte(msg),
	)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	sendMail()
}
