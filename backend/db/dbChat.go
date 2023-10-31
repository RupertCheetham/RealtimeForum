package db

import (
	"realtimeForum/utils"
	"time"
)

// inserts new chat message to db
func AddChatToDatabase(UUID string, Message string, Sender int, Recipient int) error {
	_, err := Database.Exec("INSERT INTO CHAT (ChatUUID, Body, SenderID, RecipientID) VALUES (?, ?, ?, ?)", UUID, Message, Sender, Recipient)
	if err != nil {
		utils.HandleError("Error adding CHAT to database in AddChatToDatabase:", err)
	}

	return err
}

// retrieves chat messages for a particular chatUUID from db
func GetChatFromDatabase(UUID string, offset int, limit int) ([]ChatMessage, error) {

	var chatStruct []ChatMessage
	// Query to retrieve chat entries starting from the most recent entries working backward
	query := `
		SELECT SenderID, Body, Timestamp
		FROM CHAT
		WHERE ChatUUID = ?
		ORDER BY Timestamp DESC
		LIMIT ? OFFSET ?
	`

	rows, err := Database.Query(query, UUID, limit, offset)
	if err != nil {
		utils.HandleError("Error selecting chat from UUID from database:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var message ChatMessage
		var timestamp string
		err := rows.Scan(&message.Sender, &message.Body, &timestamp)
		if err != nil {
			utils.HandleError("Error scanning row from database in GetChatFromDatabase:", err)
			return nil, err
		}

		if message.Body != "" {
			// Parse the timestamp string
			timeObj, parseErr := time.Parse(time.RFC3339, timestamp)
			if parseErr != nil {
				utils.HandleError("Error parsing timestamp:", parseErr)
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
	}

	return chatStruct, nil
}
