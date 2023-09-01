package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type RegistrationEntry struct {
	ID        int    `json:"id"`
	Nickname  string `json:"nickname"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	FirstName string `json:"first name"`
	LastName  string `json:"last name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func setupCORS(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func AddRegistrationHandler(w http.ResponseWriter, r *http.Request) {
	// Enable CORS headers for this handler
	setupCORS(&w, r)

	if r.Method == "POST" {
		var registration RegistrationEntry
		err := json.NewDecoder(r.Body).Decode(&registration)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Println("Received score:", registration.Nickname, registration.Age, registration.Gender, registration.FirstName, registration.LastName, registration.Email, registration.Password)

		err = addRegistrationToDatabase(registration.Nickname, registration.Age, registration.Gender, registration.FirstName, registration.LastName, registration.Email, registration.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}

	if r.Method == "GET" {
		registrations, err := getRegistrationFromDatabase()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if len(registrations) > 0 {
			json.NewEncoder(w).Encode(registrations)
		} else {
			w.Write([]byte("No registrations available"))
		}
	}

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Real Time Forum API"))
}

func main() {
	initDatabase()
	log.Println("Database initialized successfully")

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/registrations", AddRegistrationHandler)

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
