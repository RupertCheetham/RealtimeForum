package db

import (
	"log"
	"realtimeForum/utils"
)

func AddChatToDatabase(UUID string, Message string, Sender int, Recipient int) error {
	_, err := Database.Exec("INSERT INTO CHAT (ChatUUID, Body, SenderID, RecipientID) VALUES (?, ?, ?, ?)", UUID, Message, Sender, Recipient)
	if err != nil {
		utils.HandleError("Error adding CHAT to database in AddChatToDatabase:", err)
		log.Println("Error adding CHAT to database in AddChatToDatabase:", err)
	}
	return err
}
