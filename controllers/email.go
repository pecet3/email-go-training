package controllers

import (
	"fmt"
	"html/template"
	"log"
	"mails/models"
	"net/http"
)

func AddEmail(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Renderowanie strony HTML z formularzem
		tmpl, err := template.ParseFiles("view/index.html")
		if err != nil {
			log.Println("Failed to parse HTML template:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Println("Failed to execute HTML template:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		return
	}

	email := &models.Email{}

	err := r.ParseForm()
	if err != nil {
		log.Println("Failed to parse form:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Pobranie wartości z formularza
	email.Name = r.Form.Get("name")
	email.Email = r.Form.Get("email")

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
