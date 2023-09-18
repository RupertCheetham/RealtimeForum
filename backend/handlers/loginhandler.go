package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"realtimeForum/backend/db"
)

// Handler to login page
func AddLoginHandler(w http.ResponseWriter, r *http.Request) {
	// Enable CORS headers for this handler
	SetupCORS(&w, r)

	var login db.RegistrationEntry
	if r.Method == "POST" {

		err := json.NewDecoder(r.Body).Decode(&login)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		log.Println("Recieved login entry:", login.Username, login.Password)

		// isLoginSuccessful(login)
		db.GetLoginEntry(login)
		w.WriteHeader(http.StatusOK)
	}

}
