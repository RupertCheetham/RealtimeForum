package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"realtimeForum/db"
	"realtimeForum/utils"
	"strconv"
)

// Retrieves chat history between two users based on the UserIDs in the URL
func GetChatHistoryHandler(w http.ResponseWriter, r *http.Request) {

	SetupCORS(&w, r)

	User1String := r.URL.Query().Get("user1")
	User2String := r.URL.Query().Get("user2")
	offsetString := r.URL.Query().Get("offset")
	limitString := r.URL.Query().Get("limit")
	user1, err := strconv.Atoi(User1String)

	log.Println("User1String", User1String)
	log.Println("User2String", User2String)

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

	previousChatEntryFound, chatUUID, err := db.PreviousChatChecker(user1, user2)
	if err != nil {
		utils.HandleError("There has been a problem with previousChatChecker in GetChatHistoryHandler", err)
	}
	log.Println("chatUUID", chatUUID)
	if previousChatEntryFound {
		chatHistory, err = db.GetChatFromDatabase(chatUUID, offset, limit)
	}
	if err != nil {
		utils.HandleError("Error retrieving chat history from the database:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	fmt.Println()
	// Serialize chat history to JSON and send it as the response
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(chatHistory); err != nil {
		utils.HandleError("Error encoding chat history to JSON:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
