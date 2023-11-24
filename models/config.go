package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func Config() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error dotenv")
	}
	passwordDb := os.Getenv("PASSWORD_DB")
	db, err = sql.Open("postgres", "user=u111 dbname=db_u111 host=psql01.mikr.us password="+passwordDb+" sslmode=disable")
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
