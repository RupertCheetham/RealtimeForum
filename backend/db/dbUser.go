package db

import (
	"fmt"
	"log"
	"realtimeForum/utils"
)

// Adds User to database
func AddUserToDatabase(username string, age int, gender string, firstName string, lastName string, email string, password string) error {
	_, err := Database.Exec("INSERT INTO USERS (Username, Age, Gender, First_name, Last_name, Email, Password) VALUES (?, ?, ?, ?, ?, ?, ?)", username, age, gender, firstName, lastName, email, password)
	if err != nil {
		utils.HandleError("Error adding USER to database:", err)
		// log.Println("Error adding USER to database:", err)
	}
	return err
}

// func GetUsersFromDatabase() ([]UserEntry, error) {
// 	rows, err := Database.Query("SELECT Username, Age, Gender, First_name, Last_name, Email, Password FROM USERS ORDER BY Id ASC")
// 	if err != nil {
// 		utils.HandleError("Error querying USERS from database in GetUsersFromDatabase:", err)
// 		log.Println("Error querying USERS from database in GetUsersFromDatabase:", err)
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var users []UserEntry
// 	for rows.Next() {
// 		var entry UserEntry
// 		err := rows.Scan(&entry.Username, &entry.Age, &entry.Gender, &entry.FirstName, &entry.LastName, &entry.Email, &entry.Password)
// 		if err != nil {
// 			utils.HandleError("Error scanning row from database in GetUsersFromDatabase:", err)
// 			log.Println("Error scanning row from database in GetUsersFromDatabase:", err)
// 			return nil, err
// 		}
// 		users = append(users, entry)
// 	}

// 	return users, nil
// }

func GetRecentChatUsersFromDatabase(userID string) (ChatInfo, error) {

	var sortedUsers ChatInfo

	query := `SELECT
	CASE
	  WHEN SenderID = ? THEN RecipientID
	  WHEN RecipientID = ? THEN SenderID
	END AS OtherUserID
  FROM CHAT
  WHERE (SenderID = ? OR RecipientID = ?) AND (SenderID = ? OR RecipientID = ?)
  GROUP BY OtherUserID
  ORDER BY MAX(Timestamp) DESC;
`

	rows, err := Database.Query(query, userID, userID, userID, userID, userID, userID)
	if err != nil {
		utils.HandleError("Error querying CHAT from database in GetRecentChatUsersFromDatabase:", err)
		return sortedUsers, err
	}
	defer rows.Close()

	var chatUserIds []int

	for rows.Next() {
		var entry int
		err := rows.Scan(&entry)
		if err != nil {
			utils.HandleError("Error scanning row from database in GetRecentChatUsersFromDatabase:", err)
			return sortedUsers, err
		}
		chatUserIds = prependToSlice(chatUserIds, entry)

	}

	log.Println("chatUserIDs is:", chatUserIds)
	alphabeticalUsers, err := GetUsersFromDatabase()
	if err != nil {
		utils.HandleError("Error creating list of allUsers in GetRecentChatUsersFromDatabase", err)
		return sortedUsers, err
	}

	var recentChats []UserEntry

	// for i := 0; i < len(alphabeticalUsers); i++ {
	// 	for j := 0; j < len(chatUserIds); j++ {
	// 		if alphabeticalUsers[i].Id == chatUserIds[j] {
	// 			recentChats = append(recentChats, alphabeticalUsers[i])
	// 			// Remove the entry from allUsers
	// 			alphabeticalUsers = append(alphabeticalUsers[:i], alphabeticalUsers[i+1:]...)
	// 			i--   // Decrement i to account for the removed entry
	// 			break // Exit the inner loop, since a match was found
	// 		}
	// 	}
	// }

	for j := len(chatUserIds) - 1; j > 0; j-- {
		for i := 0; i < len(alphabeticalUsers); i++ {
			if alphabeticalUsers[i].Id == chatUserIds[j] {
				recentChats = append(recentChats, alphabeticalUsers[i])
				// Remove the entry from allUsers
				alphabeticalUsers = append(alphabeticalUsers[:i], alphabeticalUsers[i+1:]...)
				i--   // Decrement i to account for the removed entry
				break // Exit the inner loop, since a match was found
			}
		}
	}

	// for i := 0; i < len(chatUserIds); i++ {
	// 	for j := 0; j < len(alphabeticalUsers); j++ {
	// 		if chatUserIds[i] == alphabeticalUsers[j].Id {
	// 			recentChats = append(recentChats, alphabeticalUsers[i])
	// 			// Remove the entry from allUsers
	// 			chatUserIds = append(chatUserIds[:i], chatUserIds[i+1:]...)
	// 			i--
	// 			break
	// 		}
	// 	}
	// }

	sortedUsers.RecentChat = recentChats
	fmt.Println("This is sortedUsers.RecentChat:", sortedUsers.RecentChat)
	sortedUsers.Alphabetical = alphabeticalUsers
	fmt.Println("This is sortedUsers.Alphabetical:", sortedUsers.Alphabetical)

	fmt.Println("This is sortedUsers:", sortedUsers)
	return sortedUsers, nil
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

func GetUsersFromDatabase() ([]UserEntry, error) {
	rows, err := Database.Query("SELECT Id, Username FROM USERS ORDER BY Username COLLATE NOCASE ASC")
	if err != nil {
		utils.HandleError("Error querying USERS from database in GetUsernamesFromDatabase:", err)
		log.Println("Error querying USERS from database in GetUsernamesFromDatabase:", err)
		return nil, err
	}
	defer rows.Close()

	var users []UserEntry
	for rows.Next() {
		var entry UserEntry
		err := rows.Scan(&entry.Id, &entry.Username)
		if err != nil {
			utils.HandleError("Error scanning row from database in GetUsersFromDatabase:", err)
			log.Println("Error scanning row from database in GetUsersFromDatabase:", err)
			return nil, err
		}
		users = append(users, entry)
	}

	return users, nil
}

func GetUsernameFromSessionID(sessionID string) string {

	// SQL query to retrieve the username associated with the provided SessionID
	query := "SELECT u.Username FROM COOKIES AS c INNER JOIN USERS AS u ON c.UserID = u.Id WHERE c.SessionID = ?"

	// Execute the query and retrieve the username
	var username string
	err := Database.QueryRow(query, sessionID).Scan(&username)
	if err != nil {
		utils.HandleError("Error finding username in GetUsernameFromSessionID:", err)
		log.Println("Error finding username in GetUsernameFromSessionID:", err)
	}

	return username
}

func GetUserIDFromSessionID(sessionID string) int {

	// SQL query to retrieve the username associated with the provided SessionID
	query := "SELECT UserID FROM COOKIES WHERE SessionID = ?"

	// Execute the query and retrieve the username
	var userID int
	err := Database.QueryRow(query, sessionID).Scan(&userID)
	if err != nil {
		utils.HandleError("Error finding userID in GetUserIDFromSessionID:", err)
		log.Println("Error finding username in GetUserIDFromSessionID:", err)
	}
	return userID
}

// returns username when given userID
func GetUsernameFromUserID(userID string) string {

	// SQL query to retrieve the username associated with the provided userID
	query := "SELECT Username FROM USERS WHERE Id = ?"

	// Execute the query and retrieve the username
	var username string
	err := Database.QueryRow(query, userID).Scan(&username)
	if err != nil {
		utils.HandleError("Error finding username in GetUsernameFromUserID:", err)
	}

	return username
}

func FindUserFromDatabase(username string) ([]UserEntry, error) {
	rows, err := Database.Query("SELECT * FROM USERS WHERE Username = ?", username)
	if err != nil {
		utils.HandleError("Error querying USERS from database in FindUserFromDatabase:", err)
		// log.Println("Error querying USERS from database in FindUserFromDatabase:", err)
		return nil, err
	}
	defer rows.Close()

	var usr []UserEntry
	for rows.Next() {
		var entry UserEntry
		err := rows.Scan(&entry.Id, &entry.Username, &entry.Age, &entry.Gender, &entry.FirstName, &entry.LastName, &entry.Email, &entry.Password)
		if err != nil {
			utils.HandleError("Error scanning row from database in FindUserFromDatabase:", err)
			// log.Println("Error scanning row from database in FindUserFromDatabase:", err)
			return nil, err
		}
		usr = append(usr, entry)
	}
	return usr, nil
}

func DeleteUserFromDatabase(username string) error {
	_, err := Database.Exec("DELETE FROM USERS WHERE Username = ?", username)
	if err != nil {
		utils.HandleError("Error querying USERS from database in DeleteUserFromDatabase:", err)
		// log.Println("Error deleting USER from database in DeleteUserFromDatabase:", err)
	} else {
		utils.WriteMessageToLogFile("User " + username + " delete")
		fmt.Println("User deleted")
	}
	return err
}

func DeleteAllUsersFromDatabase() error {
	_, err := Database.Exec("DELETE FROM USERS")
	if err != nil {
		utils.HandleError("Error querying USERS from database in DeleteUserFromDatabase:", err)
		// log.Println("Error deleting USERS from database in DeleteAllUsersFromDatabase:", err)
	} else {
		utils.WriteMessageToLogFile("All users delete from user table")
		fmt.Println("All users deleted")
	}
	return err
}
