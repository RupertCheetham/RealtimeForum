package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"realtimeForum/db"
	"realtimeForum/utils"
	"time"

	"github.com/gorilla/websocket"
)

var userConnections = make(map[int]*websocket.Conn)
var onlineUserConnections = make(map[int]*websocket.Conn)

// deals with the websocket side of chat
func WebsocketHandler(w http.ResponseWriter, r *http.Request) {

	// Upgrade the HTTP connection to a WebSocket connection
	connection := upgradeConnection(w, r)
	defer connection.Close()

	for {
		// Read the contents of the websocket message
		messageType, payload, err := connection.ReadMessage()
		if err != nil {
			utils.HandleError("Problem with the connections in WebsocketHandler.", err)
			removeConnection(userConnections, connection)
			return
		}

		// Unmarshal the contents of the websocket message
		chatMsg := payloadUnmarshaller(payload)

		// Ensures that current connection (for the user) is the most recent one
		userConnections[chatMsg.Sender] = connection
		// Consider breaking down into functions

		if chatMsg.Type == "typing" {
			forwardTypingStatus(chatMsg)
		} else if chatMsg.Type == "chat" {
			var chatUUID string
			// finds users chatroom and then adds message to db
			log.Println("[WebsocketChatHandler] chat")
			previousChatEntryFound, chatUUID, err := db.PreviousChatChecker(chatMsg.Sender, chatMsg.Recipient)
			if err != nil {
				utils.HandleError("Error with previousChatChecker 1 in WebsocketChatHandler", err)
			}
			if !previousChatEntryFound {
				chatUUID = utils.GenerateNewUUID()
				log.Println("Initialising chat room between User", chatMsg.Sender, "and User", chatMsg.Recipient)
				// err = db.AddChatToDatabase(chatUUID, "", chatMsg.Sender, chatMsg.Recipient)
				if err != nil {
					utils.HandleError("There has been an issue with AddChatToDatabase in ChatHandler", err)
				}
			}
			err = db.AddChatToDatabase(chatUUID, chatMsg.Body, chatMsg.Sender, chatMsg.Recipient)
			if err != nil {
				utils.HandleError("There has been an issue with AddChatToDatabase in ChatHandler", err)
			}
			broadcastChatToUsers(messageType, chatMsg)
		} else if chatMsg.Type == "chat_init" {
			// checks to see if users have chatted previously, if not then creates a room in the db
			log.Println("[WebsocketChatHandler]  chat_init")
			previousChatEntryFound, _, err := db.PreviousChatChecker(chatMsg.Sender, chatMsg.Recipient)
			if err != nil {
				utils.HandleError("Error with previousChatChecker 2 in WebsocketChatHandler", err)
			}
			if !previousChatEntryFound {
				chatUUID := utils.GenerateNewUUID()
				log.Println("Initialising chat room between User", chatMsg.Sender, "and User", chatMsg.Recipient)
				err = db.AddChatToDatabase(chatUUID, "", chatMsg.Sender, chatMsg.Recipient)
				if err != nil {
					utils.HandleError("There has been an issue with AddChatToDatabase in ChatHandler", err)
				}
			}
		} else if chatMsg.Type == "user_online" {
			// If user's connection isn't in the list of userConnections then they must be freshly online! Right?
			log.Println("[WebsocketChatHandler] user_online")
			onlineUserConnections[chatMsg.Sender] = connection
			log.Println("[WebsocketChatHandler]  Added User ", chatMsg.Sender, " to onlineUserConnections")
			log.Println("[WebsocketChatHandler] length of onlineUserConnections is:", len(onlineUserConnections))
			broadcastOnlineStatusToUsers(messageType)
		} else if chatMsg.Type == "connection_close" {
			log.Println("[WebsocketChatHandler] connection_close")
			onlineUserConnections = removeConnection(onlineUserConnections, connection)
			broadcastOnlineStatusToUsers(messageType)
			log.Println("[WebsocketChatHandler] -TESTING- Removed User ", chatMsg.Sender, " to onlineUsers")
			log.Println("[WebsocketChatHandler] length of onlineUsers is:", len(onlineUserConnections))
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

func broadcastChatToUsers(messageType int, message db.ChatMessage) {

	// changes the message to a JSON, for sending to frontend
	payload, err := json.Marshal(message)
	if err != nil {
		utils.HandleError("There has been an issue with marshalling in broadcastChatToUsers", err)
	}

	recipientConnection := userConnections[message.Recipient]
	// if recipientConnection exists (if they're online) then send a message to them too
	if recipientConnection != nil {
		err = recipientConnection.WriteMessage(messageType, payload)
		if err != nil {
			utils.HandleError("Error sending message to a client in broadcastChatToUsers:", err)
		}
	}

}

func broadcastOnlineStatusToUsers(messageType int) {
	// Collect the online userIDs from onlineUserConnections
	var onlineUsers []int
	for user := range onlineUserConnections {

		onlineUsers = append(onlineUsers, user)

	}

	// Put it in a struct to send off
	var usersToSend db.OnlineUserStruct
	usersToSend.Type = "online-notification"
	usersToSend.OnlineUsers = onlineUsers

	log.Println("onlineUsers", onlineUsers)

	// changes the message to a JSON, for sending to frontend
	payload, err := json.Marshal(usersToSend)
	if err != nil {
		utils.HandleError("There has been an issue with marshalling in broadcastOnlineStatusToUsers", err)
	}

	// Iterate over online connections and send the message
	for _, conn := range onlineUserConnections {
		err := conn.WriteMessage(messageType, payload)
		if err != nil {
			utils.HandleError("Error sending message to clients in broadcastOnlineStatusToUsers:", err)
			removeConnection(onlineUserConnections, conn)
		}
	}
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

// remove an unused connection from userConnections
func removeConnection(mapOfConnections map[int]*websocket.Conn, connectionToRemove *websocket.Conn) map[int]*websocket.Conn {
	// Search for the connection to remove in the map
	for key, conn := range mapOfConnections {
		if conn == connectionToRemove {
			delete(mapOfConnections, key)
			log.Println("Removed closed connection from mapOfConnections")
			break
		}
	}
	return mapOfConnections
}

func forwardTypingStatus(ChatMsg db.ChatMessage) {

	// Assuming you have a function to get the WebSocket connection of the recipient
	recipientConnection := userConnections[ChatMsg.Recipient]

	if recipientConnection != nil {
		if err := recipientConnection.WriteJSON(ChatMsg); err != nil {
			utils.HandleError("problem in forwardTypingStatus function", err)
		}
	}

}
