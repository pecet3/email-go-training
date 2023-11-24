package models

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Config() {
	var err error
	db, err = sql.Open("postgres", "user=your_user_name dbname=your_database_name sslmode=disable")
	if err != nil {
		log.Println("config error:", err)
		return
	}

	err = createTables()
	if err != nil {
		log.Println("config error:", err)
		return
	}
}

func createTables() error {
	statement, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS emailAddresses (
			id SERIAL PRIMARY KEY,
			name TEXT,
			email TEXT
	);`)

	if err != nil {
		return err
	}
	if _, err := statement.Exec(); err != nil {
		return err
	}

	return nil
}
