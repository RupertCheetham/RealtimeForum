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

var userConnections = make(map[int]*websocket.Conn)

// deals with the websocket side of chat
func WebsocketChatHandler(w http.ResponseWriter, r *http.Request) {

	//Need to redesign connections to use users, not chatrooms

	log.Println("[WebsocketChatHandler, chatHandler]  I'm here again; is that a problem?")
	// Upgrade the HTTP connection to a WebSocket connection
	connection := upgradeConnection(w, r)
	defer connection.Close()

	for {
		messageType, payload, err := connection.ReadMessage()
		if err != nil {
			utils.HandleError("Problem with the connections in WebsocketHandler.  Probably need to figure out a different way to remove them", err)
			// removeConnection(chatUUID, connection)
			return
		}

		chatMsg := payloadUnmarshaller(payload)

		previousChatEntryFound, chatUUID, err := previousChatChecker(chatMsg.Sender, chatMsg.Recipient)
		if err != nil {
			utils.HandleError("Error with previousChatChecker in ChatHandler", err)
		}

		if !previousChatEntryFound {
			chatUUID = utils.GenerateNewUUID()
			err = db.AddChatToDatabase(chatUUID, "", chatMsg.Sender, chatMsg.Recipient)
			if err != nil {
				utils.HandleError("There has been an issue with AddChatToDatabase in ChatHandler", err)
			}
		}

		// Check if the connection already exists in the chat
		userConnections[chatMsg.Sender] = connection

		if messageType == websocket.TextMessage {

			// Handle the message and broadcast it to other clients if needed
			if chatMsg.Type == "chat" {

				err = db.AddChatToDatabase(chatUUID, chatMsg.Body, chatMsg.Sender, chatMsg.Recipient)
			}
			if err != nil {
				utils.HandleError("There has been an issue with AddChatToDatabase in ChatHandler", err)
			}
		}
		if chatMsg.Type == "chat" {

			// Write message back to each client
			broadcastToUsers(messageType, chatMsg)
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

// func connectionExists(SenderID int, conn *websocket.Conn) bool {
// 	// Check if the connection exists for the given SenderID
// 	existingConn, ok := userConnections[SenderID]
// 	if !ok {
// 		return false
// 	}

// 	// Check if the existing connection matches the provided conn
// 	return existingConn == conn
// }

func broadcastToUsers(messageType int, message db.ChatMessage) {

	// changes the message to a JSON, for sending to frontend
	payload, err := json.Marshal(message)
	if err != nil {
		utils.HandleError("There has been an issue with marshalling in broadcastToUsers", err)
	}

	// Obtains the connections for each user involved in the chat
	// Should be simple enough to turn it  into an array of connections later, to enable group chats
	// Currently needs to only send messages if the user is online
	senderConnection := userConnections[message.Sender]
	err = senderConnection.WriteMessage(messageType, payload)
	if err != nil {
		utils.HandleError("Error sending message to a client in broadcastToChat:", err)
	}

	recipientConnection := userConnections[message.Recipient]
	// if recipientConnection exists (if they're online) then send a message to them too
	if recipientConnection != nil {
		err = recipientConnection.WriteMessage(messageType, payload)
		if err != nil {
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

// // remove an unused connection from userConnections
// func removeConnection(senderID int, connectionToRemove *websocket.Conn) {
// 	connections, ok := userConnections[senderID]
// 	if !ok {
// 		utils.HandleError("Chat UUID not found in map in removeConnection, in chatHandler", nil)
// 		return
// 	}

// 	// Find the index of the connection to remove
// 	index := -1
// 	for i, conn := range connections {
// 		if conn == connectionToRemove {
// 			index = i
// 			break
// 		}
// 	}

// 	// If the connection was found, remove it
// 	if index != -1 {
// 		userConnections[senderID] = append(connections[:index], connections[index+1:]...)
// 		log.Println("Removed closed connection from userConnections")
// 	}

// }

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
