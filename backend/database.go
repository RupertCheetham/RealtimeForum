package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDatabase() {
	var err error
	db, err = sql.Open("sqlite3", "registration.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create the registration table if it doesn't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS registration (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			nickname TEXT,
			age INTEGER,
			gender TEXT,
			first name Text,
			last name TEXT,
			email TEXT,
			password TEXT
		);
	`)
	if err != nil {
		log.Fatal(err)
	}
}

func addRegistrationToDatabase(nickname string, age int, gender string, firstName string, lastName string, email string, password string) error {
	_, err := db.Exec("INSERT INTO registration (nickname, age, gender, first name, last name, email, password) VALUES (?, ?, ?, ?, ?, ?, ?)", nickname, age, gender, firstName, lastName, email, password)
	if err != nil {
		log.Println("Error adding registration to database:", err)
	}
	return err
}

func getRegistrationFromDatabase() ([]RegistrationEntry, error) {
	rows, err := db.Query("SELECT nickname, age, gender, first name, last name, email, password FROM registration ORDER BY id ASC")
	if err != nil {
		log.Println("Error querying registrations from database:", err)
		return nil, err
	}
	defer rows.Close()

	var registrations []RegistrationEntry
	for rows.Next() {
		var entry RegistrationEntry
		err := rows.Scan(&entry.Nickname, &entry.Age, &entry.Gender, &entry.FirstName, &entry.LastName, &entry.Email, &entry.Password)
		if err != nil {
			log.Println("Error scanning row from database:", err)
			return nil, err
		}
		registrations = append(registrations, entry)
	}

	return registrations, nil
}
