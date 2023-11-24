package main

import (
	"mails/models"
	"mails/router"
	email "mails/utils"
)

func main() {

	go email.SendEmailEveryHour()
	models.Config()
	router.SetupAndRun()

	select {}

}
