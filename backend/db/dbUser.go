package db

import (
	"errors"
	"realtimeForum/utils"
	"strconv"
)

// Adds User to database
func AddUserToDatabase(username string, age int, gender string, firstName string, lastName string, email string, password string) error {
	_, err := Database.Exec("INSERT INTO USERS (Username, Age, Gender, First_name, Last_name, Email, Password) VALUES (?, ?, ?, ?, ?, ?, ?)", username, age, gender, firstName, lastName, email, password)
	if err != nil {
		utils.HandleError("Error adding USER to database:", err)
	}
	return err
}

// returns a list of users from the DB, excluding the current user
func GetUsersFromDatabase(userIDString string) ([]UserEntry, error) {
	rows, err := Database.Query("SELECT Id, Username FROM USERS ORDER BY Username COLLATE NOCASE ASC")
	if err != nil {
		utils.HandleError("Error querying USERS from database in GetUsersFromDatabase:", err)
		return nil, err
	}
	defer rows.Close()

	var users []UserEntry
	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		utils.HandleError("Error AtoIing userID in GetUsersFromDatabase:", err)
		return nil, err
	}
	for rows.Next() {
		var entry UserEntry
		err := rows.Scan(&entry.Id, &entry.Username)
		if err != nil {
			utils.HandleError("Error scanning row from database in GetUsersFromDatabase:", err)
			return nil, err
		}
		if entry.Id != userID {
			users = append(users, entry)
		}
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

func FindUserFromDatabase(username string) (UserEntry, error) {
	rows, err := Database.Query("SELECT * FROM USERS WHERE Username = ?", username)
	if err != nil {
		utils.HandleError("Error querying USERS from database in FindUserFromDatabase:", err)
		return UserEntry{}, err
	}
	defer rows.Close()

	var user UserEntry
	if !rows.Next() {
		return UserEntry{}, errors.New("user not found")
	}

	err = rows.Scan(&user.Id, &user.Username, &user.Age, &user.Gender, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		utils.HandleError("Error scanning row from database in FindUserFromDatabase:", err)
		return UserEntry{}, err
	}

	// fmt.Println("user from finduserfromdb function:", user, "error:", err)

	return user, err
}

func DeleteUserFromDatabase(username string) error {
	_, err := Database.Exec("DELETE FROM USERS WHERE Username = ?", username)
	if err != nil {
		utils.HandleError("Error querying USERS from database in DeleteUserFromDatabase:", err)
	} else {
		utils.WriteMessageToLogFile("User " + username + " delete")
		// fmt.Println("User deleted")
	}
	return err
}

func DeleteAllUsersFromDatabase() error {
	_, err := Database.Exec("DELETE FROM USERS")
	if err != nil {
		utils.HandleError("Error querying USERS from database in DeleteUserFromDatabase:", err)
	} else {
		utils.WriteMessageToLogFile("All users delete from user table")
		// fmt.Println("All users deleted")
	}
	return err
}
