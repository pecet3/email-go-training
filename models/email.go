package models

import (
	"log"
)

type Email struct {
	Id    int
	Name  string
	Email string
}

func (e *Email) CreateEmailAddress() error {

	_, err := db.Exec("INSERT INTO emailAddresses (name, email) VALUES ($1, $2)", e.Name, e.Email)
	if err != nil {
		log.Println("Failed to insert email address:", err)
		return err
	}

	log.Println("Email address inserted successfully")
	return nil
}

func GetEmailAddresses() ([]Email, error) {
	rows, err := db.Query("SELECT id, name, email FROM emailAddresses")
	if err != nil {
		log.Println("Failed to execute query:", err)
		return nil, err
	}
	defer rows.Close()

	var emails []Email

	for rows.Next() {
		var e Email
		if err := rows.Scan(&e.Id, &e.Name, &e.Email); err != nil {
			log.Println("Failed to scan row:", err)
			return nil, err
		}
		emails = append(emails, e)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error in result set:", err)
		return nil, err
	}

	return emails, nil
}
