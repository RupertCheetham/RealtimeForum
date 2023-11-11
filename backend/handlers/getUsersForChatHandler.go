package handlers

import (
	"encoding/json"
	"net/http"
	"realtimeForum/db"
	"realtimeForum/utils"

	"github.com/gorilla/websocket"
)

// Returns a list of registered users; we'll use this in the frontend to start a chat with one of them
func GetUsersForChatHandler(w http.ResponseWriter, r *http.Request) {
	// Enable CORS headers for this handler
	SetupCORS(&w, r)

	userID := r.URL.Query().Get("userId")

	// This code block is handling the logic for retrieving posts from the database when the HTTP request
	// method is GET.
	if r.Method == http.MethodGet { // Use http.MethodGet constant for clarity
		users, err := db.GetRecentChatUsersFromDatabase(userID)
		if err != nil {
			utils.HandleError("Error fetching users in GetUsersForChatHandler:", err)
		}

		// Set the response content type to JSON
		w.Header().Set("Content-Type", "application/json")

		// Encode and send the username as JSON in the response
		json.NewEncoder(w).Encode(users)
	}
}

// Upgrade the HTTP connection to a WebSocket connection
func upgradeConnection(w http.ResponseWriter, r *http.Request) *websocket.Conn {

	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		utils.HandleError("There was an error with Upgrade in upgradeConnection,", err)
		return nil
	}
	return connection
}
