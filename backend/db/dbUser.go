package db

import (
	"fmt"
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

func GetUsersFromDatabase() ([]UserEntry, error) {
	rows, err := Database.Query("SELECT Id, Username FROM USERS ORDER BY Username ASC")
	if err != nil {
		utils.HandleError("Error querying USERS from database in GetUsernamesFromDatabase:", err)
		return nil, err
	}
	defer rows.Close()

	var users []UserEntry
	for rows.Next() {
		var entry UserEntry
		err := rows.Scan(&entry.Id, &entry.Username)
		if err != nil {
			utils.HandleError("Error scanning row from database in GetUsersFromDatabase:", err)
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
