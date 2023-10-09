package db

import (
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func GetLoginEntry(loginCheck UserEntry) map[string]string {
	dbLoginCheck, _ := GetUsersFromDatabase()

	message := map[string]string{
		"message": "",
	}

	// Check if the input contains an '@' sign
	isEmail := strings.Contains(loginCheck.Username, "@")

	for i := 0; i < len(dbLoginCheck); i++ {
		if (isEmail && dbLoginCheck[i].Email == loginCheck.Username) || (!isEmail && dbLoginCheck[i].Username == loginCheck.Username) {
			err := bcrypt.CompareHashAndPassword([]byte(dbLoginCheck[i].Password), []byte(loginCheck.Password))
			if err != nil {
				message["message"] = "Incorrect password"
				return message
			}
			message["message"] = "Login successful"
			return message
		}
	}
	message["message"] = "User not found"
	return message
}
