package handlers

import (
	"encoding/json"
	"net/http"
	"realtimeForum/db"
)

// Handler for getting post from DB
func GetUsernameHandler(w http.ResponseWriter, r *http.Request) {
	// Enable CORS headers for this handler
	SetupCORS(&w, r)

	// This code block is handling the logic for retrieving posts from the database when the HTTP request
	// method is GET.
	if r.Method == http.MethodGet { // Use http.MethodGet constant for clarity
		sessionID := r.URL.Query().Get("sessionID")
		username := db.GetUsernameFromSessionID(sessionID)

		// Set the response content type to JSON
		w.Header().Set("Content-Type", "application/json")

		// Encode and send the username as JSON in the response
		json.NewEncoder(w).Encode(username)
	}
}
