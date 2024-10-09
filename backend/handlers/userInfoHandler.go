package handlers

import (
	"encoding/json"
	"net/http"
	"realtimeForum/db"
)

// Handler for getting username from userID
func GetUsernameFromIDHandler(w http.ResponseWriter, r *http.Request) {
	// Enable CORS headers for this handler
	SetupCORS(&w, r)

	if r.Method == http.MethodGet { // Use http.MethodGet constant for clarity
		userID := r.URL.Query().Get("userID")

		username := db.GetUsernameFromUserID(userID)
		// Set the response content type to JSON
		w.Header().Set("Content-Type", "application/json")

		// Encode and send the username as JSON in the response
		json.NewEncoder(w).Encode(username)
	}
}
