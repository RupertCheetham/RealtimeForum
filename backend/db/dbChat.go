package db

import (
	"log"
	"realtimeForum/utils"
	"time"
)

// inserts new chat message to db
func AddChatToDatabase(UUID string, Message string, Sender int, Recipient int) error {
	_, err := Database.Exec("INSERT INTO CHAT (ChatUUID, Body, SenderID, RecipientID) VALUES (?, ?, ?, ?)", UUID, Message, Sender, Recipient)
	if err != nil {
		utils.HandleError("Error adding CHAT to database in AddChatToDatabase:", err)
		log.Println("Error adding CHAT to database in AddChatToDatabase:", err)
	}

	return err
}

// retrieves chat messages for a particular chatUUID from db
func GetChatFromDatabase(UUID string) ([]ChatMessage, error) {
	var chatStruct []ChatMessage

	rows, err := Database.Query("SELECT SenderID, Body, Timestamp FROM CHAT WHERE ChatUUID = ?", UUID)
	if err != nil {
		utils.HandleError("Error selecting chat from UUID from database:", err)
		log.Println("Error selecting chat from UUID from database:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var message ChatMessage
		var timestamp string
		err := rows.Scan(&message.Sender, &message.Message, &timestamp)
		if err != nil {
			utils.HandleError("Error scanning row from database in GetChatFromDatabase:", err)
			log.Println("Error scanning row from database GetChatFromDatabase:", err)
			return nil, err
		}

		// Parse the timestamp string
		timeObj, parseErr := time.Parse(time.RFC3339, timestamp)
		if parseErr != nil {
			utils.HandleError("Error parsing timestamp:", parseErr)
			log.Println("Error parsing timestamp:", parseErr)
			return nil, parseErr
		}

		// Add an hour to the time
		timeWithHourAdded := timeObj.Add(time.Hour)

		// Format the time as a string in the desired format
		formattedTime := timeWithHourAdded.Format("15:04:05 02-01-2006")

		// Assign the formatted time to the message
		message.Time = formattedTime

		chatStruct = append(chatStruct, message)
	}

	return chatStruct, nil
}