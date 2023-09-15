package db

import (
	"log"
	"realtimeForum/utils"
)

func AddUserToDatabase(username string, age int, gender string, firstName string, lastName string, email string, password string) error {
	_, err := Database.Exec("INSERT INTO USERS (Username, Age, Gender, First_name, Last_name, Email, Password) VALUES (?, ?, ?, ?, ?, ?, ?)", username, age, gender, firstName, lastName, email, password)
	if err != nil {
		utils.HandleError("Error adding USER to database:", err)
		log.Println("Error adding USER to database:", err)
	}
	return err
}

func GetUsersFromDatabase() ([]UserEntry, error) {
	rows, err := Database.Query("SELECT Username, Age, Gender, First_name, Last_name, Email, Password FROM USERS ORDER BY Id ASC")
	if err != nil {
		utils.HandleError("Error querying USERS from database in GetUsersFromDatabase:", err)
		log.Println("Error querying USERS from database in GetUsersFromDatabase:", err)
		return nil, err
	}
	defer rows.Close()

	var users []UserEntry
	for rows.Next() {
		var entry UserEntry
		err := rows.Scan(&entry.Username, &entry.Age, &entry.Gender, &entry.FirstName, &entry.LastName, &entry.Email, &entry.Password)
		if err != nil {
			utils.HandleError("Error scanning row from database in GetUsersFromDatabase:", err)
			log.Println("Error scanning row from database in GetUsersFromDatabase:", err)
			return nil, err
		}
		users = append(users, entry)
	}

	return users, nil
}
