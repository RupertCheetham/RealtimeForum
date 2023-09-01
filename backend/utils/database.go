package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDatabase() {
	var err error
	db, err = sql.Open("sqlite3", "scores.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create the scores table if it doesn't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS scores (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			score INTEGER,
			time INTEGER
		);
	`)
	if err != nil {
		log.Fatal(err)
	}
}

func addScoreToDatabase(name string, score int, time int) error {
	_, err := db.Exec("INSERT INTO scores (name, score, time) VALUES (?, ?, ?)", name, score, time)
	if err != nil {
		log.Println("Error adding score to database:", err)
	}
	return err
}

func getScoresFromDatabase() ([]ScoreEntry, error) {
	rows, err := db.Query("SELECT name, score, time FROM scores ORDER BY score DESC")
	if err != nil {
		log.Println("Error querying scores from database:", err)
		return nil, err
	}
	defer rows.Close()

	var scores []ScoreEntry
	for rows.Next() {
		var entry ScoreEntry
		err := rows.Scan(&entry.Name, &entry.Score, &entry.Time)
		if err != nil {
			log.Println("Error scanning row from database:", err)
			return nil, err
		}
		scores = append(scores, entry)
	}

	return scores, nil
}
