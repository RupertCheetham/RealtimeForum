package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"realtimeForum/db"
)

// Handler to login page
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Enable CORS headers for this handler
	SetupCORS(&w, r)
	fmt.Println("I'm in LoginHandler")
	var login db.UserEntry
	if r.Method == "POST" {

		err := json.NewDecoder(r.Body).Decode(&login)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		log.Println("Recieved login entry:", login.Username, login.Password)

		// isLoginSuccessful(login)
		db.VerifyUser(login)
		w.WriteHeader(http.StatusCreated)
	}
}
