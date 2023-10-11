package db

import (
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func GetLoginEntry(loginCheck UserEntry) (message map[string]any, err error) {
	// dbLoginCheck, _ := GetUsersFromDatabase()
	dbLoginCheck, _ := FindUserFromDatabase(loginCheck.Username)

	message = map[string]any{
		"message": "",
	}

	// // Check if the input contains an '@' sign
	isEmail := strings.Contains(loginCheck.Username, "@")

	if (isEmail && dbLoginCheck[0].Email == loginCheck.Username) || (!isEmail && dbLoginCheck[0].Username == loginCheck.Username) {
		err := bcrypt.CompareHashAndPassword([]byte(dbLoginCheck[0].Password), []byte(loginCheck.Password))
		if err != nil {
			message["message"] = "Incorrect username or password"
			return message, err
		}
		message["message"] = "Login successfully"
		return message, err
	}

	return message, err
}
