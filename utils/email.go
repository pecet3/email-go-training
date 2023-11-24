package email

import (
	"fmt"
	"log"
	"mails/models"
	"net/smtp"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func SendMail() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error dotenv")
	}
	startDate := time.Date(2005, time.April, 21, 37, 0, 0, 0, time.UTC)

	currentTime := time.Now()

	timeDifference := currentTime.Sub(startDate)

	days := int(timeDifference.Hours() / 24)

	hours := int(timeDifference.Seconds() / 3600)

	daysString := strconv.Itoa(days)

	hoursString := strconv.Itoa(hours)

	googlePassword := os.Getenv("GOOGLE_PASSWORD")

	auth := smtp.PlainAuth(
		"",
		"poznajprawdejp2@gmail.com",
		googlePassword,
		"smtp.gmail.com",
	)

	subject := "Informacja Papieska"
	// + currentTime.Format("2006-01-02 15:04:05")

	body := "Jan Paweł 2 nie żyje od " + daysString + " dni i " + hoursString + " godzin"

	msg := "Subject:" + subject + "\r\n\r\n" + body

	emailAddresses, err := models.GetEmailAddresses()

	var emails []string

	for _, e := range emailAddresses {
		emails = append(emails, e.Email)
	}

	err = smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"poznajprawdejp2@gmail.com",
		[]string{"pecet3107@gmail.com"},
		[]byte(msg),
	)
	if err != nil {
		log.Println(err)
	}
}

func SendEmailEveryHour() {
	ticker := time.NewTicker(time.Hour)

	for {
		<-ticker.C

		SendMail()
		log.Println("-> Email has been sent ")
	}

}
