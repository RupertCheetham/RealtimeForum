package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type ScoreEntry struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Score int    `json:"score"`
	Time  int    `json:"time"`
}

func setupCORS(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func AddScoreHandler(w http.ResponseWriter, r *http.Request) {
	// Enable CORS headers for this handler
	setupCORS(&w, r)

	if r.Method == "POST" {
		var score ScoreEntry
		err := json.NewDecoder(r.Body).Decode(&score)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Println("Received score:", score.Name, score.Score, score.Time)

		err = addScoreToDatabase(score.Name, score.Score, score.Time)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}

	if r.Method == "GET" {
		scores, err := getScoresFromDatabase()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if len(scores) > 0 {
			json.NewEncoder(w).Encode(scores)
		} else {
			w.Write([]byte("No scores available"))
		}
	}

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Super Mario API"))
}

func main() {
	initDatabase()
	log.Println("Database initialized successfully")

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/scores", AddScoreHandler)

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
