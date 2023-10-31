package handlers

import (
	"encoding/json"
	"net/http"
	"realtimeForum/db"
)

// // Handler for getting username from session
// func GetUsernameHandler(w http.ResponseWriter, r *http.Request) {
// 	// Enable CORS headers for this handler
// 	SetupCORS(&w, r)

// 	// This code block is handling the logic for retrieving posts from the database when the HTTP request
// 	// method is GET.
// 	if r.Method == http.MethodGet { // Use http.MethodGet constant for clarity
// 		sessionID := r.URL.Query().Get("sessionID")
// 		username := db.GetUsernameFromSessionID(sessionID)

// 		// Set the response content type to JSON
// 		w.Header().Set("Content-Type", "application/json")

//			// Encode and send the username as JSON in the response
//			json.NewEncoder(w).Encode(username)
//		}
//	}
//
// Handler for getting post from DB

// Handler for getting userID from session
func GetUserIDHandler(w http.ResponseWriter, r *http.Request) {
	// Enable CORS headers for this handler
	SetupCORS(&w, r)

	// This code block is handling the logic for retrieving posts from the database when the HTTP request
	// method is GET.
	if r.Method == http.MethodGet { // Use http.MethodGet constant for clarity
		sessionID := GetCookie(w, r)
		userID := db.GetUserIDFromSessionID(sessionID) //

		// Set the response content type to JSON
		w.Header().Set("Content-Type", "application/json")

		// Encode and send the username as JSON in the response
		json.NewEncoder(w).Encode(userID)
	}
}

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
