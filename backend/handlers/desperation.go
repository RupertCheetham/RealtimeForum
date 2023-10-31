package handlers

// import (
// 	"database/sql"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"realtimeForum/db"
// 	"realtimeForum/utils"
// 	"strconv"
// 	"time"

// 	"github.com/gorilla/websocket"
// )

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// 	CheckOrigin: func(r *http.Request) bool {
// 		// Allow WebSocket connections from http://localhost:8080
// 		allowedOrigins := []string{
// 			"https://localhost:8080", //backend
// 			"https://localhost:3000", //frontend
// 		}
// 		origin := r.Header.Get("Origin")
// 		for _, allowedOrigin := range allowedOrigins {
// 			if origin == allowedOrigin {
// 				return true
// 			}
// 		}
// 		return false
// 	},
// }

// // Maintain a map of connected clients
// var clients = make(map[*websocket.Conn]bool)

// // Maintain a map that associates chat UUIDs with connected clients
// var chatClientMap = make(map[string]map[*websocket.Conn]bool)

// // deals with the websocket side of chat
// func ChatHandler(w http.ResponseWriter, r *http.Request) {

// 	// Upgrade the HTTP connection to a WebSocket connection
// 	connection := upgradeConnection(w, r)

// 	// clients[connection] = true
// 	defer func() {
// 		delete(clients, connection)
// 		connection.Close()
// 	}()

// 	for {
// 		messageType, payload, err := connection.ReadMessage()
// 		if err != nil {
// 			log.Println("WebSocket read error:", err)
// 			return
// 		}

// 		chatMsg := payloadUnmarshaller(payload)

// 		if messageType == websocket.TextMessage {
// 			// Process the incoming message
// 			fmt.Println("Received a WebSocket message:", string(payload))
// 			// Handle the message and broadcast it to other clients if needed

// 			// Unmarshal the JSON into the struct

// 			previousChatEntryFound, chatUUID, err := previousChatChecker(chatMsg.Sender, chatMsg.Recipient)

// 			if err != nil {
// 				log.Println("Error with chatChecker in ChatHandler", err)
// 				utils.HandleError("Error with chatChecker in ChatHandler", err)
// 			}

// 			// if chat is new then generates new UUID for chat
// 			if !previousChatEntryFound {
// 				fmt.Println("chatIsNew", previousChatEntryFound)
// 				chatUUID = utils.GenerateNewUUID()
// 				fmt.Println("Generated UUID:", chatUUID)
// 			}

// 			currentTime := time.Now()
// 			formattedTime := currentTime.Format("15:04:05 02-01-2006") // Format for date and time up to seconds

// 			err = db.AddChatToDatabase(chatUUID, chatMsg.Message, chatMsg.Sender, chatMsg.Recipient)
// 			if err != nil {
// 				log.Println("There has been an issue with AddChatToDatabase in ChatHandler", err)
// 				utils.HandleError("There has been an issue with AddChatToDatabase in ChatHandler", err)
// 			}
// 			// Access the individual fields
// 			fmt.Println("Type:", chatMsg.Type)
// 			fmt.Println("Message:", chatMsg.Message)
// 			fmt.Println("Sender:", chatMsg.Sender)
// 			fmt.Println("Recipient:", chatMsg.Recipient)
// 			fmt.Println("Time:", formattedTime)

// 			// Print the message to the console
// 			fmt.Printf("%s sent: %s\n", connection.RemoteAddr(), string(chatMsg.Message))

// 			// for _, client := range clients {
// 			// 	// Write message back to browser
// 			// 	if err = client.WriteMessage(messageType, payload); err != nil {
// 			// 		return
// 			// 	}
// 			// }

// 		}

// 		payload, err = json.Marshal(chatMsg)
// 		if err != nil {
// 			log.Println("There has been an issue with marshalling in ChatHandler", err)
// 			utils.HandleError("There has been an issue with marshalling in ChatHandler", err)
// 		}
// 		if err = connection.WriteMessage(messageType, payload); err != nil {
// 			log.Println("WebSocket write error:", err)
// 			return
// 		}
// 	}

// }

// // Checks to see if chat between two users has taken place before, if so then returns chat UUID
// func previousChatChecker(firstID int, secondID int) (bool, string, error) {
// 	query := `
// 	SELECT ChatUUID
// 	FROM CHAT
// 	WHERE (SenderID = ? AND RecipientID = ?) OR (SenderID = ? AND RecipientID = ?)
// 	LIMIT 1
// `

// 	// Execute the query and try to fetch the ChatUUID
// 	var chatUUID string
// 	err := db.Database.QueryRow(query, firstID, secondID, secondID, firstID).Scan(&chatUUID)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			// Entry doesn't exist
// 			return false, "", nil
// 		}
// 		return false, "", err
// 	}

// 	// Entry exists, return true and the ChatUUID
// 	return true, chatUUID, nil
// }

// func GetChatHistoryHandler(w http.ResponseWriter, r *http.Request) {

// 	SetupCORS(&w, r)

// 	chatUser1 := r.URL.Query().Get("user1")
// 	chatUser2 := r.URL.Query().Get("user2")

// 	var user1 int
// 	var user2 int
// 	var chatHistory []db.ChatMessage
// 	var err error

// 	user1, _ = strconv.Atoi(chatUser1)
// 	user2, _ = strconv.Atoi(chatUser2)

// 	previousChatEntryFound, chatUUID, _ := previousChatChecker(user1, user2)

// 	if previousChatEntryFound {
// 		chatHistory, err = db.GetChatFromDatabase(chatUUID)
// 	}
// 	if err != nil {
// 		utils.HandleError("Error retrieving chat history from the database:", err)
// 		http.Error(w, "Internal server error", http.StatusInternalServerError)
// 		return
// 	}

// 	// Serialize chat history to JSON and send it as the response
// 	w.Header().Set("Content-Type", "application/json")
// 	if err = json.NewEncoder(w).Encode(chatHistory); err != nil {
// 		utils.HandleError("Error encoding chat history to JSON:", err)
// 		http.Error(w, "Internal server error", http.StatusInternalServerError)
// 		return
// 	}
// }

// func GetUsersForChatHandler(w http.ResponseWriter, r *http.Request) {
// 	// Enable CORS headers for this handler
// 	SetupCORS(&w, r)

// 	// This code block is handling the logic for retrieving posts from the database when the HTTP request
// 	// method is GET.
// 	if r.Method == http.MethodGet { // Use http.MethodGet constant for clarity
// 		users, err := db.GetUsersFromDatabase()
// 		if err != nil {
// 			utils.HandleError("Error fetching users in GetUsersForChatHandler:", err)
// 			log.Println("Error fetching users in GetUsersForChatHandler:", err)
// 		}

// 		// Set the response content type to JSON
// 		w.Header().Set("Content-Type", "application/json")

// 		// Encode and send the username as JSON in the response
// 		json.NewEncoder(w).Encode(users)
// 	}
// }

// // Handle incoming messages from clients
// func handleWebSocketConnection(conn *websocket.Conn, chatUUID string) {
// 	clients[conn] = true
// 	if _, ok := chatClientMap[chatUUID]; !ok {
// 		chatClientMap[chatUUID] = make(map[*websocket.Conn]bool)
// 	}
// 	chatClientMap[chatUUID][conn] = true

// 	defer func() {
// 		delete(clients, conn)
// 		delete(chatClientMap[chatUUID], conn)
// 		conn.Close()
// 	}()

// 	for {
// 		messageType, p, err := conn.ReadMessage()
// 		if err != nil {
// 			return
// 		}
// 		// Broadcast the message to clients associated with the same chat UUID
// 		for client := range chatClientMap[chatUUID] {
// 			if err := client.WriteMessage(messageType, p); err != nil {
// 				return
// 			}
// 		}
// 	}
// }

// // Upgrade the HTTP connection to a WebSocket connection
// func upgradeConnection(w http.ResponseWriter, r *http.Request) *websocket.Conn {

// 	connection, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		utils.HandleError("There was an error with Upgrade in upgradeConnection,", err)
// 		log.Println("There was an error with Upgrade in upgradeConnection,", err)
// 		return nil
// 	}
// 	defer connection.Close()
// 	return connection
// }

// // unpacks websocket's Payload and adds time for return
// func payloadUnmarshaller(payload []byte) db.ChatMessage {
// 	var chatMsg db.ChatMessage
// 	err := json.Unmarshal([]byte(payload), &chatMsg)
// 	if err != nil {
// 		utils.HandleError("Error unmarshaling JSON in payloadUnmarshaller:", err)
// 		log.Println("Error unmarshaling JSON in payloadUnmarshaller:", err)
// 		return chatMsg
// 	}

// 	currentTime := time.Now()
// 	formattedTime := currentTime.Format("15:04:05 02-01-2006")
// 	chatMsg.Time = formattedTime // Format for date and time up to seconds

// 	return chatMsg
// }
