package db

import (
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func GetLoginEntry(loginCheck UserEntry) (message map[string]any, err error) {
	message = map[string]any{
		"message":  "",
		"id":       0,
		"username": "",
	}

	dbLoginCheck, err := FindUserFromDatabase(loginCheck.Username)
	if err != nil {
		return message, err
	}

	if err != nil {
		return message, err
	}

	// Check if the input contains an '@' sign
	isEmail := strings.Contains(loginCheck.Username, "@")

	if (isEmail && dbLoginCheck.Email == loginCheck.Email) || (!isEmail && dbLoginCheck.Username == loginCheck.Username) {
		err := bcrypt.CompareHashAndPassword([]byte(dbLoginCheck.Password), []byte(loginCheck.Password))
		if err != nil {
			message["message"] = "Incorrect username or password"
			return message, err
		} else {
			message["message"] = "Login successfully"
			message["id"] = dbLoginCheck.Id
			message["username"] = dbLoginCheck.Username
		}
	}
	return message, err
}
