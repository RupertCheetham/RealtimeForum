package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"realtimeForum/db"
	"realtimeForum/utils"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

var chatConnections = make(map[string][]*websocket.Conn)

// deals with the websocket side of chat
func ChatHandler(w http.ResponseWriter, r *http.Request) {

	// Upgrade the HTTP connection to a WebSocket connection
	connection := upgradeConnection(w, r)
	defer connection.Close()

	sender, recipient := obtainSenderAndRecipient(r)

	previousChatEntryFound, chatUUID, err := previousChatChecker(sender, recipient)
	if err != nil {
		utils.HandleError("Error with previousChatChecker in ChatHandler", err)
	}
	// if chat is new then generates new UUID for chat
	if !previousChatEntryFound {
		chatUUID = utils.GenerateNewUUID()
		// fmt.Println("Generated UUID:", chatUUID)
		err = db.AddChatToDatabase(chatUUID, "", sender, recipient)

		if err != nil {
			utils.HandleError("There has been an issue with AddChatToDatabase in ChatHandler", err)
		}
	}

	// add current connection to list of connections that use this chatUUID
	chatConnections[chatUUID] = append(chatConnections[chatUUID], connection)

	for {
		messageType, payload, err := connection.ReadMessage()
		if err != nil {
			removeConnection(chatUUID, connection)
			return
		}

		chatMsg := payloadUnmarshaller(payload)
		// var chatUUID string

		if messageType == websocket.TextMessage {
			// log.Println("Message type is: ", chatMsg.Type)
			// Process the incoming message
			// fmt.Println("Received a WebSocket message:", string(payload))
			// Handle the message and broadcast it to other clients if needed
			if chatMsg.Type == "chat" {
				err = db.AddChatToDatabase(chatUUID, chatMsg.Body, chatMsg.Sender, chatMsg.Recipient)
			}
			if err != nil {
				utils.HandleError("There has been an issue with AddChatToDatabase in ChatHandler", err)
			}

			// fmt.Println("Type:", chatMsg.Type, "Message:", chatMsg.Body, "Sender:", chatMsg.Sender, "Recipient:", chatMsg.Recipient, "Time:", chatMsg.Time)

		}
		if chatMsg.Type == "chat" {
			payload, err = json.Marshal(chatMsg)
			if err != nil {
				utils.HandleError("There has been an issue with marshalling in ChatHandler", err)
			}
			// Write message back to each client
			broadcastToChat(chatUUID, messageType, payload)
		}

	}

}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// Allow WebSocket connections from http://localhost:8080
		allowedOrigins := []string{
			"https://localhost:8080", //backend
			"https://localhost:3000", //frontend
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

func obtainSenderAndRecipient(r *http.Request) (int, int) {
	senderString := r.URL.Query().Get("sender")
	recipientString := r.URL.Query().Get("recipient")
	sender, err := strconv.Atoi(senderString)
	if err != nil {
		utils.HandleError("There has been an issue with Atoi sender in obtainSenderAndRecipient", err)
	}
	recipient, err := strconv.Atoi(recipientString)
	if err != nil {
		utils.HandleError("There has been an issue with Atoi recipient in obtainSenderAndRecipient", err)
	}
	return sender, recipient
}

// Checks to see if chat between two users has taken place before, if so then returns chat UUID
func previousChatChecker(firstID int, secondID int) (bool, string, error) {
	query := `
	SELECT ChatUUID
	FROM CHAT
	WHERE (SenderID = ? AND RecipientID = ?) OR (SenderID = ? AND RecipientID = ?)
	LIMIT 1
`

	// Execute the query and try to fetch the ChatUUID
	var chatUUID string
	err := db.Database.QueryRow(query, firstID, secondID, secondID, firstID).Scan(&chatUUID)
	if err != nil {
		if err == sql.ErrNoRows {
			// Entry doesn't exist
			return false, "", nil
		}
		return false, "", err
	}

	// Entry exists, return true and the ChatUUID
	return true, chatUUID, nil
}

func broadcastToChat(chatUUID string, messageType int, payload []byte) {
	// Retrieve the list of connections for the specified chatUUID
	connections, ok := chatConnections[chatUUID]
	if !ok {
		utils.HandleError("There were no connections that use this chatUUID in broadcastToChat", nil)
		// Chat UUID not found in the map, handle this error if needed
		return
	}
	log.Println("broadcastToChat, length of connections:", len(connections))
	for _, conn := range connections {
		if err := conn.WriteMessage(messageType, payload); err != nil {
			// Handle the error if needed
			utils.HandleError("Error sending message to a client in broadcastToChat:", err)
		}
	}
}

// Retrieves chat history between two users based on the UserIDs in the URL
func GetChatHistoryHandler(w http.ResponseWriter, r *http.Request) {

	SetupCORS(&w, r)

	User1String := r.URL.Query().Get("user1")
	User2String := r.URL.Query().Get("user2")
	offsetString := r.URL.Query().Get("offset")
	limitString := r.URL.Query().Get("limit")
	user1, err := strconv.Atoi(User1String)

	// converts the values to ints
	if err != nil {
		utils.HandleError("There has been a problem with AtoI User1String in GetChatHistoryHandler", err)
	}
	user2, err := strconv.Atoi(User2String)
	if err != nil {
		utils.HandleError("There has been a problem with AtoI User2String in GetChatHistoryHandler", err)
	}
	offset, err := strconv.Atoi(offsetString)
	if err != nil {
		utils.HandleError("There has been a problem with AtoI offsetString in GetChatHistoryHandler", err)
	}
	limit, err := strconv.Atoi(limitString)
	if err != nil {
		utils.HandleError("There has been a problem with AtoI limitString in GetChatHistoryHandler", err)
	}

	var chatHistory []db.ChatMessage

	previousChatEntryFound, chatUUID, err := previousChatChecker(user1, user2)
	if err != nil {
		utils.HandleError("There has been a problem with previousChatChecker in GetChatHistoryHandler", err)
	}

	if previousChatEntryFound {
		chatHistory, err = db.GetChatFromDatabase(chatUUID, offset, limit)
	}
	if err != nil {
		utils.HandleError("Error retrieving chat history from the database:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Serialize chat history to JSON and send it as the response
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(chatHistory); err != nil {
		utils.HandleError("Error encoding chat history to JSON:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// remove an unused connection from chatConnections
func removeConnection(chatUUID string, connectionToRemove *websocket.Conn) {
	connections, ok := chatConnections[chatUUID]
	if !ok {
		utils.HandleError("Chat UUID not found in map in removeConnection, in chatHandler", nil)
		return
	}

	// Find the index of the connection to remove
	index := -1
	for i, conn := range connections {
		if conn == connectionToRemove {
			index = i
			break
		}
	}

	// If the connection was found, remove it
	if index != -1 {
		chatConnections[chatUUID] = append(connections[:index], connections[index+1:]...)
		log.Println("Removed closed connection from chatConnections")
	}

}

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

// unpacks websocket's Payload and adds time for return
func payloadUnmarshaller(payload []byte) db.ChatMessage {
	var chatMsg db.ChatMessage

	err := json.Unmarshal([]byte(payload), &chatMsg)
	if err != nil {
		utils.HandleError("Error unmarshaling JSON in payloadUnmarshaller:", err)
		return chatMsg
	}
	currentTime := time.Now()
	formattedTime := currentTime.Format("15:04:05 02-01-2006")
	chatMsg.Time = formattedTime // Format for date and time up to seconds

	return chatMsg
}
