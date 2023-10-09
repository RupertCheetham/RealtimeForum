package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"realtimeForum/db"
	"realtimeForum/utils"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// Allow WebSocket connections from http://localhost:8080
		allowedOrigins := []string{
			"http://localhost:8080", //backend
			"http://localhost:3000", //frontend
		}
		origin := r.Header.Get("Origin")
		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin {
				return true
			}
		}
		return false
	},
}

func ChatHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("gwew")
	// Upgrade the HTTP connection to a WebSocket connection
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("There was an error with Upgrade in WebSocketHandler,", err)
		return
	}
	defer connection.Close()

	// Handle incoming and outgoing WebSocket messages here
	// Use Go channels to broadcast messages to all connected clients

	for {
		messageType, payload, err := connection.ReadMessage()
		if err != nil {
			log.Println("WebSocket read error:", err)
			return
		}
		if messageType == websocket.TextMessage {
			// Process the incoming message
			fmt.Println("Received a WebSocket message:", string(payload))
			// Handle the message and broadcast it to other clients if needed
			var chatMsg db.ChatMessage
			// Unmarshal the JSON into the struct
			err := json.Unmarshal([]byte(payload), &chatMsg)
			if err != nil {
				fmt.Println("Error unmarshaling JSON:", err)
				return
			}
			UUID := "lknd$sfvkljsdnfgsnlkdf"
			err = db.AddChatToDatabase(UUID, chatMsg.Message, chatMsg.Sender, chatMsg.Recipient)
			if err != nil {
				log.Println("There has been an issue with AddChatToDatabase in ChatHandler", err)
				utils.HandleError("There has been an issue with AddChatToDatabase in ChatHandler", err)
			}
			// Access the individual fields
			fmt.Println("Type:", chatMsg.Type)
			fmt.Println("Message:", chatMsg.Message)
			fmt.Println("Sender:", chatMsg.Sender)
			fmt.Println("Recipient:", chatMsg.Recipient)
		}
	}
}
