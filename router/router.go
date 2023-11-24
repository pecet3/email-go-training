package router

import (
	"log"
	"mails/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupAndRun() {
	mux := mux.NewRouter()
	port := ":6660"

	mux.HandleFunc("/email", controllers.AddEmail).Methods("POST", "GET")

	log.Println("Starting the server at port", port)
	log.Fatal(http.ListenAndServe(port, mux))
}
