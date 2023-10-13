package db

import (
	"log"
	"realtimeForum/utils"
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
		err := rows.Scan(&message.Sender, &message.Message, &message.Time)
		if err != nil {
			utils.HandleError("Error scanning row from database in GetChatFromDatabase:", err)
			log.Println("Error scanning row from database GetChatFromDatabase:", err)
			return nil, err
		}
		chatStruct = append(chatStruct, message)
	}

	return chatStruct, nil
}
