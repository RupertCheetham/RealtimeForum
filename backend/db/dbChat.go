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

func GetRecentChatUsersFromDatabase(userID string) (*ChatInfo, error) {

	// finds userIDs that the current user has chatted with, returns them with the newest being first
	query := `SELECT
	CASE
	  WHEN SenderID = ? THEN RecipientID
	  WHEN RecipientID = ? THEN SenderID
	END AS OtherUserID
  FROM CHAT
  WHERE (SenderID = ? OR RecipientID = ?) AND (SenderID = ? OR RecipientID = ?)
  GROUP BY OtherUserID
  ORDER BY MAX(Timestamp) ASC;
`

	rows, err := Database.Query(query, userID, userID, userID, userID, userID, userID)
	if err != nil {
		utils.HandleError("Error querying CHAT from database in GetRecentChatUsersFromDatabase:", err)
		return nil, err
	}
	defer rows.Close()

	var recentChatUserIds []int

	for rows.Next() {
		var entry int
		err := rows.Scan(&entry)
		if err != nil {
			utils.HandleError("Error scanning row from database in GetRecentChatUsersFromDatabase:", err)
			return nil, err
		}
		// list of users that the current user has chatted to
		recentChatUserIds = prependToSlice(recentChatUserIds, entry)

	}
	alphabeticalUsers, err := GetUsersFromDatabase(userID)
	if err != nil {
		utils.HandleError("Error creating list of allUsers in GetRecentChatUsersFromDatabase", err)
	}
	sortedUsers := chatsplitter(alphabeticalUsers, recentChatUserIds)
	return sortedUsers, nil

}

// splits the array of users in to ones that you've chatted with (sorted by most recent)
// and the rest (sorted alphabetically)
func chatsplitter(alphabeticalUsers []UserEntry, recentChatUserIds []int) *ChatInfo {
	var sortedUsers ChatInfo

	var recentChats []UserEntry

	for j := 0; j < len(recentChatUserIds); j++ {
		for i := 0; i < len(alphabeticalUsers); i++ {
			if alphabeticalUsers[i].Id == recentChatUserIds[j] {
				recentChats = append(recentChats, alphabeticalUsers[i])
				// Remove the entry from alphabeticalUsers
				alphabeticalUsers = append(alphabeticalUsers[:i], alphabeticalUsers[i+1:]...)
				i--   // Decrement i to account for the removed entry
				break // Exit the inner loop, since a match was found
			}
		}
	}

	sortedUsers.RecentChat = recentChats
	sortedUsers.Alphabetical = alphabeticalUsers
	return &sortedUsers

}

func prependToSlice(slice []int, elements ...int) []int {
	// Calculate the new length of the slice after adding elements
	newLen := len(slice) + len(elements)

	// Create a new slice with the new length
	newSlice := make([]int, newLen)

	// Copy the elements to be added to the front of the new slice
	copy(newSlice[len(elements):], slice)

	// Copy the existing elements to the back of the new slice
	copy(newSlice, elements)

	return newSlice
}
