package db

import (
	"fmt"
)

func GetLoginEntry(loginCheck UserEntry) map[string]string {
	dbLoginCheck, _ := GetUsersFromDatabase()

	message := map[string]string{
		"message": "",
	}

	for i := 0; i < len(dbLoginCheck); i++ {
		if dbLoginCheck[i].Username == loginCheck.Username && dbLoginCheck[i].Password == loginCheck.Password {
			fmt.Println("Login successful")
			message["message"] = "Login successful"
			return message
		}
	}
	fmt.Println("Incorrect login details")
	message["message"] = "Incorrect login details"
	return message
}
