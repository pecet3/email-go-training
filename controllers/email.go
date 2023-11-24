package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"mails/models"
	"net/http"
)

func AddEmail(w http.ResponseWriter, r *http.Request) {
	email := &models.Email{}
	// Odczytanie ciała żądania
	err := json.NewDecoder(r.Body).Decode(&email)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Dodanie adresu e-mail do bazy danych

	err = email.CreateEmailAddress()

	if err != nil {
		log.Println("Failed to create email address:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Odpowiedź sukcesem
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Email address added successfully")
}
