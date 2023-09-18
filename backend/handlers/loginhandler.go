package handlers

import (
	"encoding/json"
	"net/http"
	"realtimeForum/db"
)

// Handler to login page
func AddLoginHandler(w http.ResponseWriter, r *http.Request) {
	// Enable CORS headers for this handler
	SetupCORS(&w, r)

	var login db.UserEntry
	if r.Method == "POST" {

		err := json.NewDecoder(r.Body).Decode(&login)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		msg := db.GetLoginEntry(login)
		jsonResponse, err := json.Marshal(msg)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}

	if r.Method == "GET" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized access"))
	}

}
